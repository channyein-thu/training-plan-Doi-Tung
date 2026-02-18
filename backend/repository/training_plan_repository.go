package repository

import (
	"errors"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type TrainingPlanRepositoryImpl struct {
	Db *gorm.DB
}

func NewTrainingPlanRepositoryImpl(db *gorm.DB) TrainingPlanRepository {
	return &TrainingPlanRepositoryImpl{Db: db}
}

// FindPaginated implements TrainingPlanRepository.
func (r *TrainingPlanRepositoryImpl) FindPaginated(offset int, limit int) ([]model.TrainingPlan, int64, error) {
	var trainingPlans []model.TrainingPlan
	var total int64

	if err := r.Db.Model(&model.TrainingPlan{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.Db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&trainingPlans).Error
	return trainingPlans, total, err
}

// Delete implements TrainingPlanRepository.
func (r *TrainingPlanRepositoryImpl) Delete(trainingPlanId int) error {
	result := r.Db.Delete(&model.TrainingPlan{}, trainingPlanId)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("training plan not found")
	}

	return nil
}

// FindAll implements TrainingPlanRepository.
func (r *TrainingPlanRepositoryImpl) FindAll() ([]model.TrainingPlan, error) {
	var trainingPlans []model.TrainingPlan

	result := r.Db.Find(&trainingPlans)
	return trainingPlans, result.Error
}

// FindById implements TrainingPlanRepository.
func (r *TrainingPlanRepositoryImpl) FindById(trainingPlanId int) (*model.TrainingPlan, error) {
	var trainingPlan model.TrainingPlan

	result := r.Db.First(&trainingPlan, trainingPlanId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("training plan not found")
		}
		return nil, result.Error
	}

	return &trainingPlan, nil
}

// Save implements TrainingPlanRepository.
func (r *TrainingPlanRepositoryImpl) Save(trainingPlan *model.TrainingPlan) error {
	result := r.Db.Create(trainingPlan)
	return result.Error
}

// Update implements TrainingPlanRepository.
func (r *TrainingPlanRepositoryImpl) Update(trainingPlan *model.TrainingPlan) error {
	result := r.Db.
		Model(&model.TrainingPlan{}).
		Where("id = ?", trainingPlan.ID).
		Updates(trainingPlan)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("training plan not found")
	}

	return nil
}
