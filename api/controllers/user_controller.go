package controllers

import (
	"log"
	"net/http"

	"github.com/giifrr/forum/api/dto"
	"github.com/giifrr/forum/api/model"
	"github.com/giifrr/forum/api/utils/formaterror"
	"github.com/gin-gonic/gin"
)

func (server *Server) CreateUser(c *gin.Context) {
	errorList = map[string]string{}

	var input dto.CreateUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errorList,
		})
		return
	}
	user := model.User{
		Password: input.Password,
		Email: input.Email,
		Username: input.Username,
	}
	user.Prepare()

	errorMessages := user.Validate("")
	log.Println(errorMessages)
	if len(errorMessages) > 0 {
		errorList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"errors": errorList,
		})
		return
	}
	log.Println("Masih ada")

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errorList = formattedError
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errorList,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"response": userCreated,
	})
}
