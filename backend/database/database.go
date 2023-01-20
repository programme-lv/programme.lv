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
	}
	log.Println("connected to database")
	err = instance.AutoMigrate(
		&models.Language{},
		&models.TaskSubmission{},
		&models.TaskSubmJob{},
		&models.ExecSubmission{},
	)
	if err != nil {
		return nil, err
	}
	log.Println("database migration completed")
	return instance, nil
}
