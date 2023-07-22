package controllers

func (s *Server) InitializeRoutes() {
	v1 := s.Router.Group("/api/v1")
	{

		// login route
		v1.POST("/login", s.Login)

		// User routes
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetUsers)
	}
}
