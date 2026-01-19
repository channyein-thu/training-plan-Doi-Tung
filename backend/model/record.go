package model

import "time"

type RecordStatus string

const (
	RecordStatusRegister RecordStatus = "Register"
	RecordStatusAttended RecordStatus = "Attended"
	RecordStatusAbsent   RecordStatus = "Absent"
)

type Record struct {
	ID        uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint         `gorm:"not null" json:"userId"`
	User      *User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CourseID  uint         `gorm:"not null" json:"courseId"`
	Course    *Course      `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Status    RecordStatus `gorm:"type:enum('Register','Attended','Absent');not null;default:'Register'" json:"status"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updatedAt"`
}
