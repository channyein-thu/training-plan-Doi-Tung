package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"
	"training-plan-api/config"
	"training-plan-api/container"
	"training-plan-api/helper"
	"training-plan-api/middleware"
	"training-plan-api/router"
	"training-plan-api/seed"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		// Caps the entire request body (including multipart file uploads).
		// Keep in sync with helper.MaxCertificateFileSize — this is a hard
		// ceiling enforced before our handlers ever run.
		BodyLimit: 6 * 1024 * 1024, // 6MB
	})

	//  Load config
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(" Cannot load config:", err)
	}

	// Fail fast if the JWT signing key is missing or weak. The token helpers
	// read JWT_SECRET via os.Getenv, but viper's AutomaticEnv only populates the
	// config struct — it never exports into the process environment. Export it
	// here (so os.Getenv works under both `go run` and Docker) and refuse to
	// start with an empty/short key, which would otherwise let anyone forge
	// tokens signed with "".
	if len(appConfig.JWTSecret) < 32 {
		log.Fatal(" JWT_SECRET is missing or too short (need at least 32 characters). Set it in app.env or the environment.")
	}
	if err := os.Setenv("JWT_SECRET", appConfig.JWTSecret); err != nil {
		log.Fatal(" Failed to export JWT_SECRET:", err)
	}

	//  DB and migration
	db := config.ConnectionDB(&appConfig)

	// Seed the default admin ONLY in development, and only when an admin
	// password is explicitly provided. This prevents a production deployment
	// from ever booting with a well-known default admin account.
	if strings.EqualFold(appConfig.GoEnv, "development") {
		seed.SeedAdmin(db, appConfig.AdminSeedPassword)
	}


	// Allowed CORS origins come from ALLOWED_ORIGINS (comma-separated) so the
	// real admin/manager domains can be configured per environment. Falls back
	// to the local dev frontend.
	allowedOrigins := appConfig.AllowedOrigins
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000"
	}

	app.Use(cors.New(cors.Config{
    AllowOrigins:     allowedOrigins,
    AllowCredentials: true,
}))

	app.Use(helmet.New(helmet.Config{
    CrossOriginResourcePolicy: "cross-origin",
}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               300,
		Expiration:        60 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		// Return JSON so the frontend's response.json() never blows up
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"status":  "ERROR",
				"message": "Too many requests. Please slow down and try again in a moment.",
			})
		},
	}))
	//  Validator
	validate := validator.New()

	calendarService := helper.NewGoogleCalendarService(context.Background())
	location := helper.LoadLocation()

	// Initialize storage
	storage := helper.NewLocalStorage(appConfig.UploadPath)

	deps := container.NewAppDependencies(
		db,
		validate,
		calendarService,
		location,
		storage,
		appConfig,
	)

	//  Routes
	router.RegisterRoutes(app, deps)

	// NOTE: certificate files are intentionally NOT served via app.Static.
	// They contain personal data and are streamed only through the
	// authenticated GET /api/v1/certificates/:id/file endpoint.

	port := appConfig.Port
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
