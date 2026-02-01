package model

import "time"

type Role string

const (
	RoleHRAdmin           Role = "Hr(admin)"
	RoleDepartmentManager Role = "DepartmentHead(manager)"
	RoleStaff             Role = "Staff"
)

type UserStatus string

const (
	UserStatusActive   UserStatus = "Active"
	UserStatusInactive UserStatus = "Inactive"
)

type CreatedByType string

const (
	CreatedBySelf    CreatedByType = "self"
	CreatedByAdmin   CreatedByType = "admin"
	CreatedByManager CreatedByType = "manager"
)

type User struct {
	ID            uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	DepartmentID  int           `gorm:"not null" json:"departmentId"`
	Department    *Department   `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Name          string        `gorm:"type:varchar(52);not null" json:"name"`
	Password      string        `gorm:"type:text;not null" json:"-"`
	Email         string        `gorm:"type:varchar(52);unique;not null" json:"email"`
	EmployeeID    string        `gorm:"type:varchar(52);unique;not null" json:"employeeID"`
	Phone         string        `gorm:"type:varchar(20)" json:"phone"`
	Role          Role          `gorm:"type:enum('Hr(admin)','DepartmentHead(manager)','Staff');not null;default:'Staff'" json:"role"`
	Status        UserStatus    `gorm:"type:enum('Active','Inactive');not null;default:'Active'" json:"status"`
	Position      string        `gorm:"type:varchar(100)" json:"position"`
	WorkStartDate *time.Time    `gorm:"type:timestamp" json:"workStartDate,omitempty"`
	CreatedBy     CreatedByType `gorm:"type:enum('self','admin','manager');not null;default:'self'" json:"createdBy"`
	CreatedByID   *uint         `gorm:"index" json:"createdById,omitempty"`
	Error         int16         `gorm:"type:smallint;default:0" json:"error"`
	CreatedAt     int64         `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     int64         `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (r Role) IsValid() bool {
	return r == RoleHRAdmin || r == RoleDepartmentManager || r == RoleStaff
}
