package model

import "time"

// PushSubscription stores a single browser's Web Push subscription for a user.
// A user can have multiple subscriptions (laptop + phone, multiple browsers).
type PushSubscription struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null;index"`
	User      *User     `gorm:"foreignKey:UserID"`
	Endpoint  string    `gorm:"type:text;not null"`
	P256dh    string    `gorm:"type:varchar(255);not null"`
	Auth      string    `gorm:"type:varchar(255);not null"`
	UserAgent string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
