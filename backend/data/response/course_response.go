package response

import "time"

type CourseResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	SpeakerInstitute *string `json:"speakerInstitute,omitempty"`
	Type             string  `json:"type"`
	Category         string  `json:"category"`

	Date    time.Time `json:"date"`
	Content *string   `json:"content,omitempty"`

	NumberOfDays   int
	NumberOfHours  *int
	Location        *string
	TotalCost       *int
	BudgetCode      *string
	NumberOfPerson  int
	CostPerPerson   *int

	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}
