package dto

type LoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
