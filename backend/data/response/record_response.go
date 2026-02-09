package response

import "time"

type RecordResponse struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"userId"`
	UserName   string    `json:"userName"`
	CourseID   uint      `json:"courseId"`
	CourseName string    `json:"courseName"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
