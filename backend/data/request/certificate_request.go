package request

type CreateCertificateRequest struct {
	TrainingID    uint    `json:"trainingId" validate:"required"`
	Image        string  `json:"image" validate:"omitempty"`
	Description  *string `json:"description" validate:"omitempty,max=1000"`
}

type UpdateCertificateRequest struct {
	Image        *string `json:"image" validate:"omitempty"`
	Description  *string `json:"description" validate:"omitempty,max=1000"`
}

type UpdateCertificateStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=Approved Rejected"`
}