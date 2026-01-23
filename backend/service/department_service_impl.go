package service

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"training-plan-api/config"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type DepartmentServiceImpl struct {
	repo     repository.DepartmentRepository
	validate *validator.Validate
	cache     *redis.Client
}
func NewDepartmentServiceImpl(
	repo repository.DepartmentRepository,
	validate *validator.Validate,
	cache *redis.Client,
) DepartmentService {
	return &DepartmentServiceImpl{
		repo:     repo,
		validate: validate,
		cache:    cache,
	}
}


// Create implements DepartmentService.
func (d *DepartmentServiceImpl) Create(req request.CreateDepartmentRequest) error {
	if err := d.validate.Struct(req); err != nil {
			return helper.ValidationError(
		helper.FormatValidationError(err),
	)
	}

	department := model.Department{
		Name: req.Name,
		Division: req.Division,
	}

	err := d.repo.Save(department)
	if err != nil {
		return err
	}
	d.invalidateDepartmentCache()
	return nil
}


// Delete implements DepartmentService.
func (d *DepartmentServiceImpl) Delete(departmentId int) error {
	err := d.repo.Delete(departmentId)
	if err != nil {
		return err
	}
	d.invalidateDepartmentCache()
	return nil
}

// FindAll implements DepartmentService.
func (d *DepartmentServiceImpl) FindPaginated(page, pageSize int) (response.PaginatedResponse[response.DepartmentResponse], error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	cacheKey := fmt.Sprintf("department:page:%d:size:%d", page, pageSize)


	// CACHE HIT
	cached, err := d.cache.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		var resp response.PaginatedResponse[response.DepartmentResponse]
		_ = json.Unmarshal([]byte(cached), &resp)
		log.Println("CACHE HIT:", cacheKey)
		return resp, nil
	} else if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	// CACHE MISS
	offset := (page - 1) * pageSize
	departments, total, err := d.repo.FindAllPaginated(offset, pageSize)

	if err != nil {
		return response.PaginatedResponse[response.DepartmentResponse]{}, err
	}

	var items []response.DepartmentResponse
	for _, dept := range departments {
		items = append(items, response.DepartmentResponse{
			ID:         dept.ID,
			Name:       dept.Name,
			Division:   dept.Division,
			TotalStaff: dept.TotalStaff,
		})
	}

	resp := response.PaginatedResponse[response.DepartmentResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      pageSize,
			TotalItems: total,
			TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		},
	}

	bytes, _ := json.Marshal(resp)
	_ = d.cache.Set(config.Ctx, cacheKey, bytes, time.Minute*10).Err()

	return resp, nil
}

// FindById implements DepartmentService.
func (d *DepartmentServiceImpl) FindById(departmentId int) (response.DepartmentResponse, error) {
	cacheKey:= fmt.Sprintf("department:id:%d", departmentId)

	// CACHE HIT
	cached, err := d.cache.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		var resp response.DepartmentResponse
		_ = json.Unmarshal([]byte(cached), &resp)
		log.Println("CACHE HIT:", cacheKey)
		return resp, nil
	} else if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	// CACHE MISS
	department, err := d.repo.FindByIdWithStaffCount(departmentId)
	if err != nil {
		return response.DepartmentResponse{}, err
	}
	resp := response.DepartmentResponse{
		ID:         department.ID,
		Name:       department.Name,
		Division:   department.Division,
		TotalStaff: department.TotalStaff,
	}

	bytes, _ := json.Marshal(resp)
	_ = d.cache.Set(config.Ctx, cacheKey, bytes, time.Minute*10).Err()

	return resp, nil
}


// Update implements DepartmentService.
func (d *DepartmentServiceImpl) Update(departmentId int, req request.UpdateDepartmentRequest) error {
	if err := d.validate.Struct(req); err != nil {
			return helper.ValidationError(
		helper.FormatValidationError(err),
	)
	}

	department, err := d.repo.FindById(departmentId)
	if err != nil {
		return err
	}

	department.Name = req.Name
	department.Division = req.Division

	err = d.repo.Update(department)
	if err != nil {
		return err
	}
	d.invalidateDepartmentCache()
	return nil
}


// CACHE INVALIDATION
func (s *DepartmentServiceImpl) invalidateDepartmentCache() {
	iter := s.cache.Scan(config.Ctx, 0, "department:*", 0).Iterator()
	for iter.Next(config.Ctx) {
		s.cache.Del(config.Ctx, iter.Val())
	}
	log.Println("Department cache invalidated")
}
