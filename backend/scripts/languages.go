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
	languages := []models.Language{
		{ID: "C11", Name: "C11 (GNU GCC)"},
		{ID: "C++17", Name: "C++17 (GNU G++)"},
		{ID: "Python3.10", Name: "Python 3.10"},
		{ID: "Java19", Name: "Java19 (OpenJDK)"},
	}

	tx := db.Begin()
	for _, lang := range languages {
		err = tx.FirstOrCreate(&lang, models.Language{Name: lang.Name}).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		log.Printf("language %+v added to database", lang.Name)
	}
	tx.Commit()
}
