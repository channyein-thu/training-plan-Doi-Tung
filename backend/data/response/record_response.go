package response

import "time"

type RecordResponse struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"userId"`
	UserName   string    `json:"userName"`
	Position   string    `json:"position"`
	Department string    `json:"department"`
	Division string      `json:"division"`
	CourseID   uint      `json:"courseId"`
	CourseName string    `json:"courseName"`
	Location  string    `json:"location"`
	CostPerPerson *int64     `json:"costPerPerson,omitempty"`
	BudgetCode   *string   `json:"budgetCode,omitempty"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
