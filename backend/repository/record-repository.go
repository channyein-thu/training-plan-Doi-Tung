package repository

import (
	"errors"
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
func (r *RecordRepositoryImpl) Exists(userId uint, courseId uint) bool {
		var count int64
	r.Db.Model(&model.Record{}).
		Where("user_id = ? AND course_id = ?", userId, courseId).
		Count(&count)

	return count > 0
}


// FindById implements RecordRepository.
func (r *RecordRepositoryImpl) FindById(id int) (*model.Record, error) {
		var record model.Record

	err := r.Db.
		Preload("User").
		Preload("Course").
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

