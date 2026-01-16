package model

import "time"

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
	Date              time.Time
	Content       *string
	NumberOfDays      int `gorm:"default:1"`
	NumberOfHours     *int
	Location          *string         `gorm:"type:text"`
	TotalCost         *int
	BudgetCode        *string         `gorm:"type:varchar(52)"`
	NumberOfPerson    int `gorm:"default:0"`
	CostPerPerson     *int

	CalendarEventID *string `gorm:"type:varchar(128);index"`

	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
}
