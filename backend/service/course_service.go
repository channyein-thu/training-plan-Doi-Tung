package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"training-plan-api/config"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/mapper"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/calendar/v3"
)

type CourseServiceImpl struct {
	repo      repository.CourseRepository
	cache     *redis.Client
	validate  *validator.Validate
	calendar  *calendar.Service
	location  *time.Location
}

func NewCourseServiceImpl(
	repo repository.CourseRepository,
	cache *redis.Client,
	validate *validator.Validate,
	calendar *calendar.Service,
	location *time.Location,
) CourseService {
	return &CourseServiceImpl{
		repo:     repo,
		cache:    cache,
		validate: validate,
		calendar: calendar,
		location: location,
	}
}


// CREATE COURSE
func (s *CourseServiceImpl) Create(req request.CreateCourseRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(
		helper.FormatValidationError(err),
	)
	}

	course := mapper.ToCourseModel(req)

	if err := s.repo.Save(&course); err != nil {
		return err
	}

	// ===== SAFETY CHECKS =====
	if s.calendar == nil || s.location == nil {
		log.Println("Calendar not initialized, skipping calendar")
		s.invalidateCourseCache()
		return nil
	}

	if course.Date.IsZero() {
		log.Println("Course date is zero, skipping calendar")
		s.invalidateCourseCache()
		return nil
	}
	
	hours := 8
	if course.NumberOfHours != nil {
		hours = *course.NumberOfHours
	}

	// description := "Training Plan"
	// if course.Content != nil && *course.Content != "" {
	// 	description = *course.Content
	// }

	eventID, err := helper.CreateCourseCalendarEvent(
		context.Background(),
		s.calendar,
		course.Name,
		course.Content,
		course.Date.In(s.location),
		hours,
	)
	if err != nil {
		return err
	}

	if err == nil {
		course.CalendarEventID = &eventID
		if err := s.repo.Update(&course); err != nil {
			log.Println("Failed to save calendar_event_id:", err)
		}
	} else {
		log.Println("Calendar create failed:", err)
	}

	s.invalidateCourseCache()
	return nil
}

// DELETE COURSE
func (s *CourseServiceImpl) Delete(courseId int) error {
	course, err := s.repo.FindById(courseId)
	if err != nil {
		return err
	}

	// Delete calendar event if exists
	if course.CalendarEventID != nil {
		if err := helper.DeleteCourseCalendarEvent(
			context.Background(),
			s.calendar,
			*course.CalendarEventID,
		); err != nil {
			log.Println("Calendar delete failed:", err)
		}
	}

	if err := s.repo.Delete(courseId); err != nil {
		return err
	}

	s.invalidateCourseCache()
	return nil
}

// FIND BY ID (CACHE)
func (s *CourseServiceImpl) FindById(courseId int) (response.CourseResponse, error) {
	cacheKey := fmt.Sprintf("course:id:%d", courseId)

	// CACHE HIT
	cached, err := s.cache.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		var resp response.CourseResponse
		_ = json.Unmarshal([]byte(cached), &resp)
		log.Println("CACHE HIT:", cacheKey)
		return resp, nil
	} else if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	// CACHE MISS
	course, err := s.repo.FindById(courseId)
	if err != nil {
		return response.CourseResponse{}, err
	}

	resp := mapper.ToCourseResponse(*course)

	bytes, _ := json.Marshal(resp)
	_ = s.cache.Set(config.Ctx, cacheKey, bytes, time.Minute*10).Err()

	return resp, nil
}

// FIND PAGINATED (CACHE)
func (s *CourseServiceImpl) FindPaginated(
	page int,
	pageSize int,
) (response.PaginatedResponse[response.CourseResponse], error) {

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	cacheKey := fmt.Sprintf("course:page:%d:size:%d", page, pageSize)

	// CACHE HIT
	cached, err := s.cache.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		var resp response.PaginatedResponse[response.CourseResponse]
		_ = json.Unmarshal([]byte(cached), &resp)
		log.Println("CACHE HIT:", cacheKey)
		return resp, nil
	} else if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	// CACHE MISS
	offset := (page - 1) * pageSize
	courses, total, err := s.repo.FindPaginated(offset, pageSize)
	if err != nil {
		return response.PaginatedResponse[response.CourseResponse]{}, err
	}

	items := mapper.ToCourseResponseList(courses)

	resp := response.PaginatedResponse[response.CourseResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      pageSize,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
		},
	}

	bytes, _ := json.Marshal(resp)
	_ = s.cache.Set(config.Ctx, cacheKey, bytes, time.Minute*10).Err()

	return resp, nil
}

// UPDATE COURSE
func (s *CourseServiceImpl) Update(courseId int, req request.UpdateCourseRequest) error {
	if err := s.validate.Struct(req); err != nil {
			return helper.ValidationError(
		helper.FormatValidationError(err),
	)
	}

	course, err := s.repo.FindById(courseId)
	if err != nil {
		return err
	}

	mapper.UpdateCourseFromRequest(course, req)

	if err := s.repo.Update(course); err != nil {
		return err
	}

	// ===== SAFETY CHECKS =====
	if s.calendar == nil || s.location == nil {
		log.Println("Calendar not initialized, skipping calendar update")
		s.invalidateCourseCache()
		return nil
	}

	if course.CalendarEventID == nil {
		log.Println("No calendar_event_id, skipping calendar update")
		s.invalidateCourseCache()
		return nil
	}

	if course.Date.IsZero() {
		log.Println("Course date is zero, skipping calendar update")
		s.invalidateCourseCache()
		return nil
	}
	// ========================

	hours := 8
	if course.NumberOfHours != nil {
		hours = *course.NumberOfHours
	}

	// description := "Training  Plan"
	// if course.Content != nil && *course.Content != "" {
	// 	description = *course.Content
	// }

	if err := helper.UpdateCourseCalendarEvent(
		context.Background(),
		s.calendar,
		*course.CalendarEventID,
		course.Name,
		course.Content,
		course.Date.In(s.location),
		hours,
	); err != nil {
		log.Println("Calendar update failed:", err)
		return err
	}

	s.invalidateCourseCache()
	return nil
}

// CACHE INVALIDATION
func (s *CourseServiceImpl) invalidateCourseCache() {
	iter := s.cache.Scan(config.Ctx, 0, "course:*", 0).Iterator()
	for iter.Next(config.Ctx) {
		s.cache.Del(config.Ctx, iter.Val())
	}
	log.Println("Course cache invalidated")
}
