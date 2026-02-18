package model

import "time"

type TrainingPlanCategory string
type TrainingPlanType string

const (
	// Category
	CategoryEnvironment      TrainingPlanCategory = "สนับสนุนนโยบายสิ่งแวดล้อม"
	CategorySafety           TrainingPlanCategory = "ความปลอดภัยและอาชีวอนามัย"
	CategorySalesService     TrainingPlanCategory = "งานขายและงานบริการ"
	CategorySoftware         TrainingPlanCategory = "การใช้งาน Software"
	CategoryPresentation     TrainingPlanCategory = "การนำเสนอ"
	CategoryLeadership       TrainingPlanCategory = "Leadership Development"
	CategoryMachine          TrainingPlanCategory = "การใช้งานเครื่องจักรและซ่อมบำรุง"
	CategoryThinking         TrainingPlanCategory = "กระบวนการคิด วิเคราะห์"
	CategoryWorkProcess      TrainingPlanCategory = "พัฒนาทักษะกระบวนการทำงาน"
	CategoryProcurement      TrainingPlanCategory = "การจัดซื้อจัดจ้าง"
	CategoryCommunication    TrainingPlanCategory = "การสื่อสาร"
	CategorySeminar          TrainingPlanCategory = "โครงการสัมมนาอื่นๆ"
	CategoryManagement       TrainingPlanCategory = "พัฒนาขีดความสามารถระดับบริหาร"
	CategoryFinance          TrainingPlanCategory = "การเงินและการบัญชี"
)

const (
	// Type
	TypeInHouse      TrainingPlanType = "In-house"
	TypePublic       TrainingPlanType = "Public"
	TypeOJT          TrainingPlanType = "OJT"
	TypeSelfLearning TrainingPlanType = "Self-learning"
	TypeOnline       TrainingPlanType = "Online/Virtual"
)

type TrainingPlan struct {
	ID                int            `gorm:"primaryKey;autoIncrement"`
	Name              string         `gorm:"type:varchar(52);not null"`
	SpeakerInstitute  *string         `gorm:"type:text"`
	Type              TrainingPlanType     `gorm:"type:enum('In-house','Public','OJT','Self-learning','Online/Virtual')"`
	Category          TrainingPlanCategory `gorm:"type:enum(
		'สนับสนุนนโยบายสิ่งแวดล้อม',
		'ความปลอดภัยและอาชีวอนามัย',
		'งานขายและงานบริการ',
		'การใช้งาน Software',
		'การนำเสนอ',
		'Leadership Development',
		'การใช้งานเครื่องจักรและซ่อมบำรุง',
		'กระบวนการคิด วิเคราะห์',
		'พัฒนาทักษะกระบวนการทำงาน',
		'การจัดซื้อจัดจ้าง',
		'การสื่อสาร',
		'โครงการสัมมนาอื่นๆ',
		'พัฒนาขีดความสามารถระดับบริหาร',
		'การเงินและการบัญชี'
	)"`
	Date              time.Time `gorm:"type:date;not null"`
	Content       string 	  `gorm:"type:text"`
	NumberOfDays      int `gorm:"default:1"`
	NumberOfHours     *int 
	Location          *string         `gorm:"type:text"`
	TotalCost         *int
	BudgetCode        *string         `gorm:"type:varchar(52)"`
	NumberOfPerson    int `gorm:"default:0"`
	CostPerPerson     *int `gorm:"type:int"`

	CalendarEventID *string `gorm:"type:varchar(128);index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (TrainingPlan) TableName() string {
	return "training_plans"
}

