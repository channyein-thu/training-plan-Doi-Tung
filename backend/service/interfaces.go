package service

import (
	"mime/multipart"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
)

type CourseService interface {
	Create(course request.CreateCourseRequest) error
	Update(courseId int, course request.UpdateCourseRequest) error
	Delete(courseId int) error
	FindById(courseId int) (response.CourseResponse, error)
	FindPaginated(page, pageSize int) (response.PaginatedResponse[response.CourseResponse], error)
}

type DepartmentService interface {
	Create(request.CreateDepartmentRequest) error
	Update(int, request.UpdateDepartmentRequest) error
	Delete(int) error
	FindById(int) (response.DepartmentResponse, error)
	FindPaginated(page, pageSize int) (response.PaginatedResponse[response.DepartmentResponse], error)
	FindDepartmentList() ([]response.DepartmentListItem, error)
}

type UserService interface {
	AdminCreate(req request.CreateUserRequest, creatorID uint) error
	AdminUpdate(userID uint, req request.UpdateUserRequest) error
	AdminDelete(userID uint) error
	AdminFindAll(page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error)
	AdminFindById(userID uint) (response.UserResponse, error)
	AdminFindAllForTable(params request.UserTableQueryParams) (response.PaginatedResponse[response.UserTableResponse], error)
	ManagerCreate(req request.ManagerCreateUserRequest, managerID uint, managerDepartmentID int) error
	ManagerFindByDepartment(departmentID, page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error)
}

type CertificateService interface {
	FindByCurrentUser(userID uint) ([]response.CertificateResponse, error)
	Upload(userID uint, req request.CreateCertificateRequest, file *multipart.FileHeader) error
	Delete(certificateID int, userID uint) error
}