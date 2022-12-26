package main

import (
	"log"

	"github.com/KrisjanisP/deikstra/service/api"
	"github.com/KrisjanisP/deikstra/service/database"
	"github.com/KrisjanisP/deikstra/service/scheduler"
)

func main() {
	config := LoadAppConfig()

	db, err := database.ConnectAndMigrate(config.DBConnString)
	if err != nil {
		log.Fatal(err)
	}

	sched := scheduler.CreateSchedulerServer()

	go sched.StartSchedulerServer(config.SchedulerPort)

	a := api.CreateAPIController(sched, db)
	a.StartAPIServer(config.APIPort)
}
