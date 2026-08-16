package middleware

import (
	"os"
	"strings"

	"training-plan-api/model"
	"training-plan-api/helper"
	"github.com/gofiber/fiber/v2"
)


func JWTProtected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authentication required",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid token format",
		})
	}

	tokenString := parts[1]

	claims, err := helper.VerifyAccessToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid or expired token",
		})
	}

	c.Locals("user_id", helper.ExtractUserID(claims))
	c.Locals("user_role", helper.ExtractUserRole(claims))

	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	role := c.Locals("user_role").(string)
	if role != string(model.RoleHRAdmin) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "Access denied. Admin only.",
		})
	}
	return c.Next()
}

func ManagerOnly(c *fiber.Ctx) error {
	role := c.Locals("user_role").(string)
	if role != string(model.RoleDepartmentManager) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "Access denied. Manager only.",
		})
	}
	return c.Next()
}

func StaffOnly(c *fiber.Ctx) error {
	role := c.Locals("user_role").(string)
	if role != string(model.RoleStaff) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "Access denied. Staff only.",
		})
	}
	return c.Next()
}

func AdminOrManager(c *fiber.Ctx) error {
	role := c.Locals("user_role").(string)
	if role != string(model.RoleHRAdmin) && role != string(model.RoleDepartmentManager) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "Access denied. Admin or Manager only.",
		})
	}
	return c.Next()
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
