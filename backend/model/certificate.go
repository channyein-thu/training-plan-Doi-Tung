// package model

// import "time"

// type CertificateCategory string

// type CertificateType string

// type Certificate struct {
// 	ID           uint                `gorm:"primaryKey;autoIncrement" json:"id"`
// 	UserID       uint                `gorm:"not null" json:"userId"`
// 	User         *User               `gorm:"foreignKey:UserID" json:"user,omitempty"`
// 	TrainingName string              `gorm:"type:varchar(255);not null" json:"trainingName"`
// 	Category     CertificateCategory `gorm:"type:enum" json:"category,omitempty"`
// 	Type         CertificateType     `gorm:"type:enum" json:"type,omitempty"`
// 	Image        string              `gorm:"type:text" json:"image,omitempty"`
// 	Description  string              `gorm:"type:text" json:"description,omitempty"`
// 	CreatedAt    time.Time           `gorm:"autoCreateTime" json:"createdAt"`
// 	UpdatedAt    time.Time           `gorm:"autoUpdateTime" json:"updatedAt"`
// }

package model

import "time"

type CertificateStatus string

const (
	CertPending  CertificateStatus = "Pending"
	CertApproved CertificateStatus = "Approved"
	CertRejected CertificateStatus = "Rejected"
)

type Certificate struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	UserID uint `gorm:"not null;index"`
	User   *User `gorm:"foreignKey:UserID"`

	TrainingName string `gorm:"type:varchar(255);not null"`
	Image        string `gorm:"type:text;not null"`
	Description  *string `gorm:"type:text"`

	Status CertificateStatus `gorm:"type:enum('Pending','Approved','Rejected');default:'Pending'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}