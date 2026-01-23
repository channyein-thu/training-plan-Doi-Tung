package request

import "time"

type CreateCourseRequest struct {
	Name             string  `json:"name" validate:"required,min=3"`
	SpeakerInstitute *string `json:"speakerInstitute"`
	Type             string  `json:"type" validate:"required"`
	Category         string  `json:"category" validate:"required"`

	Date        time.Time `json:"date" validate:"required"`
	Content     string   `json:"content" validate:"required,min=10"`

	NumberOfDays  int  `json:"numberOfDays" validate:"gte=1"`
	NumberOfHours *int `json:"numberOfHours" validate:"omitempty,gte=1"`

	Location        *string `json:"location" validate:"omitempty"`
	TotalCost       *int `json:"totalCost" validate:"omitempty,gte=0"`
	BudgetCode      *string `json:"budgetCode" validate:"omitempty"`
	NumberOfPerson  *int `json:"numberOfPerson" validate:"omitempty,gte=0"`
	CostPerPerson   *int `json:"costPerPerson" validate:"omitempty,gte=0"`
}



type UpdateCourseRequest struct {
	Name             *string    `json:"name" validate:"omitempty,min=3"`
	SpeakerInstitute *string    `json:"speakerInstitute" validate:"omitempty"`

	Type     *string `json:"type" validate:"omitempty"`
	Category *string `json:"category" validate:"omitempty"`

	Date    *time.Time `json:"date" validate:"omitempty"`
	Content *string   `json:"content" validate:"omitempty"`

	NumberOfDays  *int `json:"numberOfDays" validate:"omitempty,gte=1"`
	NumberOfHours *int `json:"numberOfHours" validate:"omitempty,gte=1"`

	Location       *string `json:"location" validate:"omitempty"`
	TotalCost      *int    `json:"totalCost" validate:"omitempty,gte=0"`
	BudgetCode     *string `json:"budgetCode" validate:"omitempty"`
	NumberOfPerson *int    `json:"numberOfPerson" validate:"omitempty,gte=0"`
	CostPerPerson  *int    `json:"costPerPerson" validate:"omitempty,gte=0"`
}


