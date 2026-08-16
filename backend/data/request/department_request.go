package request

import "training-plan-api/model"

type CreateDepartmentRequest struct {
	Name     string        `json:"name" validate:"required,min=2"`
	Division model.Division `json:"division" validate:"required"`
}

type UpdateDepartmentRequest struct {
	Name     string        `json:"name" validate:"required,min=2"`
	Division model.Division `json:"division" validate:"required"`
}