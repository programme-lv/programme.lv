package database

import (
	"log"

	"github.com/KrisjanisP/deikstra/service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectAndMigrate(connectionString string) (*gorm.DB, error) {
	instance, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
		log.Fatal(err)
		panic("Couldnt connect to DB")
	}
	log.Println("Connected to Database...")
	err = instance.AutoMigrate(&models.TaskSubmission{})
	if err != nil {
		return nil, err
		log.Fatal(err)
		panic("Couldnt migrate DB")
	}
	log.Println("Database Migration Completed...")
	return instance, nil
}
