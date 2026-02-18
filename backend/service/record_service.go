package service

import (
	"math"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
)

type RecordServiceImpl struct {
	repo     repository.RecordRepository
	userRepo repository.UserRepository
	validate *validator.Validate
}

func NewRecordServiceImpl(
	repo repository.RecordRepository,
	userRepo repository.UserRepository,
	validate *validator.Validate,
) RecordService {
	return &RecordServiceImpl{
		repo:     repo,
		userRepo: userRepo,
		validate: validate,
	}
}

// FindByUser implements RecordService.
func (s *RecordServiceImpl) FindByUser(userID uint, page int, limit int) (response.PaginatedResponse[response.RecordResponse], error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	records, total, err := s.repo.FindByUserId(userID, offset, limit)
	if err != nil {
		return response.PaginatedResponse[response.RecordResponse]{}, err
	}

	items := make([]response.RecordResponse, 0, len(records))

	for _, r := range records {

		resp := response.RecordResponse{
			ID:             r.ID,
			UserID:         r.UserID,
			TrainingPlanID: r.TrainingPlanID,
			Status:         string(r.Status),
			CreatedAt:      r.CreatedAt,
			UpdatedAt:      r.UpdatedAt,
		}

		if r.User != nil {
			resp.UserName = r.User.Name
		}

		if r.TrainingPlan != nil {
			resp.TrainingPlanName = r.TrainingPlan.Name
		}

		items = append(items, resp)
	}

	return response.PaginatedResponse[response.RecordResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		},
	}, nil
}



func (s *RecordServiceImpl) RegisterStaff(
	trainingPlanId uint,
	req request.RegisterStaffRequest,
) error {

	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	for _, userId := range req.UserIDs {
		if s.repo.Exists(userId, trainingPlanId) {
			continue // prevent duplicate registration
		}

		record := &model.Record{
			UserID:         userId,
			TrainingPlanID: trainingPlanId,
			Status:         model.RecordStatusRegister,
		}

		if err := s.repo.Save(record); err != nil {
			return err
		}
	}

	return nil
}

func (s *RecordServiceImpl) FindById(
	id int,
) (response.RecordResponse, error) {

	record, err := s.repo.FindById(id)
	if err != nil {
		return response.RecordResponse{}, err
	}

	return response.RecordResponse{
		ID:               record.ID,
		UserID:           record.UserID,
		UserName:         record.User.Name,
		TrainingPlanID:   record.TrainingPlanID,
		TrainingPlanName: record.TrainingPlan.Name,
		Status:           string(record.Status),
		CreatedAt:        record.CreatedAt,
		UpdatedAt:        record.UpdatedAt,
	}, nil
}

func (s *RecordServiceImpl) Update(
	id int,
	req request.UpdateRecordRequest,
) error {

	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	record, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	record.Status = req.Status
	return s.repo.Update(record)
}

func (s *RecordServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *RecordServiceImpl) FindByManager(
	managerID uint,
	page int,
	limit int,
) (response.PaginatedResponse[response.RecordResponse], error) {

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	manager, err := s.userRepo.FindById(managerID)
	if err != nil {
		return response.PaginatedResponse[response.RecordResponse]{}, err
	}

	if manager.Role != model.RoleDepartmentManager {
		return response.PaginatedResponse[response.RecordResponse]{},
			helper.Forbidden("only managers can access this resource")
	}

	offset := (page - 1) * limit

	records, total, err := s.repo.FindByManagerDepartment(
		manager.DepartmentID,
		offset,
		limit,
	)
	if err != nil {
		return response.PaginatedResponse[response.RecordResponse]{}, err
	}

	items := make([]response.RecordResponse, 0, len(records))

	for _, r := range records {

		resp := response.RecordResponse{
			ID:             r.ID,
			UserID:         r.UserID,
			TrainingPlanID: r.TrainingPlanID,
			Status:         string(r.Status),
			CreatedAt:      r.CreatedAt,
			UpdatedAt:      r.UpdatedAt,
		}

		if r.User != nil {
			resp.UserName = r.User.Name
		}

		if r.TrainingPlan != nil {
			resp.TrainingPlanName = r.TrainingPlan.Name
		}

		items = append(items, resp)
	}

	return response.PaginatedResponse[response.RecordResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		},
	}, nil
}
