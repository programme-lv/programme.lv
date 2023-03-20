package main

import (
	"log"

	"github.com/KrisjanisP/deikstra/service/config"
	"github.com/KrisjanisP/deikstra/service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var languages = [...]models.Language{
	{ID: "C11", Name: "C11 (GNU GCC)", Filename: "main.c", CompileCmd: "gcc -std=c11 -o main main.c", ExecuteCmd: "./main"},
	{ID: "C++17", Name: "C++17 (GNU G++)", Filename: "main.cpp", CompileCmd: "g++ -std=c++17 -o main main.cpp", ExecuteCmd: "./main"},
	{ID: "Python3.10", Name: "Python 3.10", Filename: "main.py", ExecuteCmd: "python3.10 main.py"},
	{ID: "Java19", Name: "Java19 (OpenJDK)", Filename: "Main.java", CompileCmd: "javac Main.java", ExecuteCmd: "java Main"},
}

var taskTypes = [...]models.TaskType{
	{ID: "batch", Description: "todo"},
	{ID: "interactive", Description: "todo"},
	{ID: "simple", Description: "todo"},
}

var verdicts = [...]models.Verdict{
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

func main() {
	conf := config.LoadAppConfig()
	db, err := gorm.Open(postgres.Open(conf.DBConnString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migrateTables(db)
	addLanguages(db)
	addTaskTypes(db)
	addVerdicts(db)
}

func migrateTables(db *gorm.DB) {
	log.Println("migrating tables")
	err := db.AutoMigrate(
		&models.User{},
		&models.Language{},
		&models.TaskType{},
		&models.TaskTag{},
		&models.Task{},
		&models.TaskSubmission{},
		&models.TaskSubmEvaluation{},
		&models.ExecSubmission{},
		&models.TaskTest{},
		&models.TaskSubtask{},
		&models.MarkdownStatement{},
		&models.MDStatementExample{},
		&models.PDFStatement{},
		&models.TaskSubmEvalTest{},
	)

	if err != nil {
		panic(err)
	}
}

func addLanguages(db *gorm.DB) {
	log.Println("adding languages to database")
	tx := db.Begin()
	for _, lang := range languages {
		err := tx.Save(&lang).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		log.Printf("language %+v added to database", lang.Name)
	}
	tx.Commit()
}

func addTaskTypes(db *gorm.DB) {
	log.Println("adding task types to database")
	tx := db.Begin()
	for _, taskType := range taskTypes {
		err := tx.Save(&taskType).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		log.Printf("task type %+v added to database", taskType.ID)
	}
	tx.Commit()
}

func addVerdicts(db *gorm.DB) {
	log.Println("adding verdicts to database")
	tx := db.Begin()
	for _, verdict := range verdicts {
		err := tx.Save(&verdict).Error
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		log.Printf("verdict %+v added to database", verdict.ID)
	}

}
