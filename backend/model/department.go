package model

type Division string

const (
	DivisionA Division = "Division A"
	DivisionB Division = "Division B"
	DivisionC Division = "Division C"
)


type Department struct {
	ID       int      `gorm:"primaryKey;autoIncrement"`
	Name     string   `gorm:"type:varchar(100);not null;uniqueIndex:idx_dept_division"`
	Division Division `gorm:"type:enum('Division A','Division B','Division C');not null;uniqueIndex:idx_dept_division"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
