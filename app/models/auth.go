package models

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required" enums:"ADMIN,USER"`
}
type LoginResponse struct {
	UserId      string `json:"userId"`
	AccessToken string `json:"token"`
}

type RegisterInput struct {
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Token           string `json:"token" validate:"required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type SendEmailOtpInput struct {
	Email string `json:"email" validate:"required"`
}

type SendEmailResponse struct {
	Token string `json:"token"`
	// todo: remove otp
	Otp string `json:"otp"`
}

type VerifyOtpInput struct {
	Token string `json:"token" validate:"required"`
	Otp   string `json:"otp" validate:"required"`
	Email string `json:"email" validate:"required"`
}
type VerifyOtpResponse struct {
	Token string `json:"token"`
}
