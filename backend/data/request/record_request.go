package request

import (
	"training-plan-api/model"
)

type RegisterStaffRequest struct {
	UserIDs []uint `json:"userIds" validate:"required,min=1,dive,gt=0"`
}

type UpdateRecordRequest struct {
	Status model.RecordStatus `json:"status" validate:"required,oneof=Register Attended Absent"`
	Evaluation     *string `json:"evaluation" validate:"omitempty"`
	PreTestScore  *int `json:"preTestScore" validate:"omitempty,gte=0"`
	PostTestScore *int `json:"postTestScore" validate:"omitempty,gte=0"`
}

type RecordFilterRequest struct {
	DepartmentIDs []int    `json:"departmentIds"`
	Categories    []string `json:"categories"`
	Status        *string  `json:"status"`

	// Date-only strings (e.g. "2026-08-05") as sent by the frontend date picker.
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`

	Page  int `json:"page"`
	Limit int `json:"limit"`
}
