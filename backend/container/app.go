package container

import (
	"time"
	"training-plan-api/controller"
	"training-plan-api/repository"
	"training-plan-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"google.golang.org/api/calendar/v3"
	"gorm.io/gorm"
)

type AppDependencies struct {
	DepartmentController *controller.DepartmentController
	CourseController     *controller.CourseController
	AuthController       *controller.AuthController
}

func NewAppDependencies(
	db *gorm.DB,
	redis *redis.Client,
	validate *validator.Validate,
	calendarService *calendar.Service,
	location *time.Location,
) *AppDependencies {

	// ---------- Department ----------
	departmentRepo := repository.NewDepartmentRepositoryImpl(db)
	departmentService := service.NewDepartmentServiceImpl(departmentRepo, validate, redis)
	departmentController := controller.NewDepartmentController(departmentService)

	// ---------- Course ----------
	courseRepo := repository.NewCourseRepositoryImpl(db)
	courseService := service.NewCourseServiceImpl(
		courseRepo,
		redis,
		validate,
		calendarService,
		location,
	)
	courseController := controller.NewCourseController(courseService)

	// ---------- Auth ----------
	authController := controller.NewAuthController(db)

	return &AppDependencies{
		DepartmentController: departmentController,
		CourseController:     courseController,
		AuthController:       authController,
	}
}
