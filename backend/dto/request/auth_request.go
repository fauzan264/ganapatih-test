package request

type RegisterRequest struct {
	Username				string `json:"username" validate:"required"`
	Password				string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Username				string `json:"username" validate:"required"`
	Password				string `json:"password" validate:"required"`
}

type GetUser struct {
	ID							int `json:"id"`
}