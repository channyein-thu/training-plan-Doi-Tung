package repository

import "training-plan-api/model"

type DepartmentStaffCount struct {
	ID         int
	Name       string
	Division   model.Division
	TotalStaff int64
}

type CourseRepository interface {
	Save(course *model.Course) error
	FindById(id int) (model.Course, error)
	FindPaginated(offset, limit int) ([]model.Course, int64, error)
	Update(course *model.Course) error
	Delete(id int) error
}

type DepartmentRepository interface {
	Save(department model.Department) error
	FindById(departmentId int) (model.Department, error)
	FindByIdWithStaffCount(departmentId int) (DepartmentStaffCount, error)
	FindDepartmentList() ([]model.Department, error)
	Update(department model.Department) error
	Delete(departmentId int) error
	FindAllPaginated(offset, limit int) ([]DepartmentStaffCount, int64, error)
}
