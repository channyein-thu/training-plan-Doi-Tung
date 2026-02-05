package repository

import (
	"fmt"
	"strings"
	"training-plan-api/data/request"
	"training-plan-api/helper"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

func (r *UserRepositoryImpl) Save(user model.User) error {
	err := r.Db.Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "email") {
				return helper.BadRequest("Email already registered")
			}
			if strings.Contains(err.Error(), "employee_id") {
				return helper.BadRequest("Employee ID already registered")
			}
			return helper.BadRequest("User already exists")
		}
		return helper.InternalServerError("Failed to create user")
	}
	return nil
}

func (r *UserRepositoryImpl) Update(user model.User) error {
	result := r.Db.Model(&model.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":          user.Name,
		"email":         user.Email,
		"employee_id":   user.EmployeeID,
		"phone":         user.Phone,
		"department_id": user.DepartmentID,
		"role":          user.Role,
		"status":        user.Status,
		"position":      user.Position,
	})
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			if strings.Contains(result.Error.Error(), "email") {
				return helper.BadRequest("Email already in use")
			}
			if strings.Contains(result.Error.Error(), "employee_id") {
				return helper.BadRequest("Employee ID already in use")
			}
		}
		return helper.InternalServerError("Failed to update user")
	}
	if result.RowsAffected == 0 {
		return helper.NotFound("User not found")
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(userId uint) error {
	result := r.Db.Delete(&model.User{}, userId)
	if result.Error != nil {
		return helper.InternalServerError("Failed to delete user")
	}
	if result.RowsAffected == 0 {
		return helper.NotFound("User not found")
	}
	return nil
}

func (r *UserRepositoryImpl) FindById(userId uint) (model.User, error) {
	var user model.User
	err := r.Db.First(&user, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, helper.NotFound("User not found")
		}
		return user, helper.InternalServerError("Failed to find user")
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByIdWithDepartment(userId uint) (model.User, error) {
	var user model.User
	err := r.Db.Preload("Department").First(&user, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, helper.NotFound("User not found")
		}
		return user, helper.InternalServerError("Failed to find user")
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, helper.NotFound("User not found")
		}
		return user, helper.InternalServerError("Failed to find user")
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmployeeID(employeeID string) (model.User, error) {
	var user model.User
	err := r.Db.Where("employee_id = ?", employeeID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, helper.NotFound("User not found")
		}
		return user, helper.InternalServerError("Failed to find user")
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindAllPaginated(offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	r.Db.Model(&model.User{}).Count(&total)

	err := r.Db.Preload("Department").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&users).Error
	if err != nil {
		return nil, 0, helper.InternalServerError("Failed to fetch users")
	}
	return users, total, nil
}

func (r *UserRepositoryImpl) FindByDepartmentPaginated(departmentID, offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	r.Db.Model(&model.User{}).Where("department_id = ?", departmentID).Count(&total)

	err := r.Db.Preload("Department").
		Where("department_id = ?", departmentID).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&users).Error
	if err != nil {
		return nil, 0, helper.InternalServerError("Failed to fetch users")
	}
	return users, total, nil
}

func (r *UserRepositoryImpl) ExistsByEmail(email string) bool {
	var count int64
	r.Db.Model(&model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *UserRepositoryImpl) ExistsByEmployeeID(employeeID string) bool {
	var count int64
	r.Db.Model(&model.User{}).Where("employee_id = ?", employeeID).Count(&count)
	return count > 0
}

func (r *UserRepositoryImpl) FindAllWithFilters(params request.UserTableQueryParams) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.Db.Model(&model.User{})

	if params.Search != "" {
		searchTerm := "%" + strings.ToLower(params.Search) + "%"
		query = query.Where(
			"LOWER(employee_id) LIKE ? OR LOWER(name) LIKE ? OR LOWER(email) LIKE ?",
			searchTerm, searchTerm, searchTerm,
		)
	}

	if params.DepartmentID > 0 {
		query = query.Where("department_id = ?", params.DepartmentID)
	}

	if params.Status != "" {
		status := params.Status
		if strings.ToUpper(status) == "ACTIVE" {
			status = "Active"
		} else if strings.ToUpper(status) == "INACTIVE" {
			status = "Inactive"
		}
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, helper.InternalServerError("Failed to count users")
	}

	sortField := "employee_id"
	sortOrder := "ASC"

	allowedSortFields := map[string]string{
		"employee_id": "employee_id",
		"employeeId":  "employee_id",
		"name":        "name",
		"email":       "email",
		"department":  "department_id",
		"status":      "status",
		"created_at":  "created_at",
	}

	if params.SortBy != "" {
		if field, ok := allowedSortFields[params.SortBy]; ok {
			sortField = field
		}
	}

	if strings.ToLower(params.SortOrder) == "desc" {
		sortOrder = "DESC"
	}

	orderClause := fmt.Sprintf("%s %s", sortField, sortOrder)

	page := params.Page
	limit := params.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	err := query.
		Preload("Department").
		Order(orderClause).
		Offset(offset).
		Limit(limit).
		Find(&users).Error

	if err != nil {
		return nil, 0, helper.InternalServerError("Failed to fetch users")
	}

	return users, total, nil
}
