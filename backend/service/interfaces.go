package service

import (
	"mime/multipart"
	"training-plan-api/data/request"
	"training-plan-api/data/response"
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
	ManagerFindByDepartment(departmentID, page, pageSize int) (response.PaginatedResponse[response.UserListResponse], error)
}

type CertificateService interface {
	FindByCurrentUser(userID uint) ([]response.CertificateResponse, error)
	Upload(userID uint, req request.CreateCertificateRequest, file *multipart.FileHeader) error
	Delete(certificateID int, userID uint) error
	FindAllPending(	page int,limit int,) (response.PaginatedResponse[response.CertificateResponse], error)
	Approve(certificateID int) error
	Reject(certificateID int) error
	GetTrainingIDByRecordID(recordID int, userID uint) (int, error)
}

type RecordService interface {
	RegisterStaff(trainingPlanId uint, req request.RegisterStaffRequest) error
	FindById(id int) (response.RecordResponse, error)
	Update(id int, req request.UpdateRecordRequest) error
	Delete(id int) error
 	FindByManager(managerID uint,page int,limit int,) (response.PaginatedResponse[response.RecordResponse], error)	
	FindByUser(userID uint, page int, limit int) (response.PaginatedResponse[response.RecordResponse], error)
}