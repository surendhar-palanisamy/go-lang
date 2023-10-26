package database

import (
	"log"
	"os"

	"github.com/surendhar-palanisamy/gofiber-crud/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInstance struct {
	Db *gorm.DB
}

var Database dbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed To Connect", err.Error())
		os.Exit(2)
	}
	log.Println("Connected To DB")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{})
	Database = dbInstance{Db: db}
}
