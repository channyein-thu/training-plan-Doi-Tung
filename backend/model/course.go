package model

import "time"

type CourseCategory string
type CourseType string

const (
	// Category
	CategoryEnvironment      CourseCategory = "สนับสนุนนโยบายสิ่งแวดล้อม"
	CategorySafety           CourseCategory = "ความปลอดภัยและอาชีวอนามัย"
	CategorySalesService     CourseCategory = "งานขายและงานบริการ"
	CategorySoftware         CourseCategory = "การใช้งาน Software"
	CategoryPresentation     CourseCategory = "การนำเสนอ"
	CategoryLeadership       CourseCategory = "Leadership Development"
	CategoryMachine          CourseCategory = "การใช้งานเครื่องจักรและซ่อมบำรุง"
	CategoryThinking         CourseCategory = "กระบวนการคิด วิเคราะห์"
	CategoryWorkProcess      CourseCategory = "พัฒนาทักษะกระบวนการทำงาน"
	CategoryProcurement      CourseCategory = "การจัดซื้อจัดจ้าง"
	CategoryCommunication    CourseCategory = "การสื่อสาร"
	CategorySeminar          CourseCategory = "โครงการสัมมนาอื่นๆ"
	CategoryManagement       CourseCategory = "พัฒนาขีดความสามารถระดับบริหาร"
	CategoryFinance          CourseCategory = "การเงินและการบัญชี"
)

const (
	// Type
	TypeInHouse      CourseType = "In-house"
	TypePublic       CourseType = "Public"
	TypeOJT          CourseType = "OJT"
	TypeSelfLearning CourseType = "Self-learning"
	TypeOnline       CourseType = "Online/Virtual"
)

type Course struct {
	ID                int            `gorm:"primaryKey;autoIncrement"`
	Name              string         `gorm:"type:varchar(52);not null"`
	SpeakerInstitute  *string         `gorm:"type:text"`
	Type              CourseType     `gorm:"type:enum('In-house','Public','OJT','Self-learning','Online/Virtual')"`
	Category          CourseCategory `gorm:"type:enum(
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

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
