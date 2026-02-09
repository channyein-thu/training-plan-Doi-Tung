package request

import "training-plan-api/model"

type RegisterStaffRequest struct {
	UserIDs []uint `json:"userIds" validate:"required,min=1,dive,gt=0"`
}

type UpdateRecordRequest struct {
	Status model.RecordStatus `json:"status" validate:"required,oneof=Register Attended Absent"`
}
