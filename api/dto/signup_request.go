package dto

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,max=20,min=3"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}
