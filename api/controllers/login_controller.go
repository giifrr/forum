package controllers

import (
	"log"
	"net/http"

	"github.com/giifrr/forum/api/auth"
	"github.com/giifrr/forum/api/dto"
	"github.com/giifrr/forum/api/model"
	"github.com/giifrr/forum/api/security"
	"github.com/giifrr/forum/api/utils/formaterror"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(c *gin.Context) {
	errorList = map[string]string{}

	var input dto.LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
		return
	}

	user := model.User{
		Password: input.Password,
		Email: input.Email,
	}

	user.Prepare()
	errorMessages := user.Validate("login")
	log.Println(errorMessages)
	if len(errorMessages) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"errors": errorMessages,
		})
		return
	}

	userData, err := s.SignIn(user.Email, user.Password)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  formattedError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userData,
	})
}

func (s *Server) SignIn(email, password string) (map[string]interface{}, error) {
	var err error

	userData := make(map[string]interface{})

	user := model.User{}

	err = s.DB.Debug().Model(model.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		log.Println("this is the error getting the user: ", err)
		return nil, err
	}

	err = security.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("this is the error hashing the password:", err)
		return nil, err
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		log.Println("This is the error creating the token:", err)
		return nil, err
	}

	userData["token"] = token
	userData["id"] = user.ID
	userData["email"] = user.Email
	userData["avatar_path"] = user.AvatarPath
	userData["username"] = user.Username

	return userData, nil
}
