package repository

import (
	"training-plan-api/helper"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type CertificateRepositoryImpl struct{
	DB *gorm.DB
}

func NewCertificateRepositoryImpl(db *gorm.DB) CertificateRepository {
	return &CertificateRepositoryImpl{DB: db}
}

func (r *CertificateRepositoryImpl) Save(certificate *model.Certificate) error {
	err := r.DB.Create(certificate).Error
	return err
}

func (r *CertificateRepositoryImpl) FindById(id int) (model.Certificate, error) {
	var certificate model.Certificate
	err := r.DB.First(&certificate, id).Error
	return certificate, err
}

func (r *CertificateRepositoryImpl) FindByUserId(userId int) ([]model.Certificate, error) {
	var certificates []model.Certificate

	err := r.DB.
		Preload("User").
		Preload("Training").
		Where("user_id = ?", userId).
		Find(&certificates).
		Error

	return certificates, err
}

func (r *CertificateRepositoryImpl) Delete(id int) error {
	result := r.DB.Delete(&model.Certificate{}, id)
	return result.Error
}


func (r *CertificateRepositoryImpl) UpdateStatus(
	id int,
	status model.CertificateStatus,
) error {
	result := r.DB.Model(&model.Certificate{}).
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

	var certs []model.Certificate
	var total int64

	r.DB.Model(&model.Certificate{}).
		Where("status = ?", model.CertPending).
		Count(&total)

	err := r.DB.
		Preload("User").
		Where("status = ?", model.CertPending).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&certs).
		Error

	return certs, total, err
}