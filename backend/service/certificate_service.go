package service

import (
	"fmt"
	"log"
	"math"
	"mime/multipart"
	"path/filepath"
	"time"

	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type CertificateServiceImpl struct {
	repo     repository.CertificateRepository
	validate *validator.Validate
	cache    *redis.Client
	storage  helper.Storage
}
func NewCertificateServiceImpl(
	repo repository.CertificateRepository,
	validate *validator.Validate,
	redisClient *redis.Client,
	storage helper.Storage,
) CertificateService {
	return &CertificateServiceImpl{
		repo:     repo,
		validate: validate,
		cache:    redisClient,
		storage:  storage,
	}
}

// Approve implements CertificateService.
func (c *CertificateServiceImpl) Approve(certificateID int) error {
	cert, err := c.repo.FindById(certificateID)
	if err != nil {
		return err
	}

	if cert.Status != model.CertPending {
		return helper.BadRequest("certificate is not pending")
	}

	return c.repo.UpdateStatus(certificateID, model.CertApproved)
}


// FindAllPending implements CertificateService.
func (c *CertificateServiceImpl) FindAllPending(
	page int,
	limit int,
) (response.PaginatedResponse[response.CertificateResponse], error) {

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	certs, total, err := c.repo.FindAllPending(offset, limit)
	if err != nil {
		return response.PaginatedResponse[response.CertificateResponse]{}, err
	}

	items := make([]response.CertificateResponse, 0, len(certs))
	for _, cert := range certs {
		resp := response.CertificateResponse{
			ID:          cert.ID,
			UserID:      cert.UserID,
			Image:       cert.Image,
			Description: cert.Description,
			Status:      string(cert.Status),
			CreatedAt:   cert.CreatedAt,
			UpdatedAt:   cert.UpdatedAt,
		}

		if cert.User != nil {
			resp.UserName = cert.User.Name
		}

		items = append(items, resp)
	}

	return response.PaginatedResponse[response.CertificateResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		},
	}, nil
}

// Reject implements CertificateService.
func (c *CertificateServiceImpl) Reject(certificateID int) error {
	// Find certificate
	cert, err := c.repo.FindById(certificateID)
	if err != nil {
		return err
	}

	// Only pending certificates can be rejected
	if cert.Status != model.CertPending {
		return helper.BadRequest("certificate is not pending")
	}

	// Delete DB record FIRST
	if err := c.repo.Delete(certificateID); err != nil {
		return err
	}

	// Best-effort file delete
	if cert.Image != "" {
		if err := c.storage.Delete(cert.Image); err != nil {
			// Do NOT fail request – DB already clean
			// Just log it
			log.Println("⚠ failed to delete certificate file:", err)
		}
	}

	return nil
}




// GET CERTIFICATES BY CURRENT USER
func (c *CertificateServiceImpl) FindByCurrentUser(userID uint) ([]response.CertificateResponse, error) {
	certificates, err := c.repo.FindByUserId(int(userID))
	if err != nil {
		return nil, err
	}

	responses := make([]response.CertificateResponse, 0, len(certificates))

	for _, cert := range certificates {
		resp := response.CertificateResponse{
			ID:          cert.ID,
			UserID:      cert.UserID,
			Image:       cert.Image, // object path
			Description: cert.Description,
			Status:      string(cert.Status),
			CreatedAt:   cert.CreatedAt,
			UpdatedAt:   cert.UpdatedAt,
		}

		if cert.User != nil {
			resp.UserName = cert.User.Name
		}

		if cert.Training != nil {
			resp.TrainingID = cert.TrainingID
			resp.TrainingName = cert.Training.Name
		}

		responses = append(responses, resp)
	}

	return responses, nil
}

// UPLOAD CERTIFICATE (WITH ROLLBACK)
func (c *CertificateServiceImpl) Upload(
	userID uint,
	req request.CreateCertificateRequest,
	fileHeader *multipart.FileHeader,
) error {

	if err := c.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	// Open uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return helper.BadRequest("Failed to open uploaded file")
	}
	defer file.Close()

	// Generate object path (NOT URL)
	ext := filepath.Ext(fileHeader.Filename)
	objectPath := fmt.Sprintf(
		"certificates/user_%d/%d%s",
		userID,
		time.Now().Unix(),
		ext,
	)

	// Upload to storage
	if _, err := c.storage.Upload(
		objectPath,
		file,
		fileHeader.Header.Get("Content-Type"),
	); err != nil {
		return helper.Internal("Failed to upload certificate")
	}

	// Prepare DB record
	certificate := &model.Certificate{
		UserID:      userID,
		TrainingID:  req.TrainingID,
		Image:       objectPath, // STORE OBJECT PATH
		Description: req.Description,
		Status:      model.CertPending,
	}

	// Save DB record
	if err := c.repo.Save(certificate); err != nil {
		// ROLLBACK FILE if DB fails
		_ = c.storage.Delete(objectPath)
		return err
	}

	return nil
}

// DELETE CERTIFICATE (DB FIRST, FILE SECOND)
func (c *CertificateServiceImpl) Delete(certificateID int, userID uint) error {
	certificate, err := c.repo.FindById(certificateID)
	if err != nil {
		return err
	}

	// Ownership check
	if certificate.UserID != userID {
		return helper.Forbidden("You don't have permission to delete this certificate")
	}

	// Delete DB record first
	if err := c.repo.Delete(certificateID); err != nil {
		return err
	}

	// Best-effort file delete
	if certificate.Image != "" {
		_ = c.storage.Delete(certificate.Image)
	}

	return nil
}
