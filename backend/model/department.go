package model

type Department struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
	DivisionID *int   `gorm:"default:null" json:"division_id,omitempty"`
	CreatedAt  int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  int64  `gorm:"autoUpdateTime" json:"updated_at"`
}
