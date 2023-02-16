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

	err = db.AutoMigrate(&models.Verdict{})
	if err != nil {
		return
	}

	verdicts := []models.Verdict{
		{ID: "IQS", Description: "In Queue State"},
		{ID: "ICS", Description: "In Compilation State"},
		{ID: "ITS", Description: "In Testing State"},
		{ID: "AC", Description: "Accepted"},
		{ID: "PT", Description: "Partially correct"},
		{ID: "WA", Description: "Wrong Answer"},
		{ID: "RE", Description: "Runtime Error"},
		{ID: "PE", Description: "Presentation Error"},
		{ID: "TLE", Description: "Time Limit Exceeded"},
		{ID: "MLE", Description: "Memory Limit Exceeded"},
		{ID: "ILE", Description: "Idleness limit exceeded"},
		{ID: "CE", Description: "Compilation Error"},
		{ID: "ISE", Description: "Internal Server Error"},
		{ID: "SV", Description: "Security Violation"},
		{ID: "IG", Description: "Ignored"},
		{ID: "RJ", Description: "Rejected"},
	}

	log.Println(db.Create(&verdicts).RowsAffected)
}
