package service

import (
	"fmt"
	"log"
	"math"
	"mime/multipart"
	"time"

	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
)

type CertificateServiceImpl struct {
	repo         repository.CertificateRepository
	validate     *validator.Validate
	storage      helper.Storage
	notifService NotificationService
}

func NewCertificateServiceImpl(
	repo repository.CertificateRepository,
	validate *validator.Validate,
	storage helper.Storage,
	notifService NotificationService,
) CertificateService {
	return &CertificateServiceImpl{
		repo:         repo,
		validate:     validate,
		storage:      storage,
		notifService: notifService,
	}
}



// ================= ADMIN =================

func (c *CertificateServiceImpl) Approve(certificateID int) error {
	cert, err := c.repo.FindById(certificateID)
	if err != nil {
		return err
	}

	if cert.Status != model.CertPending {
		return helper.BadRequest("certificate is not pending")
	}

	if err := c.repo.UpdateStatus(certificateID, model.CertApproved); err != nil {
		return err
	}

	trainingName := "your training"
	if cert.Training != nil {
		trainingName = cert.Training.Name
	}
	if err := c.notifService.Create(
		cert.UserID,
		model.NotifCertificateApproved,
		"Certificate Approved",
		fmt.Sprintf("Your certificate for \"%s\" has been approved.", trainingName),
	); err != nil {
		log.Println("⚠ failed to create approval notification:", err)
	}

	return nil
}

func (c *CertificateServiceImpl) Reject(certificateID int) error {
	cert, err := c.repo.FindById(certificateID)
	if err != nil {
		return err
	}

	if cert.Status != model.CertPending {
		return helper.BadRequest("certificate is not pending")
	}

	if err := c.repo.Delete(certificateID); err != nil {
		return err
	}

	if cert.Image != "" {
		if err := c.storage.Delete(cert.Image); err != nil {
			log.Println("⚠ failed to delete certificate file:", err)
		}
	}

	trainingName := "your training"
	if cert.Training != nil {
		trainingName = cert.Training.Name
	}
	if err := c.notifService.Create(
		cert.UserID,
		model.NotifCertificateRejected,
		"Certificate Rejected",
		fmt.Sprintf("Your certificate for \"%s\" has been rejected.", trainingName),
	); err != nil {
		log.Println("⚠ failed to create rejection notification:", err)
	}

	return nil
}

func (c *CertificateServiceImpl) FindAllPending(
	page, limit int,
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
			resp.EmployeeID = cert.User.EmployeeID
		}
		if cert.User.Department != nil {
			resp.Department = cert.User.Department.Name
			resp.Division = string(cert.User.Department.Division)
		}
		if cert.Training != nil {
			resp.TrainingID = cert.TrainingID
			resp.TrainingName = cert.Training.Name
			resp.Category = string(cert.Training.Category)
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

// ================= USER =================

func (c *CertificateServiceImpl) FindByCurrentUser(
	userID uint,
) ([]response.CertificateResponse, error) {

	certificates, err := c.repo.FindByUserId(int(userID))
	if err != nil {
		return nil, err
	}

	responses := make([]response.CertificateResponse, 0, len(certificates))
	for _, cert := range certificates {
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
			resp.EmployeeID = cert.User.EmployeeID
		}
		if cert.User.Department != nil {
			resp.Department = cert.User.Department.Name
			resp.Division = string(cert.User.Department.Division)
		}
		if cert.Training != nil {
			resp.TrainingID = cert.TrainingID
			resp.TrainingName = cert.Training.Name
			resp.Category = string(cert.Training.Category)
		}

		responses = append(responses, resp)
	}

	return responses, nil
}

func (c *CertificateServiceImpl) Upload(
	userID uint,
	req request.CreateCertificateRequest,
	fileHeader *multipart.FileHeader,
) error {

	if err := c.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	// Enforces a size cap and an extension whitelist, and sniffs the file's
	// actual leading bytes to confirm they match the claimed type — the
	// client-supplied filename and Content-Type header are not trusted.
	file, ext, mimeType, err := helper.ValidateCertificateFile(fileHeader)
	if err != nil {
		return err
	}
	defer file.Close()

	objectPath := fmt.Sprintf(
		"certificates/user_%d/%d%s",
		userID,
		time.Now().Unix(),
		ext,
	)

	if _, err := c.storage.Upload(
		objectPath,
		file,
		mimeType,
	); err != nil {
		return helper.Internal("Failed to upload certificate")
	}

	certificate := &model.Certificate{
		UserID:      userID,
		TrainingID:  req.TrainingID,
		Image:       "uploads/" + objectPath,
		Description: req.Description,
		Status:      model.CertPending,
	}

	if err := c.repo.Save(certificate); err != nil {
		_ = c.storage.Delete(objectPath)
		return err
	}

	return nil
}

func (c *CertificateServiceImpl) Delete(
	certificateID int,
	userID uint,
) error {

	certificate, err := c.repo.FindById(certificateID)
	if err != nil {
		return err
	}

	if certificate.UserID != userID {
		return helper.Forbidden("You don't have permission to delete this certificate")
	}

	if err := c.repo.Delete(certificateID); err != nil {
		return err
	}

	if certificate.Image != "" {
		_ = c.storage.Delete(certificate.Image)
	}

	return nil
}
