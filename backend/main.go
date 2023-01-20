package main

import (
	"log"

	"github.com/KrisjanisP/deikstra/service/controller"
	"github.com/KrisjanisP/deikstra/service/database"
	"github.com/KrisjanisP/deikstra/service/scheduler"
)

func main() {
	config := LoadAppConfig()

	db, err := database.ConnectAndMigrate(config.DBConnString)
	if err != nil {
		log.Fatal(err)
	}

	sched := scheduler.NewScheduler()

	go sched.StartSchedulerServer(config.SchedulerPort)

	tm := database.CreateTaskManager(config.TasksDir)

	c := controller.CreateAPIController(sched, db, tm)
	c.StartAPIServer(config.APIPort)
}
