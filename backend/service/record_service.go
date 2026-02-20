package service

import (
	"math"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/xuri/excelize/v2"
)

type RecordServiceImpl struct {
	repo     repository.RecordRepository
	userRepo repository.UserRepository
	validate *validator.Validate
}

func NewRecordServiceImpl(
	repo repository.RecordRepository,
	userRepo repository.UserRepository,
	validate *validator.Validate,
) RecordService {
	return &RecordServiceImpl{
		repo:     repo,
		userRepo: userRepo,
		validate: validate,
	}
}

// Search implements RecordService.
func (s *RecordServiceImpl) Search(
	req request.RecordFilterRequest,
) (response.PaginatedResponse[response.AdminRecordResponse], error) {

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 10
	}

	records, total, err := s.repo.Search(req)
	if err != nil {
		return response.PaginatedResponse[response.AdminRecordResponse]{}, err
	}

	items := make([]response.AdminRecordResponse, 0, len(records))

	for _, r := range records {

		var trainingName string
		var location *string
		var costPerPerson *int
		var budgetCode *string

		if r.TrainingPlan != nil {
			trainingName = r.TrainingPlan.Name
			location = r.TrainingPlan.Location
			costPerPerson = r.TrainingPlan.CostPerPerson
			budgetCode = r.TrainingPlan.BudgetCode
		}

		var employeeName string
		var position string
		var department string
		var division string

		if r.User != nil {
			employeeName = r.User.Name
			position = r.User.Position

			if r.User.Department != nil {
				department = r.User.Department.Name
				division = string(r.User.Department.Division)
			}
		}

		items = append(items, response.AdminRecordResponse{
			ID:               r.ID,
			TrainingPlanName: trainingName,
			Location:         location,
			CostPerPerson:    costPerPerson,
			BudgetCode:       budgetCode,
			EmployeeID:       r.User.EmployeeID,
			EmployeeName:     employeeName,
			Position:         position,
			Department:       department,
			Division:         division,
			Status:           string(r.Status),
			UpdatedAt:        r.UpdatedAt,
		})
	}

	return response.PaginatedResponse[response.AdminRecordResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       req.Page,
			Limit:      req.Limit,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(req.Limit))),
		},
	}, nil
}


// FindByUser implements RecordService.
func (s *RecordServiceImpl) FindByUser(userID uint, page int, limit int) (response.PaginatedResponse[response.StaffRecordResponse], error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	records, total, err := s.repo.FindByUserId(userID, offset, limit)
	if err != nil {
		return response.PaginatedResponse[response.StaffRecordResponse]{}, err
	}

	items := make([]response.StaffRecordResponse, 0, len(records))

	for _, r := range records {

		resp := response.StaffRecordResponse{
			ID:             r.ID,
			TrainingPlanID: r.TrainingPlanID,
			Status:         string(r.Status),
			CreatedAt:      r.CreatedAt,
			UpdatedAt:      r.UpdatedAt,
		}

		if r.TrainingPlan != nil {
			resp.TrainingPlanName = r.TrainingPlan.Name
			resp.TrainingDate = r.TrainingPlan.Date
			if(r.TrainingPlan.Location != nil) {
				resp.Location = r.TrainingPlan.Location
			}
			if(r.TrainingPlan.NumberOfHours != nil) {
				resp.NumberOfHours = *r.TrainingPlan.NumberOfHours
			}
			resp.SpeakerInstitute = r.TrainingPlan.SpeakerInstitute
			resp.TrainingType = string(r.TrainingPlan.Type)
		}

		items = append(items, resp)
	}

	return response.PaginatedResponse[response.StaffRecordResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		},
	}, nil
}



func (s *RecordServiceImpl) RegisterStaff(
	trainingPlanId uint,
	req request.RegisterStaffRequest,
) error {

	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	for _, userId := range req.UserIDs {
		if s.repo.Exists(userId, trainingPlanId) {
			continue // prevent duplicate registration
		}

		record := &model.Record{
			UserID:         userId,
			TrainingPlanID: trainingPlanId,
			Status:         model.RecordStatusRegister,
		}

		if err := s.repo.Save(record); err != nil {
			return err
		}
	}

	return nil
}

func (s *RecordServiceImpl) FindById(
	id int,
) (response.RecordResponse, error) {

	record, err := s.repo.FindById(id)
	if err != nil {
		return response.RecordResponse{}, err
	}

	return response.RecordResponse{
		ID:               record.ID,
		UserID:           record.UserID,
		UserName:         record.User.Name,
		TrainingPlanID:   record.TrainingPlanID,
		TrainingPlanName: record.TrainingPlan.Name,
		Status:           string(record.Status),
		CreatedAt:        record.CreatedAt,
		UpdatedAt:        record.UpdatedAt,
	}, nil
}

