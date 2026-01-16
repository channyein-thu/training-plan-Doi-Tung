package service

import (
	"training-plan-api/data/request"
	"training-plan-api/data/response"
)

type DepartmentService interface {
	Create(request.CreateDepartmentRequest) error
	Update(int, request.UpdateDepartmentRequest) error
	Delete(int) error
	FindById(int) (response.DepartmentResponse, error)
	FindAll() ([]response.DepartmentResponse, error)
}
