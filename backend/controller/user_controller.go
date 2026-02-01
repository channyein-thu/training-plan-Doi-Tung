package controller

import (
	"strconv"

	"training-plan-api/data/request"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	userService service.UserService
	db          *gorm.DB
}

func NewUserController(userService service.UserService, db *gorm.DB) *UserController {
	return &UserController{
		userService: userService,
		db:          db,
	}
}

// ============== ADMIN ENDPOINTS ==============

func (uc *UserController) AdminCreate(c *fiber.Ctx) error {
	var req request.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	creatorID := c.Locals("user_id").(uint)

	if err := uc.userService.AdminCreate(req, creatorID); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
	})
}

func (uc *UserController) AdminUpdate(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.BadRequest("Invalid user ID")
	}

	var req request.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	if err := uc.userService.AdminUpdate(uint(userID), req); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
	})
}

func (uc *UserController) AdminDelete(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.BadRequest("Invalid user ID")
	}

	if err := uc.userService.AdminDelete(uint(userID)); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}

func (uc *UserController) AdminFindAll(c *fiber.Ctx) error {
	params := request.UserTableQueryParams{
		Search:       c.Query("search"),
		DepartmentID: 0,
		Status:       c.Query("status"),
		Page:         1,
		Limit:        10,
		SortBy:       c.Query("sortBy", "employee_id"),
		SortOrder:    c.Query("sortOrder", "asc"),
	}

	if deptID, err := strconv.Atoi(c.Query("departmentId", "0")); err == nil {
		params.DepartmentID = deptID
	}
	if page, err := strconv.Atoi(c.Query("page", "1")); err == nil {
		params.Page = page
	}
	if limit, err := strconv.Atoi(c.Query("limit", "10")); err == nil {
		params.Limit = limit
	}

	result, err := uc.userService.AdminFindAllForTable(params)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

func (uc *UserController) AdminFindById(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.BadRequest("Invalid user ID")
	}

	user, err := uc.userService.AdminFindById(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

// ============== MANAGER ENDPOINTS ==============

func (uc *UserController) ManagerCreate(c *fiber.Ctx) error {
	var req request.ManagerCreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid request body")
	}

	managerID := c.Locals("user_id").(uint)

	managerDepartmentID, err := uc.getManagerDepartmentID(managerID)
	if err != nil {
		return helper.InternalServerError("Failed to get manager department")
	}

	if err := uc.userService.ManagerCreate(req, managerID, managerDepartmentID); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
	})
}

func (uc *UserController) ManagerFindDepartmentUsers(c *fiber.Ctx) error {
	managerID := c.Locals("user_id").(uint)

	managerDepartmentID, err := uc.getManagerDepartmentID(managerID)
	if err != nil {
		return helper.InternalServerError("Failed to get manager department")
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))

	result, err := uc.userService.ManagerFindByDepartment(managerDepartmentID, page, pageSize)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    result,
	})
}

func (uc *UserController) getManagerDepartmentID(managerID uint) (int, error) {
	var departmentID int
	err := uc.db.Table("users").
		Select("department_id").
		Where("id = ?", managerID).
		Scan(&departmentID).Error
	return departmentID, err
}
