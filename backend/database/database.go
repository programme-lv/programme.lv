package database

import (
	"gorm.io/gorm"
	"log"

	"github.com/KrisjanisP/deikstra/service/models"
	"gorm.io/driver/postgres"
)

func ConnectAndMigrate(connStr string) (*gorm.DB, error) {
	instance, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
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
