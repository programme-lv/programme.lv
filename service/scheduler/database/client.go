package database

import (
	"log"

	"github.com/KrisjanisP/deikstra/service/scheduler/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Couldnt connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	err = Instance.AutoMigrate(&models.TaskSubmission{})
	if err != nil {
		log.Fatal(err)
		panic("Couldnt migrate DB")
	}
	log.Println("Database Migration Completed...")
}
