package response

import "time"

type CertificateResponse struct {
	ID uint `json:"id"`

	UserID   uint   `json:"userId"`
	UserName string `json:"userName"`
	EmployeeID string `json:"employeeId"`
	Department string `json:"department"`
	Division string   `json:"division"`
	Category   string  `json:"category"`

	TrainingID   uint   `json:"trainingId"`
	TrainingName string  `json:"trainingName"`
	Image        string  `json:"image"`
	Description  *string `json:"description,omitempty"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserCertificateResponse struct {
	ID uint `json:"id"`

	TrainingName string `json:"trainingName"`
	Status       string `json:"status"`
	Image        string `json:"image"`

	Description *string `json:"description,omitempty"`
}
