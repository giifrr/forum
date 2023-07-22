package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) InitializeRoutes() {
	v1 := s.Router.Group("/api/v1")
	{
		// Root path
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello, world!",
			})
		})

		// login route
		v1.POST("/login", s.Login)

		// User routes
		v1.POST("/users", s.CreateUser)
	}
}
