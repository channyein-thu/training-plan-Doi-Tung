package service

import (
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"
	"training-plan-api/utils"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
	deptRepo repository.DepartmentRepository
	validate *validator.Validate
}

func NewUserServiceImpl(
	userRepo repository.UserRepository,
	deptRepo repository.DepartmentRepository,
	validate *validator.Validate,
) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
		deptRepo: deptRepo,
		validate: validate,
	}
}

func (s *UserServiceImpl) AdminCreate(req request.CreateUserRequest, creatorID uint) error {
	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	if !req.Role.IsValid() {
		return helper.BadRequest("Invalid system role specified")
	}

	if _, err := s.deptRepo.FindById(req.DepartmentID); err != nil {
		return helper.BadRequest("Invalid department selected")
	}

	if s.userRepo.ExistsByEmail(req.Email) {
		return helper.BadRequest("Email already registered")
	}
	if s.userRepo.ExistsByEmployeeID(req.EmployeeID) {
		return helper.BadRequest("Employee ID already registered")
	}

	user := model.User{
		Name:         req.Name,
		EmployeeID:   req.EmployeeID,
		Email:        req.Email,
		Phone:        req.Phone,
		DepartmentID: req.DepartmentID,
		Role:         req.Role,
		Position:     req.Position,
		Status:       req.Status,
		Password:     utils.GeneratePassword(req.Password),
		CreatedBy:    model.CreatedByAdmin,
		CreatedByID:  &creatorID,
	}

	return s.userRepo.Save(user)
}

func (s *UserServiceImpl) AdminUpdate(userID uint, req request.UpdateUserRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	if !req.Role.IsValid() {
		return helper.BadRequest("Invalid system role specified")
	}

	existingUser, err := s.userRepo.FindById(userID)
	if err != nil {
		return err
	}

	if _, err := s.deptRepo.FindById(req.DepartmentID); err != nil {
		return helper.BadRequest("Invalid department selected")
	}

	existingUser.Name = req.Name
	existingUser.EmployeeID = req.EmployeeID
	existingUser.Email = req.Email
	existingUser.Phone = req.Phone
	existingUser.DepartmentID = req.DepartmentID
	existingUser.Role = req.Role
	existingUser.Position = req.Position
	existingUser.Status = req.Status

	return s.userRepo.Update(existingUser)
}

func (s *UserServiceImpl) AdminDelete(userID uint) error {
	return s.userRepo.Delete(userID)
}

func (s *UserServiceImpl) AdminFindAll(page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	users, total, err := s.userRepo.FindAllPaginated(offset, pageSize)
	if err != nil {
		return response.PaginatedResponse[response.UserListResponse]{}, err
	}

	return response.PaginatedResponse[response.UserListResponse]{
		Items: response.ToUserListResponses(users),
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      pageSize,
			TotalItems: total,
			TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		},
	}, nil
}

func (s *UserServiceImpl) AdminFindById(userID uint) (response.UserResponse, error) {
	user, err := s.userRepo.FindByIdWithDepartment(userID)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(user), nil
}

func (s *UserServiceImpl) AdminFindAllForTable(params request.UserTableQueryParams) (response.PaginatedResponse[response.UserTableResponse], error) {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}
	if params.Limit > 100 {
		params.Limit = 100
	}

	users, total, err := s.userRepo.FindAllWithFilters(params)
	if err != nil {
		return response.PaginatedResponse[response.UserTableResponse]{}, err
	}

	return response.PaginatedResponse[response.UserTableResponse]{
		Items: response.ToUserTableResponses(users),
		Meta: response.PaginationMeta{
			Page:       params.Page,
			Limit:      params.Limit,
			TotalItems: total,
			TotalPages: int((total + int64(params.Limit) - 1) / int64(params.Limit)),
		},
	}, nil
}

func (s *UserServiceImpl) ManagerCreate(req request.ManagerCreateUserRequest, managerID uint, managerDepartmentID int) error {
	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	if s.userRepo.ExistsByEmail(req.Email) {
		return helper.BadRequest("Email already registered")
	}
	if s.userRepo.ExistsByEmployeeID(req.EmployeeID) {
		return helper.BadRequest("Employee ID already registered")
	}

	user := model.User{
		Name:         req.Name,
		EmployeeID:   req.EmployeeID,
		Email:        req.Email,
		Phone:        req.Phone,
		DepartmentID: managerDepartmentID,
		Role:         model.RoleStaff,
		Position:     req.Position,
		Status:       req.Status,
		Password:     utils.GeneratePassword(req.Password),
		CreatedBy:    model.CreatedByManager,
		CreatedByID:  &managerID,
	}

	return s.userRepo.Save(user)
}

func (s *UserServiceImpl) ManagerFindByDepartment(departmentID, page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	users, total, err := s.userRepo.FindByDepartmentPaginated(departmentID, offset, pageSize)
	if err != nil {
		return response.PaginatedResponse[response.UserListResponse]{}, err
	}

	return response.PaginatedResponse[response.UserListResponse]{
		Items: response.ToUserListResponses(users),
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      pageSize,
			TotalItems: total,
			TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		},
	}, nil
}
