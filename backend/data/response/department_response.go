package response

import "training-plan-api/model"

type DepartmentResponse struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Division   model.Division `json:"division"`
	TotalStaff int64          `json:"totalStaff"`
}