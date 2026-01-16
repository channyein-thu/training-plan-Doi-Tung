package repository

import "training-plan-api/model"

type CourseRepository interface {
	Save(course *model.Course) error
	FindById(id int) (model.Course, error)
	FindPaginated(offset, limit int) ([]model.Course, int64, error)
	Update(course *model.Course) error
	Delete(id int) error
}