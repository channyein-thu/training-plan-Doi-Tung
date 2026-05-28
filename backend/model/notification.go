package model

import "time"

type NotificationType string

const (
	NotifTrainingRegistered  NotificationType = "training_registered"
	NotifCertificateApproved NotificationType = "certificate_approved"
	NotifCertificateRejected NotificationType = "certificate_rejected"
)

type Notification struct {
	ID        uint             `gorm:"primaryKey;autoIncrement"`
	UserID    uint             `gorm:"not null;index"`
	User      *User            `gorm:"foreignKey:UserID"`
	Type      NotificationType `gorm:"type:enum('training_registered','certificate_approved','certificate_rejected');not null"`
	Title     string           `gorm:"type:varchar(255);not null"`
	Message   string           `gorm:"type:text;not null"`
	IsRead    bool             `gorm:"default:false"`
	CreatedAt time.Time        `gorm:"autoCreateTime"`
	UpdatedAt time.Time        `gorm:"autoUpdateTime"`
}
