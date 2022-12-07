package data

import (
	"log"

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
	err = Instance.AutoMigrate(&TaskSubmission{})
	if err != nil {
		log.Fatal(err)
		panic("Couldnt migrate DB")
	}
	log.Println("Database Migration Completed...")
}
