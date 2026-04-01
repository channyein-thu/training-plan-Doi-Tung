package response

import "time"

type AdminRecordResponse struct {
	ID               uint      `json:"id"`
	TrainingPlanName string    `json:"trainingPlanName"`
	Location 	   *string    `json:"location"`
	CostPerPerson    *int    	`json:"costPerPerson,omitempty"`
	BudgetCode       *string   `json:"budgetCode,omitempty"`
	EmployeeID       string     `json:"employeeId"`
	EmployeeName     string    `json:"employeeName"`
	Position         string    `json:"position"`
	Department       string    `json:"department"`
	Division         string    `json:"division"`
	Status           string    `json:"status"`
	Evaluation       *string   `json:"evaluation,omitempty"`
	PreTestScore     *int      `json:"preTestScore,omitempty"`
	PostTestScore    *int      `json:"postTestScore,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type RecordResponse struct {
	ID               uint      `json:"id"`
	UserID           uint      `json:"userId"`
	UserName         string    `json:"userName"`
	Position         string    `json:"position"`
	Department       string    `json:"department"`
	Division         string    `json:"division"`
	TrainingPlanID   uint      `json:"trainingPlanId"`
	TrainingPlanName string    `json:"trainingPlanName"`
	Location         *string    `json:"location"`
	CostPerPerson    *int    `json:"costPerPerson,omitempty"`
	BudgetCode       *string   `json:"budgetCode,omitempty"`
	Status           string    `json:"status"`
	Evaluation       *string   `json:"evaluation,omitempty"`
	PreTestScore     *int      `json:"preTestScore,omitempty"`
	PostTestScore    *int      `json:"postTestScore,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type StaffRecordResponse struct {
	ID               uint      `json:"id"`
	Division		 string    `json:"division"`
	Department 	 string    `json:"department"`
	BudgetCode       *string   `json:"budgetCode,omitempty"`
	CostPerPerson    *int    `json:"costPerPerson,omitempty"`
	Position		 string    `json:"position"`
	TrainingPlanID   uint      `json:"trainingPlanId"`
	TrainingPlanName string    `json:"trainingPlanName"`
	Status           string    `json:"status"`
	Location         *string    `json:"location"`
	TrainingDate     time.Time `json:"trainingDate"`
	NumberOfHours    int       `json:"numberOfHours"`
	SpeakerInstitute *string   `json:"speakerInstitute,omitempty"`
	TrainingType     string    `json:"trainingType"`
	Evaluation       *string   `json:"evaluation,omitempty"`
	PreTestScore     *int      `json:"preTestScore,omitempty"`
	PostTestScore    *int      `json:"postTestScore,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type FinalRecord struct{
	ID               uint      `json:"id"`
	UserID		   uint      `json:"userId"`
	TrainingPlanID  uint      `json:"trainingPlanId"`
	Status           string    `json:"status"`
	Evaluation       *string   `json:"evaluation,omitempty"`
	PreTestScore     *int      `json:"preTestScore,omitempty"`
	PostTestScore    *int      `json:"postTestScore,omitempty"`
	EmployeeID   string              `json:"employeeID"`
	StaffName         string              `json:"staffName"`
	DepartmentID int                 `json:"departmentId"`
	Departmentname string              `json:"departmentName"`
	Division   string              `json:"division"`
	Position     string              `json:"position"`
	TrainingPlanName string `json:"trainingPlanName"`

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
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`

} 

type RecordResponseFinal struct {
	ID               uint      `json:"id"`
	UserID		   uint      `json:"userId"`
	User        *UserResponse	   `json:"user"`
	TrainingPlanID  uint      `json:"trainingPlanId"`
	TrainingPlan    TrainingPlanResponse `json:"trainingPlan"`
	Status           string    `json:"status"`
	Evaluation       *string   `json:"evaluation,omitempty"`
	PreTestScore     *int      `json:"preTestScore,omitempty"`
	PostTestScore    *int      `json:"postTestScore,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}



