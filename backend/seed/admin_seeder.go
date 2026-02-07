package seed

import (
	"log"
	"time"

	"training-plan-api/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&model.User{}).
		Where("role = ?", model.RoleHRAdmin).
		Count(&count)

	if count > 0 {
		log.Println("✅ Admin user already exists")
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte("admin123"),
		bcrypt.DefaultCost,
	)

	user := model.User{
		DepartmentID: 1,
		Name:         "System Admin",
		Email:        "admin@company.com",
		EmployeeID:   "ADMIN001",
		Phone:        "0999999999",
		Password:     string(hashed),
		Role:         model.RoleHRAdmin,
		Status:       model.UserStatusActive,
		Position:     "HR Administrator",
		CreatedBy:    model.CreatedByAdmin,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatal("❌ Failed to seed admin:", err)
	}

	log.Println("✅ Admin user seeded successfully")
}
