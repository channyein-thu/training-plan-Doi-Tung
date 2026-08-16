package request

type PushSubscribeRequest struct {
	Endpoint string `json:"endpoint" validate:"required"`
	Keys     struct {
		P256dh string `json:"p256dh" validate:"required"`
		Auth   string `json:"auth" validate:"required"`
	} `json:"keys"`
}

type PushUnsubscribeRequest struct {
	Endpoint string `json:"endpoint" validate:"required"`
}
