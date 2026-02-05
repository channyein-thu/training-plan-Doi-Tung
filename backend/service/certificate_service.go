package service

import (
	"fmt"
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
	certificateRepo repository.CertificateRepository,
	validate *validator.Validate,
	redisClient *redis.Client,
	storage helper.Storage,
) CertificateService {
	return &CertificateServiceImpl{
		repo:     certificateRepo,
		validate: validate,
		cache:    redisClient,
		storage:  storage,
	}
}


// FindByCurrentUser implements CertificateService.
func (c *CertificateServiceImpl) FindByCurrentUser(userID uint) ([]response.CertificateResponse, error) {
	certificates, err := c.repo.FindByUserId(int(userID))
	if err != nil {
		return nil, err
	}

	var responses []response.CertificateResponse
	for _, cert := range certificates {
		trainingName := ""
		if cert.Training != nil {
			trainingName = cert.Training.Name
		}

		userName := ""
		if cert.User != nil {
			userName = cert.User.Name
		}

		responses = append(responses, response.CertificateResponse{
			ID:           cert.ID,
			UserID:       cert.UserID,
			UserName:     userName,
			TrainingName: trainingName,
			Image:        cert.Image,
			Description:  cert.Description,
			Status:       string(cert.Status),
			CreatedAt:    cert.CreatedAt,
			UpdatedAt:    cert.UpdatedAt,
		})
	}

	return responses, nil
}

// Upload implements CertificateService.
func (c *CertificateServiceImpl) Upload(userID uint, req request.CreateCertificateRequest, fileHeader *multipart.FileHeader) error {
	if err := c.validate.Struct(req); err != nil {
		return helper.ValidationError(
		helper.FormatValidationError(err),
	)
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return helper.BadRequest("Failed to open uploaded file")
	}
	defer file.Close()

	// Generate unique file path: certificates/{userID}/{timestamp}_{filename}
	ext := filepath.Ext(fileHeader.Filename)
	timestamp := time.Now().Unix()
	filePath := fmt.Sprintf("certificates/%d/%d%s", userID, timestamp, ext)

	// Upload file to storage
	fileURL, err := c.storage.Upload(filePath, file, fileHeader.Header.Get("Content-Type"))
	if err != nil {
		return helper.Internal("Failed to upload certificate")
	}

	certificate := &model.Certificate{
		UserID:      userID,
		TrainingID:   req.TrainingID,
		Image:       fileURL,
		Description: req.Description,
		Status:      model.CertPending,
	}

	return c.repo.Save(certificate)
}

// Delete implements CertificateService.
func (c *CertificateServiceImpl) Delete(certificateID int, userID uint) error {
	certificate, err := c.repo.FindById(certificateID)
	if err != nil {
		return err
	}

	// Verify ownership
	if certificate.UserID != userID {
		return helper.Forbidden("You don't have permission to delete this certificate")
	}

	// Delete from database first
	if err := c.repo.Delete(certificateID); err != nil {
		return err
	}

	// Try to delete file from storage (don't fail if file doesn't exist)
	if certificate.Image != "" {
		// Extract the file path from the URL (remove leading slash)
		filePath := certificate.Image
		if len(filePath) > 0 && filePath[0] == '/' {
			filePath = filePath[1:]
		}
		_ = c.storage.Delete(filePath) // Ignore error if file doesn't exist
	}

	return nil
}

