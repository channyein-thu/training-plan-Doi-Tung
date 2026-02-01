package service

import (
	"training-plan-api/data/request"
	"training-plan-api/data/response"
)

type UserService interface {
	AdminCreate(req request.CreateUserRequest, creatorID uint) error
	AdminUpdate(userID uint, req request.UpdateUserRequest) error
	AdminDelete(userID uint) error
	AdminFindAll(page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error)
	AdminFindById(userID uint) (response.UserResponse, error)
	AdminFindAllForTable(params request.UserTableQueryParams) (response.PaginatedResponse[response.UserTableResponse], error)
	ManagerCreate(req request.ManagerCreateUserRequest, managerID uint, managerDepartmentID int) error
	ManagerFindByDepartment(departmentID, page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error)
}
