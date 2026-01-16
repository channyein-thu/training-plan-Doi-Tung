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