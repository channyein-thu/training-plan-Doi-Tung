package repository

import (
	"errors"
	"training-plan-api/data/request"
	"training-plan-api/helper"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type RecordRepositoryImpl struct {
	Db *gorm.DB
}

func NewRecordRepositoryImpl(db *gorm.DB) RecordRepository {
	return &RecordRepositoryImpl{Db: db}
}

// Search implements RecordRepository.
func (r *RecordRepositoryImpl) Search(
	req request.RecordFilterRequest,
) ([]model.Record, int64, error) {

	var records []model.Record
	var total int64

	query := r.Db.
		Model(&model.Record{}).
		Preload("User").
		Preload("User.Department").
		Preload("TrainingPlan").
		Joins("JOIN users ON users.id = records.user_id").
		Joins("JOIN training_plans ON training_plans.id = records.training_plan_id")

	// Department filter
	if len(req.DepartmentIDs) > 0 {
		query = query.Where("users.department_id IN ?", req.DepartmentIDs)
	}

	// Category filter
	if len(req.Categories) > 0 {
		query = query.Where("training_plans.category IN ?", req.Categories)
	}

	// Status filter
	if req.Status != nil && *req.Status != "" {
		query = query.Where("records.status = ?", *req.Status)
	}

	// Date range filter
	if req.StartDate != nil && req.EndDate != nil {
		query = query.Where("training_plans.date BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}

	// Count first
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit

	err := query.
		Order("records.created_at DESC").
		Offset(offset).
		Limit(req.Limit).
		Find(&records).
		Error

	return records, total, err
}


// FindByUserId implements RecordRepository.
func (r *RecordRepositoryImpl) FindByUserId(userID uint, offset int, limit int) ([]model.Record, int64, error) {
	var records []model.Record
	var total int64

	query := r.Db.
		Preload("User").
		Preload("TrainingPlan").
		Where("user_id = ?", userID)

	err := query.Model(&model.Record{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&records).Error
	return records, total, err
}


// FindByManagerDepartment implements RecordRepository.
func (r *RecordRepositoryImpl) FindByManagerDepartment(departmetnID int, offset int, limit int) ([]model.Record, int64, error) {
	var records []model.Record
	var total int64

	query := r.Db.
		Preload("User").
		Preload("User.Department").
		Preload("TrainingPlan").
		Joins("JOIN users ON users.id = records.user_id").
		Where("users.department_id = ?", departmetnID)

	err := query.Model(&model.Record{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).Order("records.created_at DESC").Find(&records).Error
	return records, total, err
}

// Delete implements RecordRepository.
func (r *RecordRepositoryImpl) Delete(id int) error {
	result := r.Db.Delete(&model.Record{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return helper.NotFound("record not found")
	}
	return nil
}

// Exists implements RecordRepository.
func (r *RecordRepositoryImpl) Exists(userId uint, trainingPlanId uint) bool {
	var count int64
	r.Db.Model(&model.Record{}).
		Where("user_id = ? AND training_plan_id = ?", userId, trainingPlanId).
		Count(&count)

	return count > 0
}

// FindById implements RecordRepository.
func (r *RecordRepositoryImpl) FindById(id int) (*model.Record, error) {
	var record model.Record

	err := r.Db.
		Preload("User").
		Preload("TrainingPlan").
		First(&record, id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, helper.NotFound("record not found")
	}

	return &record, err
}

// Save implements RecordRepository.
func (r *RecordRepositoryImpl) Save(record *model.Record) error {
	return r.Db.Create(record).Error
}

// Update implements RecordRepository.
func (r *RecordRepositoryImpl) Update(record *model.Record) error {
	return r.Db.Save(record).Error
}
