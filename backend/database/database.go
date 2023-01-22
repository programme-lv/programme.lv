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
		&models.User{},
		&models.Language{},
		&models.TaskSubmission{},
		&models.TaskSubmJob{},
		&models.ExecSubmission{},
		&models.Task{},
		&models.Test{},
		&models.Subtask{},
		&models.Tag{},
		&models.MarkdownStatement{},
		&models.MDStatementExample{},
		&models.PDFStatement{},
	)
	if err != nil {
		return nil, err
	}
	log.Println("database migration completed")
	return instance, nil
}
