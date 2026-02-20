package model

type Division string

const (
	SocialEnterprise Division = "Social Enterprise"
	DevelopProject Division = "Development Project"
	NatureBasedSolutionAndSpecialProject Division = "Nature-based Solution and Special Project"
	Sustainability Division = "Sustainability"
	AccountingAndFinance Division = "Accounting and Finance"
	Administration Division = "Administration"
	Other Division = "Other (under CEO)"
)

type Department struct {
	ID       int      `gorm:"primaryKey;autoIncrement"`
	Name     string   `gorm:"type:varchar(100);not null;uniqueIndex:idx_dept_division"`
	Division Division `gorm:"type:enum('Social Enterprise','Development Project','Nature-based Solution and Special Project','Sustainability','Accounting and Finance','Administration','Other (under CEO)');not null;uniqueIndex:idx_dept_division"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
