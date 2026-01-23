package repository

import "training-plan-api/model"

type DepartmentRepository interface {
	Save(department model.Department) error
	FindById(departmentId int) (model.Department, error)
	FindByIdWithStaffCount(departmentId int) (DepartmentStaffCount, error)
	FindAll() ([]model.Department, error)
	Update(department model.Department) error
	Delete(departmentId int) error
	FindAllPaginated(offset, limit int) ([]DepartmentStaffCount, int64, error)
}
