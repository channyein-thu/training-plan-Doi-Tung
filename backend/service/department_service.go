package service

import (
	"fmt"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
)

type DepartmentServiceImpl struct {
	repo     repository.DepartmentRepository
	validate *validator.Validate
}

func NewDepartmentServiceImpl(
	repo repository.DepartmentRepository,
	validate *validator.Validate,
) DepartmentService {
	return &DepartmentServiceImpl{
		repo:     repo,
		validate: validate,
	}
}

// Create implements DepartmentService.
func (d *DepartmentServiceImpl) Create(req request.CreateDepartmentRequest) error {
	if err := d.validate.Struct(req); err != nil {
		return helper.ValidationError(
			helper.FormatValidationError(err),
		)
	}

	department := &model.Department{
		Name: req.Name,
		Division: req.Division,
	}

	err := d.repo.Save(department)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements DepartmentService.
func (d *DepartmentServiceImpl) Delete(departmentId int) error {
	err := d.repo.Delete(departmentId)
	if err != nil {
		return err
	}
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

	return resp, nil
}

// FindById implements DepartmentService.
func (d *DepartmentServiceImpl) FindById(departmentId int) (response.DepartmentResponse, error) {

	department, err := d.repo.FindByIdWithStaffCount(departmentId)
	if err != nil {
		return response.DepartmentResponse{}, err
	}

	// map staffs
	staffs := make([]response.UserListResponse, 0, len(department.Staffs))

	for _, u := range department.Staffs {
		staffs = append(staffs, response.UserListResponse{
			ID:         u.ID,
			Name:       u.Name,
			EmployeeID: u.EmployeeID,
			Position:   u.Position,
		})
	}

	resp := response.DepartmentResponse{
		ID:         department.ID,
		Name:       department.Name,
		Division:   department.Division,
		TotalStaff: department.TotalStaff,
		Staffs:     staffs,
	}

	return resp, nil
}

func (d *DepartmentServiceImpl) FindDepartmentList() ([]response.DepartmentListItem, error) {
	result := []response.DepartmentListItem{}

	departments, err := d.repo.FindDepartmentList()
	if err != nil {
		return result, err
	}

	for _, dept := range departments {
		result = append(result, response.DepartmentListItem{
			ID:   dept.ID,
			Name: fmt.Sprintf("%s (%s)", dept.Name, dept.Division),
		})
	}

	return result, nil
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
	return nil
}
