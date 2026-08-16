package controller

import (
	"time"
	"training-plan-api/data/response"
	"training-plan-api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// DashboardController serves the aggregate counts shown on the dashboard stat
// cards. It queries the DB directly (like AuthController) since these are simple
// COUNT aggregates with no business logic worth a full service layer.
type DashboardController struct {
	db *gorm.DB
}

func NewDashboardController(db *gorm.DB) *DashboardController {
	return &DashboardController{db: db}
}

// AdminStats: organization-wide totals for the admin dashboard.
func (dc *DashboardController) AdminStats(c *fiber.Ctx) error {
	var totalUsers, totalTrainingPlans, totalDepartments int64

	dc.db.Model(&model.User{}).Count(&totalUsers)
	dc.db.Model(&model.TrainingPlan{}).Count(&totalTrainingPlans)
	dc.db.Model(&model.Department{}).Count(&totalDepartments)

	return c.JSON(response.Response{
		Status: "SUCCESS",
		Data: fiber.Map{
			"totalUsers":         totalUsers,
			"totalTrainingPlans": totalTrainingPlans,
			"totalDepartments":   totalDepartments,
		},
	})
}

// ManagerStats: department staff count, upcoming trainings, and the manager's
// own pending certificates.
func (dc *DashboardController) ManagerStats(c *fiber.Ctx) error {
	managerID := c.Locals("user_id").(uint)

	var departmentID int
	if err := dc.db.Table("users").
		Select("department_id").
		Where("id = ?", managerID).
		Scan(&departmentID).Error; err != nil {
		return err
	}

	var departmentStaff, activeTrainings, pendingCertificates int64

	dc.db.Model(&model.User{}).
		Where("department_id = ?", departmentID).
		Count(&departmentStaff)

	// "Active" = today or later (upcoming/ongoing). Compare against the date
	// portion so trainings scheduled for today are included.
	today := time.Now().Format("2006-01-02")
	dc.db.Model(&model.TrainingPlan{}).
		Where("date >= ?", today).
		Count(&activeTrainings)

	dc.db.Model(&model.Certificate{}).
		Where("user_id = ? AND status = ?", managerID, model.CertPending).
		Count(&pendingCertificates)

	return c.JSON(response.Response{
		Status: "SUCCESS",
		Data: fiber.Map{
			"departmentStaff":     departmentStaff,
			"activeTrainings":     activeTrainings,
			"pendingCertificates": pendingCertificates,
		},
	})
}

// StaffStats: the current staff member's own training and certificate counts.
func (dc *DashboardController) StaffStats(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var myTrainings, registeredTrainings, attendedTrainings, absentTrainings int64
	var pendingCertificates, approvedCertificates int64

	dc.db.Model(&model.Record{}).
		Where("user_id = ?", userID).
		Count(&myTrainings)

	dc.db.Model(&model.Record{}).
		Where("user_id = ? AND status = ?", userID, model.RecordStatusRegister).
		Count(&registeredTrainings)

	dc.db.Model(&model.Record{}).
		Where("user_id = ? AND status = ?", userID, model.RecordStatusAttended).
		Count(&attendedTrainings)

	dc.db.Model(&model.Record{}).
		Where("user_id = ? AND status = ?", userID, model.RecordStatusAbsent).
		Count(&absentTrainings)

	dc.db.Model(&model.Certificate{}).
		Where("user_id = ? AND status = ?", userID, model.CertPending).
		Count(&pendingCertificates)

	dc.db.Model(&model.Certificate{}).
		Where("user_id = ? AND status = ?", userID, model.CertApproved).
		Count(&approvedCertificates)

	return c.JSON(response.Response{
		Status: "SUCCESS",
		Data: fiber.Map{
			"myTrainings":          myTrainings,
			"registeredTrainings":  registeredTrainings,
			"attendedTrainings":    attendedTrainings,
			"absentTrainings":      absentTrainings,
			"pendingCertificates":  pendingCertificates,
			"approvedCertificates": approvedCertificates,
		},
	})
}
