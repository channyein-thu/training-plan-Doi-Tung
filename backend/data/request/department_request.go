package request

type CreateDepartmentRequest struct {
	Name string `validate:"required,min=3,max=100" json:"name"`
}
type UpdateDepartmentRequest struct {
	Name string `validate:"required,min=3,max=100" json:"name"`
}