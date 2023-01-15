package data

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
	}
	log.Println("connected to data")
	err = instance.AutoMigrate(
		&models.TaskSubmission{},
	)
	if err != nil {
		return nil, err
	}
	log.Println("data migration completed")
	return instance, nil
}
