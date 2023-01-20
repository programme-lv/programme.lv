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

	tm := database.NewTaskManager(config.TasksDir)

	sched := scheduler.NewScheduler(db, tm)

	go sched.StartSchedulerServer(config.SchedulerPort)

	c := controller.CreateAPIController(sched, db, tm)
	c.StartAPIServer(config.APIPort)
}
