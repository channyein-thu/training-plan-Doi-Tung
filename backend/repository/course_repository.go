package repository

import (
	"errors"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	Db *gorm.DB
}

func NewCourseRepositoryImpl(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{Db: db}
}


// FindPaginated implements CourseRepository.
func (r *CourseRepositoryImpl) FindPaginated(offset int, limit int) ([]model.Course, int64, error) {
	var courses []model.Course
	var total int64

	if err := r.Db.Model(&model.Course{}).Count(&total).Error; err != nil {
	return nil, 0, err
}

	err := r.Db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&courses).Error
	return courses, total, err
}

// Delete implements CourseRepository.
func (r *CourseRepositoryImpl) Delete(courseId int) error {
	result := r.Db.Delete(&model.Course{}, courseId)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("course not found")
	}

	return nil
}

// FindAll implements CourseRepository.
func (r *CourseRepositoryImpl) FindAll() ([]model.Course, error) {
	var courses []model.Course

	result := r.Db.Find(&courses)
	return courses, result.Error
}

// FindById implements CourseRepository.
func (r *CourseRepositoryImpl) FindById(courseId int) (*model.Course, error) {
	var course model.Course

	result := r.Db.First(&course, courseId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("course not found")
		}
		return nil, result.Error
	}

	return &course, nil
}

// Save implements CourseRepository.
func (r *CourseRepositoryImpl) Save(course *model.Course) error {
	result := r.Db.Create(course)
	return result.Error
}

// Update implements CourseRepository.
func (r *CourseRepositoryImpl) Update(course *model.Course) error {
	result := r.Db.
		Model(&model.Course{}).
		Where("id = ?", course.ID).
		Updates(course)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("course not found")
	}

	return nil
}
