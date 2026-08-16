package response

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
}

type PaginatedResponse[T any] struct {
	Items []T           `json:"items"`
	Meta  PaginationMeta `json:"meta"`
}
