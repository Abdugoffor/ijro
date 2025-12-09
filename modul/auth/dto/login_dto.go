package auth_dto

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RefreshToken struct {
	Token string `json:"token" validate:"required"`
}