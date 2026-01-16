package service

import (
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
)

type DepartmentServiceImpl struct {
	repo     repository.DepartmentRepository
	validate *validator.Validate
}

// Create implements DepartmentService.
func (d *DepartmentServiceImpl) Create(req request.CreateDepartmentRequest) error {
	if err := d.validate.Struct(req); err != nil {
		return err
	}

	department := model.Department{
		Name: req.Name,
	}

	return d.repo.Save(department)
}


// Delete implements DepartmentService.
func (d *DepartmentServiceImpl) Delete(departmentId int) error {
	return d.repo.Delete(departmentId)
}

// FindAll implements DepartmentService.
func (d *DepartmentServiceImpl) FindAll() ([]response.DepartmentResponse, error) {
	departments, err := d.repo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]response.DepartmentResponse, len(departments))
	for i, department := range departments {
		responses[i] = response.DepartmentResponse{
			ID:   department.ID,
			Name: department.Name,
		}
	}

	return responses, nil
}

// FindById implements DepartmentService.
func (d *DepartmentServiceImpl) FindById(departmentId int) (response.DepartmentResponse, error) {
	department, err := d.repo.FindById(departmentId)
	if err != nil {
		return response.DepartmentResponse{}, err
	}

	return response.DepartmentResponse{
		ID:   department.ID,
		Name: department.Name,
	}, nil
}


// Update implements DepartmentService.
func (d *DepartmentServiceImpl) Update(departmentId int, req request.UpdateDepartmentRequest) error {
	if err := d.validate.Struct(req); err != nil {
		return err
	}

	department, err := d.repo.FindById(departmentId)
	if err != nil {
		return err
	}

	department.Name = req.Name

	return d.repo.Update(department)
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
