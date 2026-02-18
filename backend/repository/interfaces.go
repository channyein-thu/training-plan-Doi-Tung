package repository

import (
	"training-plan-api/data/request"
	"training-plan-api/model"
)

type DepartmentStaffCount struct {
	ID         int
	Name       string
	Division   model.Division
	TotalStaff int64
}

type CourseRepository interface {
	Save(course *model.Course) error
	FindById(id int) (*model.Course, error)
	FindPaginated(offset, limit int) ([]model.Course, int64, error)
	Update(course *model.Course) error
	Delete(id int) error
}

type DepartmentRepository interface {
	Save(department *model.Department) error
	FindById(departmentId int) (*model.Department, error)
	FindByIdWithStaffCount(departmentId int) (DepartmentStaffCount, error)
	FindDepartmentList() ([]model.Department, error)
	Update(department *model.Department) error
	Delete(departmentId int) error
	FindAllPaginated(offset, limit int) ([]DepartmentStaffCount, int64, error)
}

type CertificateRepository interface {
	Save(certificate *model.Certificate) error
	FindById(id int) (*model.Certificate, error)
	FindByUserId(userId int) ([]model.Certificate, error)
	Delete(id int) error
	FindAllPending(offset, limit int) ([]model.Certificate, int64, error)
	UpdateStatus(id int, status model.CertificateStatus) error
	FindRecordByIDAndUserID(recordID int, userID uint) (*model.Record, error)
}

type UserRepository interface {
	Save(user *model.User) error
	Update(user *model.User) error
	Delete(userId uint) error
	FindById(userId uint) (*model.User, error)
	FindByIdWithDepartment(userId uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByEmployeeID(employeeID string) (*model.User, error)
	FindAllPaginated(offset, limit int) ([]model.User, int64, error)
	FindByDepartmentPaginated(departmentID, offset, limit int) ([]model.User, int64, error)
	ExistsByEmail(email string) bool
	ExistsByEmployeeID(employeeID string) bool
	FindAllWithFilters(params request.UserTableQueryParams) ([]model.User, int64, error)
}

type RecordRepository interface {
	Save(record *model.Record) error
	FindById(id int) (*model.Record, error)
	Update(record *model.Record) error
	Delete(id int) error
	Exists(userId uint, courseId uint) bool
	FindByManagerDepartment(departmetnID int, offset, limit int) ([]model.Record, int64, error)
	FindByUserId(userID uint, offset, limit int) ([]model.Record, int64, error)
}
