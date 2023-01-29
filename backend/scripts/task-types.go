package main

import (
	"github.com/KrisjanisP/deikstra/service/config"
	"github.com/KrisjanisP/deikstra/service/database"
	"github.com/KrisjanisP/deikstra/service/models"
	"log"
)

func main() {
	conf := config.LoadAppConfig()
	db, err := database.Connect(conf.DBConnString)

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.TaskType{})
	if err != nil {
		return
	}

	types := []models.TaskType{
		{ID: "batch", Description: "idk"},
	}

	tx := db.Begin()
	for _, taskType := range types {
		err = tx.FirstOrCreate(&taskType).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		log.Printf("task type %+v added to database", taskType.ID)
	}
	tx.Commit()
}
