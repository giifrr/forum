package api

import (
	"fmt"
	"os"

	"github.com/giifrr/forum/api/controllers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Errorln("sad .env not found")
	}

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func Run() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		log.Infoln("We are getting values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Infoln("Listening port on", apiPort)

	server.Run(apiPort)
}
