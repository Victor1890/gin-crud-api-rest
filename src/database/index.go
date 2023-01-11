package database

import (
	"log"

	"github.com/Victor1890/gin-crud-api-rest/src/config"
	"github.com/Victor1890/gin-crud-api-rest/src/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {

	var config = config.Config()

	db, dbErr := gorm.Open(postgres.Open(config.Db_connect), &gorm.Config{})

	if dbErr != nil {
		log.Fatalln("Error with connect db")
		panic(dbErr)
	}

	db.Migrator().DropTable(&models.Product{}, &models.File{})
	migrationError := db.AutoMigrate(&models.Product{}, &models.File{})

	if migrationError != nil {
		log.Fatalln("Error with migration")
		panic(migrationError)
	}

	DB = db

}
