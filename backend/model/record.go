package model

import "time"

type RecordStatus string

const (
	RecordStatusRegister RecordStatus = "Register"
	RecordStatusAttended RecordStatus = "Attended"
	RecordStatusAbsent   RecordStatus = "Absent"
)

type Record struct {
	ID             uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uint         `gorm:"not null" json:"userId"`
	User           *User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TrainingPlanID uint         `gorm:"not null" json:"trainingPlanId"`
	TrainingPlan   *TrainingPlan `gorm:"foreignKey:TrainingPlanID" json:"trainingPlan,omitempty"`
	Status         RecordStatus `gorm:"type:enum('Register','Attended','Absent');not null;default:'Register'" json:"status"`
	Evaluation     *string      `gorm:"type:text" json:"evaluation,omitempty"`
	PreTestScore  *int         `gorm:"type:int" json:"preTestScore,omitempty"`
	PostTestScore *int         `gorm:"type:int" json:"postTestScore,omitempty"`
	CreatedAt      time.Time    `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime" json:"updatedAt"`
}
