package repository

import (
	"errors"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type DepartmentRepositoryImpl struct {
	Db *gorm.DB
}

func NewDepartmentRepositoryImpl(db *gorm.DB) DepartmentRepository {
	return &DepartmentRepositoryImpl{Db: db}
}

func (repository *DepartmentRepositoryImpl) Save(department model.Department) error {
	result := repository.Db.Create(&department)
	return result.Error
}

func (repository *DepartmentRepositoryImpl) FindById(departmentId int) (model.Department, error) {
	var department model.Department

	result := repository.Db.First(&department, departmentId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return department, errors.New("department not found")
		}
		return department, result.Error
	}

	return department, nil
}

func (repository *DepartmentRepositoryImpl) FindAll() ([]model.Department, error) {
	var departments []model.Department

	result := repository.Db.Find(&departments)
	return departments, result.Error
}

func (repository *DepartmentRepositoryImpl) Update(department model.Department) error {
	result := repository.Db.
		Model(&model.Department{}).
		Where("id = ?", department.ID).
		Updates(map[string]interface{}{
			"name": department.Name,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("department not found")
	}

	return nil
}

func (repository *DepartmentRepositoryImpl) Delete(departmentId int) error {
	result := repository.Db.Delete(&model.Department{}, departmentId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("department not found")
	}

	return nil
}
