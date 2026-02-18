package mapper

import (
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/model"
)

func ToTrainingPlanModel(req request.CreateTrainingPlanRequest) model.TrainingPlan {
	trainingPlan := model.TrainingPlan{
		Name:             req.Name,
		SpeakerInstitute: req.SpeakerInstitute,
		Type:             model.TrainingPlanType(req.Type),
		Category:         model.TrainingPlanCategory(req.Category),

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
		trainingPlan.NumberOfPerson = *req.NumberOfPerson
	}

	return trainingPlan
}

func ToTrainingPlanResponse(trainingPlan model.TrainingPlan) response.TrainingPlanResponse {
	return response.TrainingPlanResponse{
		ID:   trainingPlan.ID,
		Name: trainingPlan.Name,

		SpeakerInstitute: trainingPlan.SpeakerInstitute,
		Type:             string(trainingPlan.Type),
		Category:         string(trainingPlan.Category),

		Date:    trainingPlan.Date,
		Content: trainingPlan.Content,

		NumberOfDays:  trainingPlan.NumberOfDays,
		NumberOfHours: trainingPlan.NumberOfHours,
		Location:       trainingPlan.Location,
		TotalCost:      trainingPlan.TotalCost,
		BudgetCode:     trainingPlan.BudgetCode,
		NumberOfPerson: trainingPlan.NumberOfPerson,
		CostPerPerson:  trainingPlan.CostPerPerson,

		CreatedAt: trainingPlan.CreatedAt,
		UpdatedAt: trainingPlan.UpdatedAt,
	}
}


func UpdateTrainingPlanFromRequest(trainingPlan *model.TrainingPlan, req request.UpdateTrainingPlanRequest) {
	if req.Name != nil {
		trainingPlan.Name = *req.Name
	}

	if req.SpeakerInstitute != nil {
		trainingPlan.SpeakerInstitute = req.SpeakerInstitute
	}

	if req.Type != nil {
		trainingPlan.Type = model.TrainingPlanType(*req.Type)
	}

	if req.Category != nil {
		trainingPlan.Category = model.TrainingPlanCategory(*req.Category)
	}

if req.Date != nil {
	trainingPlan.Date = *req.Date
}

if req.Content != nil {
	trainingPlan.Content = *req.Content
}

if req.NumberOfDays != nil {
	trainingPlan.NumberOfDays = *req.NumberOfDays
}


	if req.NumberOfHours != nil {
		trainingPlan.NumberOfHours = req.NumberOfHours
	}

	if req.Location != nil {
		trainingPlan.Location = req.Location
	}

	if req.TotalCost != nil {
		trainingPlan.TotalCost = req.TotalCost
	}

	if req.BudgetCode != nil {
		trainingPlan.BudgetCode = req.BudgetCode
	}

	if req.NumberOfPerson != nil {
		trainingPlan.NumberOfPerson = *req.NumberOfPerson
	}

	if req.CostPerPerson != nil {
		trainingPlan.CostPerPerson = req.CostPerPerson
	}
}

func ToTrainingPlanResponseList(trainingPlans []model.TrainingPlan) []response.TrainingPlanResponse {
	responses := make([]response.TrainingPlanResponse, len(trainingPlans))

	for i, trainingPlan := range trainingPlans {
		responses[i] = ToTrainingPlanResponse(trainingPlan)
	}

	return responses
}

