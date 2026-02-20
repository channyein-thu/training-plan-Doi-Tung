package request

import (
	"time"
	"training-plan-api/model"
)

type RegisterStaffRequest struct {
	UserIDs []uint `json:"userIds" validate:"required,min=1,dive,gt=0"`
}

type UpdateRecordRequest struct {
	Status model.RecordStatus `json:"status" validate:"required,oneof=Register Attended Absent"`
}

type RecordFilterRequest struct {
	DepartmentIDs []int    `json:"departmentIds"`
	Categories    []string `json:"categories"`
	Status        *string  `json:"status"`

	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`

	Page  int `json:"page"`
	Limit int `json:"limit"`
}
