package service

import (
	"context"
	"log"
	"math"
	"time"

	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/mapper"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
	"google.golang.org/api/calendar/v3"
)

type TrainingPlanServiceImpl struct {
	repo      repository.TrainingPlanRepository
	validate  *validator.Validate
	calendar  *calendar.Service
	location  *time.Location
}

func NewTrainingPlanServiceImpl(
	repo repository.TrainingPlanRepository,
	validate *validator.Validate,
	calendar *calendar.Service,
	location *time.Location,
) TrainingPlanService {
	return &TrainingPlanServiceImpl{
		repo:     repo,
		validate: validate,
		calendar: calendar,
		location: location,
	}
}


// CREATE TRAINING PLAN
func (s *TrainingPlanServiceImpl) Create(req request.CreateTrainingPlanRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(
		helper.FormatValidationError(err),
	)
	}

	trainingPlan := mapper.ToTrainingPlanModel(req)

	if err := s.repo.Save(&trainingPlan); err != nil {
		return err
	}

	// ===== SAFETY CHECKS =====
	if s.calendar == nil || s.location == nil {
		log.Println("Calendar not initialized, skipping calendar")
		return nil
	}

	if trainingPlan.Date.IsZero() {
		log.Println("Training plan date is zero, skipping calendar")
		return nil
	}
	
	hours := 8
	if trainingPlan.NumberOfHours != nil {
		hours = *trainingPlan.NumberOfHours
	}

	// description := "Training Plan"
	// if trainingPlan.Content != nil && *trainingPlan.Content != "" {
	// 	description = *trainingPlan.Content
	// }

	eventID, err := helper.CreateTrainingPlanCalendarEvent(
		context.Background(),
		s.calendar,
		trainingPlan.Name,
		trainingPlan.Content,
		trainingPlan.Date.In(s.location),
		hours,
	)
	if err != nil {
		return err
	}

	if err == nil {
		trainingPlan.CalendarEventID = &eventID
		if err := s.repo.Update(&trainingPlan); err != nil {
			log.Println("Failed to save calendar_event_id:", err)
		}
	} else {
		log.Println("Calendar create failed:", err)
	}

	return nil
}

// DELETE TRAINING PLAN
func (s *TrainingPlanServiceImpl) Delete(trainingPlanId int) error {
	trainingPlan, err := s.repo.FindById(trainingPlanId)
	if err != nil {
		return err
	}

	// Delete calendar event if exists
	if trainingPlan.CalendarEventID != nil {
		if err := helper.DeleteTrainingPlanCalendarEvent(
			context.Background(),
			s.calendar,
			*trainingPlan.CalendarEventID,
		); err != nil {
			log.Println("Calendar delete failed:", err)
		}
	}

	if err := s.repo.Delete(trainingPlanId); err != nil {
		return err
	}

	return nil
}

// FIND BY ID (CACHE)
func (s *TrainingPlanServiceImpl) FindById(trainingPlanId int) (response.TrainingPlanResponse, error) {
	trainingPlan, err := s.repo.FindById(trainingPlanId)
	if err != nil {
		return response.TrainingPlanResponse{}, err
	}

	resp := mapper.ToTrainingPlanResponse(*trainingPlan)

	return resp, nil
}

// FIND PAGINATED (CACHE)
func (s *TrainingPlanServiceImpl) FindPaginated(
	page int,
	pageSize int,
) (response.PaginatedResponse[response.TrainingPlanResponse], error) {

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	trainingPlans, total, err := s.repo.FindPaginated(offset, pageSize)
	if err != nil {
		return response.PaginatedResponse[response.TrainingPlanResponse]{}, err
	}

	items := mapper.ToTrainingPlanResponseList(trainingPlans)

	resp := response.PaginatedResponse[response.TrainingPlanResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      pageSize,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
		},
	}

	return resp, nil
}

// UPDATE TRAINING PLAN
func (s *TrainingPlanServiceImpl) Update(trainingPlanId int, req request.UpdateTrainingPlanRequest) error {
	if err := s.validate.Struct(req); err != nil {
			return helper.ValidationError(
		helper.FormatValidationError(err),
	)
	}

	trainingPlan, err := s.repo.FindById(trainingPlanId)
	if err != nil {
		return err
	}

	mapper.UpdateTrainingPlanFromRequest(trainingPlan, req)

	if err := s.repo.Update(trainingPlan); err != nil {
		return err
	}

	// ===== SAFETY CHECKS =====
	if s.calendar == nil || s.location == nil {
		log.Println("Calendar not initialized, skipping calendar update")
		return nil
	}

	if trainingPlan.CalendarEventID == nil {
		log.Println("No calendar_event_id, skipping calendar update")
		return nil
	}

	if trainingPlan.Date.IsZero() {
		log.Println("Training plan date is zero, skipping calendar update")
		return nil
	}
	// ========================

	hours := 8
	if trainingPlan.NumberOfHours != nil {
		hours = *trainingPlan.NumberOfHours
	}

	// description := "Training  Plan"
	// if trainingPlan.Content != nil && *trainingPlan.Content != "" {
	// 	description = *trainingPlan.Content
	// }

	if err := helper.UpdateTrainingPlanCalendarEvent(
		context.Background(),
		s.calendar,
		*trainingPlan.CalendarEventID,
		trainingPlan.Name,
		trainingPlan.Content,
		trainingPlan.Date.In(s.location),
		hours,
	); err != nil {
		log.Println("Calendar update failed:", err)
		return err
	}

	return nil
}

