package middleware

import (
	"os"
	"strings"

	"training-plan-api/model"
	"training-plan-api/utils"

	"github.com/gofiber/fiber/v2"
)


func JWTProtected(c *fiber.Ctx) error {
	var tokenString string

	tokenString = utils.GetAccessTokenFromCookie(c)

	if tokenString == "" {
		authHeader := c.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			}
		}
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authentication required",
		})
	}

	claims, err := utils.VerifyAccessToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid or expired token",
		})
	}

	c.Locals("user_id", utils.ExtractUserID(claims))
	c.Locals("user_role", utils.ExtractUserRole(claims))

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
