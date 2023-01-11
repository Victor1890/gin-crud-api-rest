package config

import (
	"log"
	"os"

	"github.com/Victor1890/gin-crud-api-rest/src/interfaces"
	"github.com/joho/godotenv"
)

func Config() *interfaces.ConfigI {

	db_conn := getEnvironment("DB_CONNECTION")
	releaseMode := getEnvironment("GIN_MODE")

	return &interfaces.ConfigI{
		Db_connect:  db_conn,
		ReleaseMode: releaseMode,
	}

}

func getEnvironment(str string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loafing .env file")
	}

	return os.Getenv(str)
}
