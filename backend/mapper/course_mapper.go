package mapper

import (
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/model"
)

func ToCourseModel(req request.CreateCourseRequest) model.Course {
	course := model.Course{
		Name:             req.Name,
		SpeakerInstitute: req.SpeakerInstitute,
		Type:             model.CourseType(req.Type),
		Category:         model.CourseCategory(req.Category),

		Date:        req.Date,
		Content:     req.Content,
		NumberOfDays: req.NumberOfDays,

		NumberOfHours: req.NumberOfHours,
		Location:       req.Location,
		TotalCost:      req.TotalCost,
		BudgetCode:     req.BudgetCode,
		CostPerPerson:  req.CostPerPerson,
	}

	if req.NumberOfPerson != nil {
		course.NumberOfPerson = *req.NumberOfPerson
	}

	return course
}

func ToCourseResponse(course model.Course) response.CourseResponse {
	return response.CourseResponse{
		ID:   course.ID,
		Name: course.Name,

		SpeakerInstitute: course.SpeakerInstitute,
		Type:             string(course.Type),
		Category:         string(course.Category),

		Date:    course.Date,
		Content: course.Content,

		NumberOfDays:  course.NumberOfDays,
		NumberOfHours: course.NumberOfHours,
		Location:       course.Location,
		TotalCost:      course.TotalCost,
		BudgetCode:     course.BudgetCode,
		NumberOfPerson: course.NumberOfPerson,
		CostPerPerson:  course.CostPerPerson,

		CreatedAt: course.CreatedAt,
		UpdatedAt: course.UpdatedAt,
	}
}


func UpdateCourseFromRequest(course *model.Course, req request.UpdateCourseRequest) {
	if req.Name != nil {
		course.Name = *req.Name
	}

	if req.SpeakerInstitute != nil {
		course.SpeakerInstitute = req.SpeakerInstitute
	}

	if req.Type != nil {
		course.Type = model.CourseType(*req.Type)
	}

	if req.Category != nil {
		course.Category = model.CourseCategory(*req.Category)
	}

if req.Date != nil {
	course.Date = *req.Date
}

if req.Content != nil {
	course.Content = *req.Content
}

if req.NumberOfDays != nil {
	course.NumberOfDays = *req.NumberOfDays
}


	if req.NumberOfHours != nil {
		course.NumberOfHours = req.NumberOfHours
	}

	if req.Location != nil {
		course.Location = req.Location
	}

	if req.TotalCost != nil {
		course.TotalCost = req.TotalCost
	}

	if req.BudgetCode != nil {
		course.BudgetCode = req.BudgetCode
	}

	if req.NumberOfPerson != nil {
		course.NumberOfPerson = *req.NumberOfPerson
	}

	if req.CostPerPerson != nil {
		course.CostPerPerson = req.CostPerPerson
	}
}

func ToCourseResponseList(courses []model.Course) []response.CourseResponse {
	responses := make([]response.CourseResponse, len(courses))

	for i, course := range courses {
		responses[i] = ToCourseResponse(course)
	}

	return responses
}

