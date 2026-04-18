package response

import "training-plan-api/model"

type UserResponse struct {
	ID           uint                `json:"id"`
	EmployeeID   string              `json:"employeeID"`
	Name         string              `json:"name"`
	Email        string              `json:"email"`
	Phone        string              `json:"phone"`
	DepartmentID int                 `json:"departmentId"`
	Department   *DepartmentResponse `json:"department,omitempty"`
	Certificate  []UserCertificateResponse `json:"certificates,omitempty"`
	Role         model.Role          `json:"role"`
	Status       model.UserStatus    `json:"status"`
	Position     string              `json:"position"`
	Avatar       string              `json:"avatar,omitempty"`
	Provider     string              `json:"provider,omitempty"`
	CreatedBy    model.CreatedByType `json:"createdBy"`
	CreatedAt    int64               `json:"createdAt"`
	UpdatedAt    int64               `json:"updatedAt"`
}

type UserListResponse struct {
	ID           uint             `json:"id"`
	EmployeeID   string           `json:"employeeID"`
	Name         string           `json:"name"`
	Email        string           `json:"email"`
	Phone        string           `json:"phone"`
	DepartmentID int              `json:"departmentId"`
	Department   string           `json:"department,omitempty"`
	Role         model.Role       `json:"role"`
	Status       model.UserStatus `json:"status"`
	Position     string           `json:"position"`
}

type UserTableResponse struct {
	ID             uint   `json:"id"`
	EmployeeID     string `json:"employeeId"`
	FullName       string `json:"fullName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	DepartmentID   int    `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
	Position	   string `json:"position"`
	Role 		 model.Role `json:"role"`
	JobRole        string `json:"jobRole"`
	Status         string `json:"status"`
	IsManager      bool   `json:"isManager"`
}

func ToUserTableResponse(user model.User) UserTableResponse {
	deptName := ""
	if user.Department != nil {
		deptName = user.Department.Name
	}

	status := string(user.Status)
	if status == "Active" {
		status = "ACTIVE"
	} else if status == "Inactive" {
		status = "INACTIVE"
	}

	return UserTableResponse{
		ID:             user.ID,
		EmployeeID:     user.EmployeeID,
		FullName:       user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		DepartmentID:   user.DepartmentID,
		DepartmentName: deptName,
		Position:	   user.Position,
		Role: 		 user.Role,
		JobRole:        user.Position,
		Status:         status,
		IsManager:      user.Role == model.RoleDepartmentManager,
	}
}

func ToUserTableResponses(users []model.User) []UserTableResponse {
	responses := make([]UserTableResponse, len(users))
	for i, user := range users {
		responses[i] = ToUserTableResponse(user)
	}
	return responses
}

func ToUserResponse(user model.User) UserResponse {
	resp := UserResponse{
		ID:           user.ID,
		EmployeeID:   user.EmployeeID,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		DepartmentID: user.DepartmentID,
		Role:         user.Role,
		Status:       user.Status,
		Position:     user.Position,
		Avatar:       user.Avatar,
		Provider:     user.Provider,
		CreatedBy:    user.CreatedBy,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
	if user.Department != nil {
		resp.Department = &DepartmentResponse{
			ID:   user.Department.ID,
			Name: user.Department.Name,
			Division: user.Department.Division,
		}
	}
	if user.Certificates != nil {
		certs := make([]UserCertificateResponse, len(user.Certificates))
		for i, cert := range user.Certificates {
			certs[i] = UserCertificateResponse{
				ID:          cert.ID,
				TrainingName: cert.Training.Name,
				Image:       cert.Image,
				Description: cert.Description,
				Status:      string(cert.Status),
			}
		}
		resp.Certificate = certs
	}
	return resp
}

func ToUserListResponse(user model.User) UserListResponse {
	deptName := ""
	if user.Department != nil {
		deptName = user.Department.Name
	}
	return UserListResponse{
		ID:           user.ID,
		EmployeeID:   user.EmployeeID,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		DepartmentID: user.DepartmentID,
		Department:   deptName,
		Role:         user.Role,
		Status:       user.Status,
		Position:     user.Position,
	}
}

func ToUserListResponses(users []model.User) []UserListResponse {
	responses := make([]UserListResponse, len(users))
	for i, user := range users {
		responses[i] = ToUserListResponse(user)
	}
	return responses
}
