package repository

import (
	"training-plan-api/data/request"
	"training-plan-api/model"
)

type UserRepository interface {
	Save(user model.User) error
	Update(user model.User) error
	Delete(userId uint) error
	FindById(userId uint) (model.User, error)
	FindByIdWithDepartment(userId uint) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByEmployeeID(employeeID string) (model.User, error)
	FindAllPaginated(offset, limit int) ([]model.User, int64, error)
	FindByDepartmentPaginated(departmentID, offset, limit int) ([]model.User, int64, error)
	ExistsByEmail(email string) bool
	ExistsByEmployeeID(employeeID string) bool
	FindAllWithFilters(params request.UserTableQueryParams) ([]model.User, int64, error)
}
