package controllers

import (
	"log"
	"net/http"
	"strconv"

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
		Email:    input.Email,
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
		"status":   http.StatusCreated,
		"response": userCreated,
	})
}

func (s *Server) GetUsers(c *gin.Context) {
	// clear previous error if any
	errorList = map[string]string{}

	user := model.User{}

	users, err := user.FindAllUsers(s.DB)

	if err != nil {
		errorList["No_users"] = "No user found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"errors": errorList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": users,
	})
}

func (s *Server) GetUser(c *gin.Context) {
	errorList = map[string]string{}

	userId := c.Param("id")

	uid, err := strconv.Atoi(userId)
	if err != nil {
		errorList["Invalid_request"] = "Invalid request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"errors": errorList,
		})
		return
	}

	user := model.User{}

	userGotten, err := user.FindUserById(s.DB, uint32(uid))
	if err != nil {
		errorList["No_user"] = "No user found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"errors": errorList,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"response": userGotten,
	})
}
