package response

import "time"

type RecordResponse struct {
	ID               uint      `json:"id"`
	UserID           uint      `json:"userId"`
	UserName         string    `json:"userName"`
	Position         string    `json:"position"`
	Department       string    `json:"department"`
	Division         string    `json:"division"`
	TrainingPlanID   uint      `json:"trainingPlanId"`
	TrainingPlanName string    `json:"trainingPlanName"`
	Location         string    `json:"location"`
	CostPerPerson    *int64    `json:"costPerPerson,omitempty"`
	BudgetCode       *string   `json:"budgetCode,omitempty"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type StaffRecordResponse struct {
	ID               uint      `json:"id"`
	TrainingPlanID   uint      `json:"trainingPlanId"`
	TrainingPlanName string    `json:"trainingPlanName"`
	Location         string    `json:"location"`
	TrainingDate     time.Time `json:"trainingDate"`
	NumberOfHours    int       `json:"numberOfHours"`
	SpeakerInstitute *string   `json:"speakerInstitute,omitempty"`
	TrainingType     string    `json:"trainingType"`
}