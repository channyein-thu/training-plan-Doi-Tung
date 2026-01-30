package service

import (
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
