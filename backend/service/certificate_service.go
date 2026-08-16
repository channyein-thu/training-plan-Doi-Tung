package service

import (
	"fmt"
	"log"
	"math"
	"mime/multipart"
	"strings"
	"time"

	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
)

// imagePathPrefix is prepended to the storage object path when a certificate is
// saved (see Upload). The Storage layer works with object paths relative to its
// own base dir, so this prefix must be stripped back off before calling
// storage.Delete — otherwise the base path gets doubled up and the file is
// never actually removed.
const imagePathPrefix = "uploads/"

// storageObjectPath converts a stored Certificate.Image value back into the
// storage-relative object path that the Storage layer expects.
func storageObjectPath(image string) string {
	return strings.TrimPrefix(image, imagePathPrefix)
}

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
		if err := c.storage.Delete(storageObjectPath(cert.Image)); err != nil {
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

			// Guard Department under the User nil-check — a certificate whose
			// owner was deleted has a nil User, and reading cert.User.Department
			// directly would panic (crashing the whole list endpoint).
			if cert.User.Department != nil {
				resp.Department = cert.User.Department.Name
				resp.Division = string(cert.User.Department.Division)
			}
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

			// Guard Department under the User nil-check — a certificate whose
			// owner was deleted has a nil User, and reading cert.User.Department
			// directly would panic (crashing the whole list endpoint).
			if cert.User.Department != nil {
				resp.Department = cert.User.Department.Name
				resp.Division = string(cert.User.Department.Division)
			}
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
		Image:       imagePathPrefix + objectPath,
		Description: req.Description,
		Status:      model.CertPending,
	}

	if err := c.repo.Save(certificate); err != nil {
		_ = c.storage.Delete(objectPath)
		return err
	}

	return nil
}

// GetCertificateFilePath authorizes the caller and returns the stored image
// path for streaming. Admins may fetch any certificate file; everyone else may
// only fetch their own. The returned path is the server-generated Image value
// (e.g. "uploads/certificates/user_5/173...jpeg") — never a client-supplied
// path — so there is no traversal risk.
func (c *CertificateServiceImpl) GetCertificateFilePath(
	certificateID int,
	callerID uint,
	callerRole string,
) (string, error) {
	cert, err := c.repo.FindById(certificateID)
	if err != nil {
		return "", err
	}

	if model.Role(callerRole) != model.RoleHRAdmin && cert.UserID != callerID {
		return "", helper.Forbidden("You don't have permission to view this certificate")
	}

	if cert.Image == "" {
		return "", helper.NotFound("certificate file not found")
	}

	return cert.Image, nil
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
		if err := c.storage.Delete(storageObjectPath(certificate.Image)); err != nil {
			log.Println("⚠ failed to delete certificate file:", err)
		}
	}

	return nil
}
