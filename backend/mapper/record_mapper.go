package mapper

import (
	"training-plan-api/data/response"
	"training-plan-api/model"
)

func ToRecordResponse(record model.Record) response.RecordResponseFinal  {
	resp := response.RecordResponseFinal {
		ID:             record.ID,
		UserID:         record.UserID,
		TrainingPlanID: record.TrainingPlanID,
		TrainingPlan:   ToTrainingPlanResponse(*record.TrainingPlan),
		Status:         string(record.Status),
		Evaluation:     record.Evaluation,
		PreTestScore:  record.PreTestScore,
		PostTestScore: record.PostTestScore,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
	}
	if record.User != nil {
		resp.User = &response.UserResponse{
			ID:           record.User.ID,
			EmployeeID:   record.User.EmployeeID,
			Name:         record.User.Name,
			Email:        record.User.Email,
			Phone:        record.User.Phone,
			DepartmentID: record.User.DepartmentID,
			Role:         record.User.Role,
			Status:       record.User.Status,
			Position:     record.User.Position,		
		}
		if record.User.Department != nil {
			resp.User.Department = &response.DepartmentResponse{
				ID:   record.User.Department.ID,
				Name: record.User.Department.Name,
				Division: record.User.Department.Division,
			}
		}
	}
		
	return resp
	
}