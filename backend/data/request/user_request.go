package request

import "training-plan-api/model"

type CreateUserRequest struct {
	Name         string           `json:"name" validate:"required,min=2,max=52"`
	EmployeeID   string           `json:"employeeID" validate:"required,min=1,max=52"`
	Email        string           `json:"email" validate:"required,email,max=52"`
	Phone        string           `json:"phone" validate:"max=20"`
	DepartmentID int              `json:"departmentId" validate:"required,gt=0"`
	Role         model.Role       `json:"role" validate:"required"`
	Position     string           `json:"position" validate:"required,min=1,max=100"`
	Status       model.UserStatus `json:"status" validate:"required,oneof=Active Inactive"`
	Password     string           `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Name         string           `json:"name" validate:"required,min=2,max=52"`
	EmployeeID   string           `json:"employeeID" validate:"required,min=1,max=52"`
	Email        string           `json:"email" validate:"required,email,max=52"`
	Phone        string           `json:"phone" validate:"max=20"`
	DepartmentID int              `json:"departmentId" validate:"required,gt=0"`
	Role         model.Role       `json:"role" validate:"required"`
	Position     string           `json:"position" validate:"required,min=1,max=100"`
	Status       model.UserStatus `json:"status" validate:"required,oneof=Active Inactive"`
}

type ManagerCreateUserRequest struct {
	Name       string           `json:"name" validate:"required,min=2,max=52"`
	EmployeeID string           `json:"employeeID" validate:"required,min=1,max=52"`
	Email      string           `json:"email" validate:"required,email,max=52"`
	Phone      string           `json:"phone" validate:"max=20"`
	Position   string           `json:"position" validate:"required,min=1,max=100"`
	Status     model.UserStatus `json:"status" validate:"required,oneof=Active Inactive"`
	Password   string           `json:"password" validate:"required,min=6"`
}

type UserTableQueryParams struct {
	Search       string `query:"search"`
	DepartmentID int    `query:"departmentId"`
	Status       string `query:"status"`
	Page         int    `query:"page"`
	Limit        int    `query:"limit"`
	SortBy       string `query:"sortBy"`
	SortOrder    string `query:"sortOrder"`
}
