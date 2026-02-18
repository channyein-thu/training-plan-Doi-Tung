package response

import "time"

type TrainingPlanResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	SpeakerInstitute *string `json:"speakerInstitute,omitempty"`
	Type             string  `json:"type"`
	Category         string  `json:"category"`

	Date    time.Time `json:"date"`
	Content string   `json:"content"`

	NumberOfDays  int     `json:"numberOfDays"`
	NumberOfHours *int    `json:"numberOfHours,omitempty"`

	Location       *string `json:"location,omitempty"`
	TotalCost      *int    `json:"totalCost,omitempty"`
	BudgetCode     *string `json:"budgetCode,omitempty"`
	NumberOfPerson int     `json:"numberOfPerson"`
	CostPerPerson  *int    `json:"costPerPerson,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
