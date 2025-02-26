package logout

type LogoutRequest struct {
	Token string `json:"token" validate:"required,jwt"`
}

type LogoutResponse struct {
	UserUUID string `json:"user_uuid"`
}
