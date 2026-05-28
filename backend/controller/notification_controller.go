package controller

import (
	"strconv"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
)

type NotificationController struct {
	service service.NotificationService
}

func NewNotificationController(service service.NotificationService) *NotificationController {
	return &NotificationController{service: service}
}

func (c *NotificationController) FindByCurrentUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	result, err := c.service.FindByUser(userID, page, limit)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status: "SUCCESS",
		Data:   result,
	})
}

func (c *NotificationController) CountUnread(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	count, err := c.service.CountUnread(userID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status: "SUCCESS",
		Data:   fiber.Map{"unread": count},
	})
}

func (c *NotificationController) MarkAsRead(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return helper.BadRequest("Invalid notification ID")
	}

	if err := c.service.MarkAsRead(uint(id), userID); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Notification marked as read",
	})
}

func (c *NotificationController) MarkAllRead(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	if err := c.service.MarkAllRead(userID); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "All notifications marked as read",
	})
}

// -------- Web Push --------

func (c *NotificationController) GetVAPIDPublicKey(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status: "SUCCESS",
		Data:   fiber.Map{"publicKey": c.service.GetVAPIDPublicKey()},
	})
}

func (c *NotificationController) SubscribePush(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	var req request.PushSubscribeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid subscription payload")
	}
	if req.Endpoint == "" || req.Keys.P256dh == "" || req.Keys.Auth == "" {
		return helper.BadRequest("Subscription is missing endpoint or keys")
	}

	userAgent := ctx.Get("User-Agent")
	if err := c.service.SubscribePush(userID, req.Endpoint, req.Keys.P256dh, req.Keys.Auth, userAgent); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Push subscription saved",
	})
}

func (c *NotificationController) SendTestPush(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)
	result := c.service.SendTestPush(userID)
	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status: "SUCCESS",
		Data:   result,
	})
}

func (c *NotificationController) UnsubscribePush(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	var req request.PushUnsubscribeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.BadRequest("Invalid unsubscribe payload")
	}
	if req.Endpoint == "" {
		return helper.BadRequest("Endpoint required")
	}

	if err := c.service.UnsubscribePush(userID, req.Endpoint); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Push subscription removed",
	})
}

func (c *NotificationController) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return helper.BadRequest("Invalid notification ID")
	}

	if err := c.service.Delete(uint(id), userID); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "SUCCESS",
		Message: "Notification deleted",
	})
}
