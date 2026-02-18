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
	TrainingPlanController     *controller.TrainingPlanController
	AuthController       *controller.AuthController
	UserController       *controller.UserController
	CertificateController *controller.CertificateController
	RecordController     *controller.RecordController
}

func NewAppDependencies(
	db *gorm.DB,
	redis *redis.Client,
	validate *validator.Validate,
	calendarService *calendar.Service,
	location *time.Location,
	storage helper.Storage,
) *AppDependencies {

		// ---------- Department ----------
	departmentRepo := repository.NewDepartmentRepositoryImpl(db)
	departmentService := service.NewDepartmentServiceImpl(departmentRepo, validate, redis)
	departmentController := controller.NewDepartmentController(departmentService)
		// ---------- User ----------
	userRepo := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepo, departmentRepo, validate)
	userController := controller.NewUserController(userService, db)

	// ---------- Record ----------
	recordRepo := repository.NewRecordRepositoryImpl(db)
	recordService := service.NewRecordServiceImpl(recordRepo, userRepo, validate)
	recordController := controller.NewRecordController(recordService)

	// ---------- Certificate ----------
	certificateRepo := repository.NewCertificateRepositoryImpl(db)
	certificateService := service.NewCertificateServiceImpl(certificateRepo, validate, storage)
	certificateController := controller.NewCertificateController(certificateService)



	// ---------- TrainingPlan ----------
	trainingPlanRepo := repository.NewTrainingPlanRepositoryImpl(db)
	trainingPlanService := service.NewTrainingPlanServiceImpl(
		trainingPlanRepo,
		redis,
		validate,
		calendarService,
		location,
	)
	trainingPlanController := controller.NewTrainingPlanController(trainingPlanService)

	// ---------- Auth ----------
	authController := controller.NewAuthController(db)


	return &AppDependencies{
		DepartmentController: departmentController,
		TrainingPlanController:     trainingPlanController,
		AuthController:       authController,
		UserController:       userController,
		CertificateController: certificateController,
		RecordController:     recordController,
	}
}