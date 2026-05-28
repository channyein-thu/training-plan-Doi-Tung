package main

import (
	"context"
	"log"
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
	})

	//  Load config
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(" Cannot load config:", err)
	}

	//  DB and migration
	db := config.ConnectionDB(&appConfig)
	seed.SeedAdmin(db)/// for development purpose only


	app.Use(cors.New(cors.Config{
    AllowOrigins:     "http://localhost:3000",
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
		// Skip rate limiting for static uploads so file viewing never gets throttled
		Next: func(c *fiber.Ctx) bool {
			return len(c.Path()) >= 8 && c.Path()[:8] == "/uploads"
		},
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

	// Serve static files for uploads
	app.Static("/uploads", appConfig.UploadPath)
	

	log.Fatal(app.Listen(":8080"))
}
