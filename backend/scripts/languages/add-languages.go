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

	err = db.AutoMigrate(&models.Language{})
	if err != nil {
		return
	}

	languages := []models.Language{
		{ID: "C11", Name: "C11 (GNU GCC)", Filename: "main.c", CompileCmd: "gcc -std=c11 -o main main.c", ExecuteCmd: "./main"},
		{ID: "C++17", Name: "C++17 (GNU G++)", Filename: "main.cpp", CompileCmd: "g++ -std=c++17 -o main main.cpp", ExecuteCmd: "./main"},
		{ID: "Python3.10", Name: "Python 3.10", Filename: "main.py", ExecuteCmd: "python3 main.py"},
		{ID: "Java19", Name: "Java19 (OpenJDK)", Filename: "Main.java", CompileCmd: "javac Main.java", ExecuteCmd: "java Main"},
	}

	tx := db.Begin()
	for _, lang := range languages {
		err = tx.Save(&lang).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		log.Printf("language %+v added to database", lang.Name)
	}
	tx.Commit()
}
