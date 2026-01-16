package request

import "time"

type CreateCourseRequest struct {
	Name             string  `json:"name" validate:"required,min=3"`
	SpeakerInstitute *string `json:"speakerInstitute"`
	Type             string  `json:"type" validate:"required"`
	Category         string  `json:"category" validate:"required"`

	Date        time.Time `json:"date" validate:"required"`
	Content     *string   `json:"content"`

	NumberOfDays  int  `json:"numberOfDays" validate:"gte=1"`
	NumberOfHours *int `json:"numberOfHours" validate:"omitempty,gte=1"`

	Location        *string
	TotalCost       *int `validate:"omitempty,gte=0"`
	BudgetCode      *string
	NumberOfPerson  *int `validate:"omitempty,gte=0"`
	CostPerPerson   *int `validate:"omitempty,gte=0"`
}



type UpdateCourseRequest struct {
	Name             *string
	SpeakerInstitute *string
	Type             *string
	Category         *string

	Date        *time.Time
	Content     *string

	NumberOfDays  *int
	NumberOfHours *int

	Location        *string
	TotalCost       *int
	BudgetCode      *string
	NumberOfPerson  *int
	CostPerPerson   *int
}


