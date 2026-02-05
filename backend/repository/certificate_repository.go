package repository

import (
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
	err := r.DB.Where("user_id = ?", userId).Find(&certificates).Error
	return certificates, err
}
func (r *CertificateRepositoryImpl) Delete(id int) error {
	result := r.DB.Delete(&model.Certificate{}, id)
	return result.Error
}