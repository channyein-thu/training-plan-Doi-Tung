package response

import "time"

type NotificationResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"userId"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	IsRead    bool   `json:"isRead"`
	CreatedAt time.Time `json:"createdAt"`
}
