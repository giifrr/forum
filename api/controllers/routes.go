package controllers

import "github.com/giifrr/forum/api/middleware"

func (s *Server) InitializeRoutes() {
	v1 := s.Router.Group("/api/v1")
	{

		// login route
		v1.POST("/login", s.Login)

		// User routes
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetUsers)
		v1.GET("/users/:id", s.GetUser)
		v1.DELETE("/users/:id", middleware.AuthMiddleware(), s.DeleteUser)

	}
}
