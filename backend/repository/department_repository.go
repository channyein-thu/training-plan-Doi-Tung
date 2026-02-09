package repository

import (
	"errors"
	"strings"
	"training-plan-api/helper"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type DepartmentRepositoryImpl struct {
	Db *gorm.DB
}

func NewDepartmentRepositoryImpl(db *gorm.DB) DepartmentRepository {
	return &DepartmentRepositoryImpl{Db: db}
}

// -------------------- CRUD --------------------

func (r *DepartmentRepositoryImpl) Save(department *model.Department) error {
	err := r.Db.Create(department).Error
	if err != nil {
		// MySQL duplicate entry error (1062)
		if strings.Contains(err.Error(), "idx_dept_division") {
			return helper.BadRequest("department already exists in this division")
		}
		return helper.InternalServerError("failed to create department")
	}
	return nil
}

func (r *DepartmentRepositoryImpl) FindById(departmentId int) (*model.Department, error) {
	var result model.Department

	err := r.Db.First(&result, departmentId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("department not found")
		}
		return nil, err
	}

	return &result, nil
}

func (r *DepartmentRepositoryImpl) FindByIdWithStaffCount(departmentId int) (DepartmentStaffCount, error) {
	var result DepartmentStaffCount

	err := r.Db.
		Table("departments").
		Select(`
			departments.id,
			departments.name,
			departments.division,
			COUNT(users.id) AS total_staff
		`).
		Joins("LEFT JOIN users ON users.department_id = departments.id").
		Where("departments.id = ?", departmentId).
		Group("departments.id").
		Scan(&result).Error

	if err != nil {
		return result, err
	}

	if result.ID == 0 {
		return result, errors.New("department not found")
	}

	return result, nil
}

func (r *DepartmentRepositoryImpl) FindDepartmentList() ([]model.Department, error) {
	var departments []model.Department
	err := r.Db.Order("id ASC").Find(&departments).Error
	return departments, err
}

func (r *DepartmentRepositoryImpl) Update(department *model.Department) error {
	result := r.Db.
		Model(&model.Department{}).
		Where("id = ?", department.ID).
		Updates(map[string]interface{}{
			"name":     department.Name,
			"division": department.Division,
		})

	if result.Error != nil {
		// Handle duplicate department name per division
		if strings.Contains(result.Error.Error(), "idx_dept_division") {
			return helper.BadRequest("department already exists in this division")
		}
		return helper.InternalServerError("failed to update department")
	}

	if result.RowsAffected == 0 {
		return helper.NotFound("department not found")
	}

	return nil
}

func (r *DepartmentRepositoryImpl) Delete(departmentId int) error {
	result := r.Db.Delete(&model.Department{}, departmentId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("department not found")
	}

	return nil
}



func (r *DepartmentRepositoryImpl) FindAllPaginated(offset, limit int) ([]DepartmentStaffCount, int64, error) {
	var result []DepartmentStaffCount
	var total int64
	err := r.Db.Model(&model.Department{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.Db.
		Table("departments").
		Select(`
			departments.id,
			departments.name,
			departments.division,
			COUNT(users.id) AS total_staff
		`).
		Joins("LEFT JOIN users ON users.department_id = departments.id").
		Group("departments.id").
		Order("departments.id ASC").
		Offset(offset).
		Limit(limit).
		Scan(&result).Error

	return result, total, err
}

