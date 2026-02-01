package controller

import (
	"time"

	"training-plan-api/helper"
	"training-plan-api/model"
	"training-plan-api/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{db: db}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name            string `json:"name"`
	EmployeeID      string `json:"employeeID"`
	Email           string `json:"email"`
	DepartmentID    int    `json:"departmentId"`
	Position        string `json:"position"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (ac *AuthController) AdminLogin(c *fiber.Ctx) error {
	return ac.handleLogin(c, model.RoleHRAdmin)
}

func (ac *AuthController) ManagerLogin(c *fiber.Ctx) error {
	return ac.handleLogin(c, model.RoleDepartmentManager)
}

func (ac *AuthController) ManagerRegister(c *fiber.Ctx) error {
	return ac.handleRegister(c, model.RoleDepartmentManager, "Manager registration successful")
}

func (ac *AuthController) StaffLogin(c *fiber.Ctx) error {
	return ac.handleLogin(c, model.RoleStaff)
}

func (ac *AuthController) StaffRegister(c *fiber.Ctx) error {
	return ac.handleRegister(c, model.RoleStaff, "Staff registration successful")
}

func (ac *AuthController) GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user model.User
	if err := ac.db.Preload("Department").First(&user, userID).Error; err != nil {
		return helper.NotFound("User not found")
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user":    ac.buildUserResponse(&user),
	})
}

func (ac *AuthController) handleLogin(c *fiber.Ctx, role model.Role) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	var user model.User
	query := ac.db.Preload("Department").Where("email = ? AND role = ?", req.Email, role)
	if err := query.First(&user).Error; err != nil {
		return helper.Unauthorized("Invalid credentials")
	}

	if user.Status != model.UserStatusActive {
		return helper.Unauthorized("Account is deactivated")
	}

	if !utils.ComparePassword(user.Password, req.Password) {
		return helper.Unauthorized("Invalid credentials")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return helper.InternalServerError("Failed to generate access token")
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return helper.InternalServerError("Failed to generate refresh token")
	}

	ac.db.Model(&model.RefreshToken{}).
		Where("user_id = ? AND revoked = ?", user.ID, false).
		Update("revoked", true)

	refreshTokenRecord := model.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(utils.RefreshTokenExpiry),
		Revoked:   false,
	}
	if err := ac.db.Create(&refreshTokenRecord).Error; err != nil {
		return helper.InternalServerError("Failed to store refresh token")
	}

	utils.SetAccessTokenCookie(c, accessToken)
	utils.SetRefreshTokenCookie(c, refreshToken)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"user":    ac.buildUserResponse(&user),
	})
}

func (ac *AuthController) Refresh(c *fiber.Ctx) error {
	refreshToken := utils.GetRefreshTokenFromCookie(c)
	if refreshToken == "" {
		return helper.Unauthorized("Refresh token required")
	}

	var tokenRecord model.RefreshToken
	if err := ac.db.Where("token = ?", refreshToken).First(&tokenRecord).Error; err != nil {
		return helper.Unauthorized("Invalid refresh token")
	}

	if !tokenRecord.IsValid() {
		utils.ClearAuthCookies(c)
		return helper.Unauthorized("Refresh token is invalid or expired")
	}

	var user model.User
	if err := ac.db.First(&user, tokenRecord.UserID).Error; err != nil {
		return helper.Unauthorized("User not found")
	}

	if user.Status != model.UserStatusActive {
		ac.db.Model(&tokenRecord).Update("revoked", true)
		utils.ClearAuthCookies(c)
		return helper.Unauthorized("Account is deactivated")
	}

	ac.db.Model(&tokenRecord).Update("revoked", true)

	newAccessToken, err := utils.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return helper.InternalServerError("Failed to generate access token")
	}

	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return helper.InternalServerError("Failed to generate refresh token")
	}

	newTokenRecord := model.RefreshToken{
		UserID:    user.ID,
		Token:     newRefreshToken,
		ExpiresAt: time.Now().Add(utils.RefreshTokenExpiry),
		Revoked:   false,
	}
	if err := ac.db.Create(&newTokenRecord).Error; err != nil {
		return helper.InternalServerError("Failed to store refresh token")
	}

	utils.SetAccessTokenCookie(c, newAccessToken)
	utils.SetRefreshTokenCookie(c, newRefreshToken)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Tokens refreshed successfully",
	})
}

func (ac *AuthController) Logout(c *fiber.Ctx) error {
	refreshToken := utils.GetRefreshTokenFromCookie(c)

	if refreshToken != "" {
		ac.db.Model(&model.RefreshToken{}).
			Where("token = ?", refreshToken).
			Update("revoked", true)
	}

	utils.ClearAuthCookies(c)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Logged out successfully",
	})
}

func (ac *AuthController) handleRegister(c *fiber.Ctx, role model.Role, successMsg string) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	if err := ac.validateRegisterRequest(&req); err != nil {
		return err
	}

	if err := ac.checkDuplicateUser(req.Email, req.EmployeeID); err != nil {
		return err
	}

	var department model.Department
	if err := ac.db.First(&department, req.DepartmentID).Error; err != nil {
		return helper.BadRequest("Invalid department selected")
	}

	user := ac.createUserFromRequest(&req, role)
	if err := ac.db.Create(&user).Error; err != nil {
		return helper.BadRequest("Failed to create user")
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": successMsg,
	})
}

func (ac *AuthController) validateRegisterRequest(req *RegisterRequest) error {
	if req.Name == "" || req.EmployeeID == "" || req.Email == "" ||
		req.Position == "" || req.Password == "" || req.DepartmentID == 0 {
		return helper.BadRequest("All required fields must be filled")
	}
	if req.Password != req.ConfirmPassword {
		return helper.BadRequest("Passwords do not match")
	}
	return nil
}

func (ac *AuthController) checkDuplicateUser(email, employeeID string) error {
	var existingUser model.User
	if ac.db.Where("email = ?", email).First(&existingUser).Error == nil {
		return helper.BadRequest("Email already registered")
	}
	if ac.db.Where("employee_id = ?", employeeID).First(&existingUser).Error == nil {
		return helper.BadRequest("Employee ID already registered")
	}
	return nil
}

func (ac *AuthController) createUserFromRequest(req *RegisterRequest, role model.Role) model.User {
	user := model.User{
		Name:         req.Name,
		EmployeeID:   req.EmployeeID,
		Email:        req.Email,
		DepartmentID: req.DepartmentID,
		Position:     req.Position,
		Role:         role,
		Password:     utils.GeneratePassword(req.Password),
		Status:       model.UserStatusActive,
	}
	return user
}

func (ac *AuthController) buildUserResponse(user *model.User) fiber.Map {
	response := fiber.Map{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"employeeID": user.EmployeeID,
		"role":       user.Role,
		"status":     user.Status,
		"position":   user.Position,
	}

	if user.Department != nil {
		response["department"] = fiber.Map{
			"id":   user.Department.ID,
			"name": user.Department.Name,
		}
	}

	return response
}