func (s *RecordServiceImpl) Update(
	id int,
	req request.UpdateRecordRequest,
) error {

	if err := s.validate.Struct(req); err != nil {
		return helper.ValidationError(helper.FormatValidationError(err))
	}

	record, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	record.Status = req.Status
	return s.repo.Update(record)
}

func (s *RecordServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *RecordServiceImpl) FindByManager(
	managerID uint,
	page int,
	limit int,
) (response.PaginatedResponse[response.AdminRecordResponse], error) {

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	manager, err := s.userRepo.FindById(managerID)
	if err != nil {
		return response.PaginatedResponse[response.AdminRecordResponse]{}, err
	}

	if manager.Role != model.RoleDepartmentManager {
		return response.PaginatedResponse[response.AdminRecordResponse]{},
			helper.Forbidden("only managers can access this resource")
	}

	offset := (page - 1) * limit

	records, total, err := s.repo.FindByManagerDepartment(
		manager.DepartmentID,
		offset,
		limit,
	)
	if err != nil {
		return response.PaginatedResponse[response.AdminRecordResponse]{}, err
	}

	items := make([]response.AdminRecordResponse, 0, len(records))

	for _, r := range records {

		resp := response.AdminRecordResponse{
			ID:             r.ID,
			Status:         string(r.Status),
			CreatedAt:      r.CreatedAt,
			UpdatedAt:      r.UpdatedAt,
		}

		if r.User != nil {
			resp.EmployeeName = r.User.Name
			resp.EmployeeID = r.User.EmployeeID
			resp.Position = r.User.Position
			if r.User.Department != nil {
				resp.Department = r.User.Department.Name
				resp.Division = string(r.User.Department.Division)
			}
		}

		if r.TrainingPlan != nil {
			resp.TrainingPlanName = r.TrainingPlan.Name
			resp.Location = r.TrainingPlan.Location
			resp.CostPerPerson = r.TrainingPlan.CostPerPerson
			resp.BudgetCode = r.TrainingPlan.BudgetCode
		}

		items = append(items, resp)
	}

	return response.PaginatedResponse[response.AdminRecordResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		},
	}, nil
}

func (s *RecordServiceImpl) Export(
	req request.RecordFilterRequest,
) (*excelize.File, error) {

	// Ignore pagination for export
	req.Page = 1
	req.Limit = 1000000

	records, _, err := s.repo.Search(req)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheet := "Records"
	f.SetSheetName("Sheet1", sheet)

	headers := []string{
		"Training Plan",
		"Location",
		"Cost Per Person",
		"Budget Code",
		"Employee ID",
		"Employee Name",
		"Position",
		"Department",
		"Division",
		"Status",
		"Updated At",
	}

	// ===== Header =====
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E9EFF7"},
			Pattern: 1,
		},
	})

	f.SetCellStyle(sheet, "A1", "K1", headerStyle)

	// ===== Data =====
	for i, r := range records {

		row := i + 2

		var trainingName, location, budgetCode string
		var costPerPerson interface{}
		var employeeID, employeeName, position string
		var department, division string

		if r.TrainingPlan != nil {
			trainingName = r.TrainingPlan.Name

			if r.TrainingPlan.Location != nil {
				location = *r.TrainingPlan.Location
			}

			if r.TrainingPlan.CostPerPerson != nil {
				costPerPerson = *r.TrainingPlan.CostPerPerson
			}

			if r.TrainingPlan.BudgetCode != nil {
				budgetCode = *r.TrainingPlan.BudgetCode
			}
		}

		if r.User != nil {
			employeeID = r.User.EmployeeID
			employeeName = r.User.Name
			position = r.User.Position

			if r.User.Department != nil {
				department = r.User.Department.Name
				division = string(r.User.Department.Division)
			}
		}

		values := []interface{}{
			trainingName,
			location,
			costPerPerson,
			budgetCode,
			employeeID,
			employeeName,
			position,
			department,
			division,
			string(r.Status),
			r.UpdatedAt.Format("2006-01-02 15:04"),
		}

		for col, val := range values {
			cell, _ := excelize.CoordinatesToCellName(col+1, row)
			f.SetCellValue(sheet, cell, val)
		}
	}

	// ===== Column Width =====
	for i := 1; i <= len(headers); i++ {
		col, _ := excelize.ColumnNumberToName(i)
		f.SetColWidth(sheet, col, col, 22)
	}

	// ===== Freeze Header Row =====
	f.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	// ===== Currency Format (Column C) =====
	currencyStyle, _ := f.NewStyle(&excelize.Style{
		NumFmt: 3, // 1,234 format
	})
	f.SetColStyle(sheet, "C", currencyStyle)

	return f, nil
}
