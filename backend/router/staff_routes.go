package router

import (
	"training-plan-api/container"

	"github.com/gofiber/fiber/v2"
)

func StaffRoutes(r fiber.Router, deps *container.AppDependencies) {

	r.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Training Plan API is running",
		})
	})

	

	// Dashboard stats
	r.Get("/dashboard-stats", deps.DashboardController.StaffStats)

	// // Records (own)
	r.Get("/records", deps.RecordController.FindByCurrentUser)
	r.Get("/records/:id", deps.RecordController.FindById)

	// // Certificates
	r.Get("/certificates", deps.CertificateController.FindByCurrentUser) // only approved certificates
	r.Post("/certificates", deps.CertificateController.Upload)
	r.Delete("/certificates/:id", deps.CertificateController.Delete)

	// Notifications
	r.Get("/notifications", deps.NotificationController.FindByCurrentUser)
	r.Get("/notifications/unread-count", deps.NotificationController.CountUnread)
	r.Get("/notifications/vapid-public-key", deps.NotificationController.GetVAPIDPublicKey)
	r.Post("/notifications/push-subscribe", deps.NotificationController.SubscribePush)
	r.Post("/notifications/push-unsubscribe", deps.NotificationController.UnsubscribePush)
	r.Post("/notifications/push-test", deps.NotificationController.SendTestPush)
	r.Put("/notifications/read-all", deps.NotificationController.MarkAllRead)
	r.Put("/notifications/:id/read", deps.NotificationController.MarkAsRead)
	r.Delete("/notifications/:id", deps.NotificationController.Delete)
}
