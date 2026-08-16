package controller

import (
	"crypto/rand"
	"encoding/base64"
	"training-plan-api/helper"
	"training-plan-api/service"

	"github.com/gofiber/fiber/v2"
)

const oauthStateCookie = "oauth_state"

type AuthOAuthController struct {
	authOAuthService service.AuthOAuthService
}

func NewAuthOAuthController(authOAuthService service.AuthOAuthService) *AuthOAuthController {
	return &AuthOAuthController{authOAuthService: authOAuthService}
}

func (c *AuthOAuthController) GoogleLogin(ctx *fiber.Ctx) error {
	// The OAuth callback lands on the frontend origin, so the CSRF `state` must
	// be bound to a cookie on that origin (the frontend sets it and verifies it
	// on callback). The frontend passes the state it committed to here via the
	// `state` query param; we simply thread it into the Google auth URL so the
	// value Google echoes back matches the frontend's cookie.
	//
	// We keep a locally generated fallback only so a direct hit to this endpoint
	// without a state still produces a valid Google URL (that path is not
	// CSRF-protected and should not be linked to from the app).
	state := ctx.Query("state")
	if state == "" {
		generated, err := generateStateToken()
		if err != nil {
			return helper.InternalServerError("Failed to initialize OAuth login")
		}
		state = generated
	}

	loginURL := c.authOAuthService.GetGoogleLoginURL(state)
	return ctx.Redirect(loginURL, fiber.StatusTemporaryRedirect)
}

func (c *AuthOAuthController) GoogleExchange(ctx *fiber.Ctx) error {
	type request struct {
		Code string `json:"code"`
	}

	var req request
	if err := ctx.BodyParser(&req); err != nil || req.Code == "" {
		return helper.BadRequest("Missing code")
	}

	jwtToken, user, err := c.authOAuthService.HandleGoogleCallback(req.Code)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"accessToken": jwtToken,
		"isProfileComplete": user.IsProfileComplete,
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
	})
}

func generateStateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
