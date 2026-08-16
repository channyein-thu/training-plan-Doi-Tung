package request

import "training-plan-api/model"


type CreateUserRequest struct {
	Name         string           `json:"name" validate:"required,min=2,max=52"`
	EmployeeID   string           `json:"employeeID" validate:"required,min=1,max=52"`
	Email        string           `json:"email" validate:"required,email,max=52"`
	Phone        string           `json:"phone" validate:"required,max=20"`
	DepartmentID int              `json:"departmentId" validate:"required,gt=0"`
	Role         model.Role       `json:"role" validate:"required" oneof:"Hr(admin) DepartmentHead(manager) Staff"`
	Position     string           `json:"position" validate:"required,min=1,max=100"`
	Status       model.UserStatus `json:"status" validate:"required,oneof=Active Inactive Suspended"`
	Password     string           `json:"password" validate:"required,min=6"`
	WorkStartDate string          `json:"workStartDate" validate:"required"`
}

type UpdateUserRequest struct {
	Name         string           `json:"name" validate:"required,min=2,max=52"`
	EmployeeID   string           `json:"employeeID" validate:"required,min=1,max=52"`
	Email        string           `json:"email" validate:"required,email,max=52"`
	Phone        string           `json:"phone" validate:"required,max=20"`
	DepartmentID int              `json:"departmentId" validate:"required,gt=0"`
	Role         model.Role       `json:"role" validate:"required" oneof:"Hr(admin) DepartmentHead(manager) Staff"`
	Position     string           `json:"position" validate:"required,min=1,max=100"`
	Status       model.UserStatus `json:"status" validate:"required,oneof=Active Inactive Suspended"`
}

type ManagerCreateUserRequest struct {
	Name       string           `json:"name" validate:"required,min=2,max=52"`
	EmployeeID string           `json:"employeeID" validate:"required,min=1,max=52"`
	Email      string           `json:"email" validate:"required,email,max=52"`
	Phone      string           `json:"phone" validate:"required,max=20"`
	Position   string           `json:"position" validate:"required,min=1,max=100"`
	Status     model.UserStatus `json:"status" validate:"required,oneof=Active Inactive Suspended"`
	Password   string           `json:"password" validate:"required,min=6"`
	WorkStartDate string        `json:"workStartDate" validate:"required"`
}

// ManagerUpdateUserRequest is the subset of user fields a department manager may
// edit on one of their own staff. Department and role are intentionally not
// editable here (a manager only manages their own department, and staff stay
// staff); password changes are also out of scope.
type ManagerUpdateUserRequest struct {
	Name       string           `json:"name" validate:"required,min=2,max=52"`
	EmployeeID string           `json:"employeeID" validate:"required,min=1,max=52"`
	Email      string           `json:"email" validate:"required,email,max=52"`
	Phone      string           `json:"phone" validate:"required,max=20"`
	Position   string           `json:"position" validate:"required,min=1,max=100"`
	Status     model.UserStatus `json:"status" validate:"required,oneof=Active Inactive Suspended"`
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

type CompleteProfileRequest struct {
	EmployeeID   string `json:"employeeId" validate:"required,min=1,max=52"`
	DepartmentID int    `json:"departmentId" validate:"required,gt=0"`
	Phone        string `json:"phone" validate:"required,max=20"`
	Position     string `json:"position" validate:"max=100"`
}
