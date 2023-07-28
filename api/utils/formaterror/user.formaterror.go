package formaterror

import (
	"strings"

	"github.com/giifrr/forum/api/dto"
	"github.com/giifrr/forum/api/model"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func Validate(action string, user model.User) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "login":
		{
			loginUserRequest := dto.LoginRequest{
				Password: user.Password,
				Email:    user.Email,
			}
			err = validate.Struct(loginUserRequest)
		}
	default:
		{
			createUserRequest := dto.CreateUserRequest{
				Username: user.Username,
				Password: user.Password,
				Email:    user.Email,
			}
			err = validate.Struct(createUserRequest)
		}
	}

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, error := range validationErrors {
			if error.Field() == "Email" {
				errorMessages["Invalid_email"] = "Invalid email or email cannot be empty"
			}
			if error.Field() == "Password" {
				errorMessages["Invalid_password"] = "Invalid password, make sure length of password is greater than 5"
			}
			if error.Field() == "Username" {
				errorMessages["Required_username"] = "Username cannot be empty"
			}
		}
	}

	return errorMessages
}
