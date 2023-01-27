package main

import (
	"github.com/KrisjanisP/deikstra/service/config"
	"github.com/KrisjanisP/deikstra/service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	conf := config.LoadAppConfig()
	db, err := gorm.Open(postgres.Open(conf.DBConnString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("connected to database")
	err = db.AutoMigrate(
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
		&models.TaskSubmJobTest{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("database migration completed")
}
