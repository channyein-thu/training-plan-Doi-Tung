package service

import (
	"mime/multipart"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
	"training-plan-api/model"

	"github.com/xuri/excelize/v2"
)

type TrainingPlanService interface {
	Create(trainingPlan request.CreateTrainingPlanRequest) error
	Update(trainingPlanId int, trainingPlan request.UpdateTrainingPlanRequest) error
	Delete(trainingPlanId int) error
	FindById(trainingPlanId int) (response.TrainingPlanResponse, error)
	FindPaginated(page, pageSize int) (response.PaginatedResponse[response.TrainingPlanResponse], error)
}

type DepartmentService interface {
	Create(department request.CreateDepartmentRequest) error
	Update(departmentId int, department request.UpdateDepartmentRequest) error
	Delete(departmentId int) error
	FindById(departmentId int) (response.DepartmentResponse, error)
	FindPaginated(page, pageSize int) (response.PaginatedResponse[response.DepartmentResponse], error)
	FindDepartmentList() ([]response.DepartmentListItem, error)
}

type UserService interface {
	AdminCreate(req request.CreateUserRequest, creatorID uint) error
	AdminUpdate(userID uint, req request.UpdateUserRequest) error
	AdminDelete(userID uint) error
	AdminFindAll(page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error)
	AdminFindById(userID uint) (response.UserResponse, error)
	AdminFindAllForTable(params request.UserTableQueryParams) (response.PaginatedResponse[response.UserTableResponse], error)
	ManagerCreate(req request.ManagerCreateUserRequest, managerID uint, managerDepartmentID int) error
	ManagerUpdate(userID uint, req request.ManagerUpdateUserRequest, managerDepartmentID int) error
	ManagerFindByDepartment(departmentID, page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error)
	CompleteProfile(userID uint, req request.CompleteProfileRequest) error
}

type AuthOAuthService interface {
	GetGoogleLoginURL(state string) string
	HandleGoogleCallback(code string) (string, *model.User, error)
}

type CertificateService interface {
	FindByCurrentUser(userID uint) ([]response.CertificateResponse, error)
	Upload(userID uint, req request.CreateCertificateRequest, file *multipart.FileHeader) error
	Delete(certificateID int, userID uint) error
	FindAllPending(	page int,limit int,) (response.PaginatedResponse[response.CertificateResponse], error)
	Approve(certificateID int) error
	Reject(certificateID int) error
	// GetCertificateFilePath authorizes the caller and returns the on-disk
	// path of the certificate image so it can be streamed by an authenticated
	// handler (never served as an open static file).
	GetCertificateFilePath(certificateID int, callerID uint, callerRole string) (string, error)
}

type RecordService interface {
	RegisterStaff(trainingPlanId uint, req request.RegisterStaffRequest) error
	FindById(id int, callerID uint, callerRole string) (response.RecordResponseFinal, error)
	Update(id int, req request.UpdateRecordRequest, callerID uint, callerRole string) error
	Delete(id int, callerID uint, callerRole string) error
 	FindByManager(managerID uint,page int,limit int,) (response.PaginatedResponse[response.AdminRecordResponse], error)
	FindByUser(userID uint, page int, limit int) (response.PaginatedResponse[response.StaffRecordResponse], error)
	Search(req request.RecordFilterRequest) (response.PaginatedResponse[response.AdminRecordResponse], error)
	Export(req request.RecordFilterRequest) (*excelize.File, error)
}

type NotificationService interface {
	Create(userID uint, notifType model.NotificationType, title, message string) error
	FindByUser(userID uint, page, limit int) (response.PaginatedResponse[response.NotificationResponse], error)
	CountUnread(userID uint) (int64, error)
	MarkAsRead(id, userID uint) error
	MarkAllRead(userID uint) error
	Delete(id, userID uint) error

	// Web Push
	GetVAPIDPublicKey() string
	SubscribePush(userID uint, endpoint, p256dh, auth, userAgent string) error
	UnsubscribePush(userID uint, endpoint string) error
	SendTestPush(userID uint) PushTestResult
}

type PushTestResult struct {
	VAPIDConfigured    bool                     `json:"vapidConfigured"`
	SubscriptionsFound int                      `json:"subscriptionsFound"`
	Results            []PushSubscriptionResult `json:"results"`
}

type PushSubscriptionResult struct {
	EndpointPrefix string `json:"endpoint"`
	StatusCode     int    `json:"statusCode"`
	Error          string `json:"error,omitempty"`
}