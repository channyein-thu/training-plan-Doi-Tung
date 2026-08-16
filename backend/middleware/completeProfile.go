package middleware

import (
	"training-plan-api/repository"

	"github.com/gofiber/fiber/v2"
)

func RequireProfileComplete(userRepo repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)

		user, err := userRepo.FindById(userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User not found",
			})
		}

		if !user.IsProfileComplete {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "profile incomplete",
			})
		}

		return c.Next()
	}
}