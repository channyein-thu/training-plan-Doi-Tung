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
	seed.SeedAdmin(db)/// for development purpose only////////////

	// Init Redis
	redisClient := config.NewRedisClient()

	app.Use(cors.New())

	app.Use(helmet.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	//  Validator
	validate := validator.New()

	calendarService := helper.NewGoogleCalendarService(context.Background())
	location := helper.LoadLocation()

	// Initialize storage
	storage := helper.NewLocalStorage(appConfig.UploadPath)

	deps := container.NewAppDependencies(
		db,
		redisClient,
		validate,
		calendarService,
		location,
		storage,
	)

	//  Routes
	router.RegisterRoutes(app, deps)

	// Serve static files for uploads
	app.Static("/uploads", appConfig.UploadPath)
	

	log.Fatal(app.Listen(":8080"))
}
