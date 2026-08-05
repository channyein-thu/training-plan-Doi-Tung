package seed

import (
	"log"
	"time"

	"training-plan-api/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedAdmin creates the initial HR department and admin user for local
// development. The admin password must be supplied by the caller (from
// ADMIN_SEED_PASSWORD) — there is no hardcoded default, so a deployment can
// never end up with a well-known "admin123" account. Callers should only invoke
// this in development (see main.go).
func SeedAdmin(db *gorm.DB, adminPassword string) {

	if adminPassword == "" {
		log.Println(" Skipping admin seed: ADMIN_SEED_PASSWORD is not set")
		return
	}

	//  Ensure HR Department exists
	var department model.Department
	err := db.
		Where("name = ? AND division = ?", "HR", model.Administration).
		First(&department).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			department = model.Department{
				Name:     "HR",
				Division: model.Administration,
			}
			if err := db.Create(&department).Error; err != nil {
				log.Fatal(" Failed to create HR department:", err)
			}
			log.Println("HR department created")
		} else {
			log.Fatal(" Department query failed:", err)
		}
	}

	//  Check if Admin already exists
	var count int64
	db.Model(&model.User{}).
		Where("role = ?", model.RoleHRAdmin).
		Count(&count)

	if count > 0 {
		log.Println(" Admin user already exists")
		return
	}

	//  Hash password
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(adminPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Fatal(" Failed to hash password:", err)
	}

	now := time.Now()

	//  Create Admin
	admin := model.User{
		DepartmentID: department.ID,
		Name:         "System Admin",
		Email:        "admin@company.com",
		EmployeeID:   "ADMIN001",
		Phone:        "0999999999",
		Password:     string(hashed),
		Role:         model.RoleHRAdmin,
		Status:       model.UserStatusActive,
		Position:     "HR Administrator",
		CreatedBy:    model.CreatedByAdmin,
		CreatedByID:  nil,
		WorkStartDate: &now,
		CreatedAt:     now.Unix(),
		UpdatedAt:     now.Unix(),
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal(" Failed to seed admin:", err)
	}

	log.Println(" Admin user seeded successfully")
}
