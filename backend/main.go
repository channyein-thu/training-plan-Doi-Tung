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

	deps := container.NewAppDependencies(
		db,
		redisClient,
		validate,
		calendarService,
		location,
	)

	//  Routes
	router.RegisterRoutes(app, deps)
	

	log.Fatal(app.Listen(":8080"))
}


// //  Dependency Injection
// 	departmentRepository := repository.NewDepartmentRepositoryImpl(db)
// 	departmentService := service.NewDepartmentServiceImpl(departmentRepository, validate, redisClient)
// 	departmentController := controller.NewDepartmentController(departmentService)

// 	// Course Injection
// 	courseRepo := repository.NewCourseRepositoryImpl(db)
// 	courseService := service.NewCourseServiceImpl(
// 		courseRepo,
// 		redisClient,
// 		validate,
// 		calendarService,
// 		location,
// 	)
// 	courseController := controller.NewCourseController(courseService)

// 	// Auth Injection
// 	authController := controller.NewAuthController(db)

// 	//  Routes
// 	api := app.Group("/api/v1")

// 	api.Get("/healthchecker", func(c *fiber.Ctx) error {
// 		return c.Status(200).JSON(fiber.Map{
// 			"status":  "success",
// 			"message": "Training Plan API is running",
// 		})
// 	})

// 	router.DepartmentRoutes(api, departmentController)
// 	router.CourseRoutes(api, courseController)
// 	router.AuthRoutes(api, authController)