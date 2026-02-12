package repository

import (
	"errors"
	"training-plan-api/helper"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type CertificateRepositoryImpl struct {
	Db *gorm.DB
}
func NewCertificateRepositoryImpl(db *gorm.DB) CertificateRepository {
	return &CertificateRepositoryImpl{Db: db}
}

// FindRecordByIDAndUserID implements CertificateRepository.
func (r *CertificateRepositoryImpl) FindRecordByIDAndUserID(recordID int, userID uint) (*model.Record, error) {
	var record model.Record

	err := r.Db.
		Where("id = ? AND user_id = ?", recordID, userID).
		First(&record).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &record, err
}



func (r *CertificateRepositoryImpl) Save(certificate *model.Certificate) error {
	return r.Db.Create(certificate).Error
}

func (r *CertificateRepositoryImpl) FindById(id int) (*model.Certificate, error) {
	var certificate model.Certificate

	err := r.Db.First(&certificate, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NotFound("certificate not found")
		}
		return nil, err
	}

	return &certificate, nil
}

func (r *CertificateRepositoryImpl) FindByUserId(userId int) ([]model.Certificate, error) {
	var certificates []model.Certificate

	err := r.Db.
		Preload("User").
		Preload("Training").
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Find(&certificates).
		Error

	return certificates, err
}

func (r *CertificateRepositoryImpl) Delete(id int) error {
	result := r.Db.Delete(&model.Certificate{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return helper.NotFound("certificate not found")
	}

	return nil
}

func (r *CertificateRepositoryImpl) UpdateStatus(
	id int,
	status model.CertificateStatus,
) error {
	result := r.Db.Model(&model.Certificate{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return helper.NotFound("certificate not found")
	}

	return nil
}

func (r *CertificateRepositoryImpl) FindAllPending(
	offset, limit int,
) ([]model.Certificate, int64, error) {

	var certificates []model.Certificate
	var total int64

	r.Db.Model(&model.Certificate{}).
		Where("status = ?", model.CertPending).
		Count(&total)

	err := r.Db.
		Preload("User").
		Where("status = ?", model.CertPending).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&certificates).
		Error

	return certificates, total, err
}
