package service

import (
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
)

type RecordServiceImpl struct {
	repo     repository.RecordRepository
	validate *validator.Validate
}

func NewRecordServiceImpl(
	repo repository.RecordRepository,
	validate *validator.Validate,
) RecordService {
	return &RecordServiceImpl{
		repo:     repo,
		validate: validate,
	}
}

func (s *RecordServiceImpl) RegisterStaff(
	courseId uint,
	req request.RegisterStaffRequest,
) error {

	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	for _, userId := range req.UserIDs {
		if s.repo.Exists(userId, courseId) {
			continue // prevent duplicate registration
		}

		record := &model.Record{
			UserID:   userId,
			CourseID: courseId,
			Status:   model.RecordStatusRegister,
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
		ID:         record.ID,
		UserID:     record.UserID,
		UserName:   record.User.Name,
		CourseID:   record.CourseID,
		CourseName: record.Course.Name,
		Status:     string(record.Status),
		CreatedAt:  record.CreatedAt,
		UpdatedAt:  record.UpdatedAt,
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
