package api

import (
	"fmt"
	"log"
	"os"

	"github.com/giifrr/forum/api/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("sad .env not found")
	}
}

func Run() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		log.Println("We are getting values")
	}
	
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Println("Listening port on", apiPort)

	server.Run(apiPort)
}
