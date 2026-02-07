package container

import (
	"time"
	"training-plan-api/controller"
	"training-plan-api/helper"
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
	UserController       *controller.UserController
	CertificateController *controller.CertificateController
}

func NewAppDependencies(
	db *gorm.DB,
	redis *redis.Client,
	validate *validator.Validate,
	calendarService *calendar.Service,
	location *time.Location,
	storage helper.Storage,
) *AppDependencies {
		// ---------- Certificate ----------
	certificateRepo := repository.NewCertificateRepositoryImpl(db)
	certificateService := service.NewCertificateServiceImpl(certificateRepo, validate, storage)
	certificateController := controller.NewCertificateController(certificateService)

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

	// ---------- User ----------
	userRepo := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepo, departmentRepo, validate)
	userController := controller.NewUserController(userService, db)

	return &AppDependencies{
		DepartmentController: departmentController,
		CourseController:     courseController,
		AuthController:       authController,
		UserController:       userController,
		CertificateController: certificateController,
}
}