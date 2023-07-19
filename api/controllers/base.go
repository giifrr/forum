package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/giifrr/forum/api/middleware"
	"github.com/giifrr/forum/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var errorList = make(map[string]string)

func (server *Server) Initialize(Dbdriver, Dbuser, Dbpassword, Dbport, Dbhost, Dbname string) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", Dbhost, Dbuser, Dbpassword, Dbname, Dbport)
	// log.Println("Hello Bang")
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Cannot connect to database")
		log.Fatal("This is the error connecting to postgres:", err)
	} else {
		log.Println("Connected to database successfully")
	}

	server.DB.Debug().AutoMigrate(&model.User{})

	server.Router = gin.Default()
	server.Router.Use(middleware.CORSMiddleware())

	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
