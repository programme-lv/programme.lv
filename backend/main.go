package main

import (
	"log"

	"github.com/KrisjanisP/deikstra/service/controller"
	"github.com/KrisjanisP/deikstra/service/data"
	"github.com/KrisjanisP/deikstra/service/scheduler"
)

func main() {
	config := LoadAppConfig()

	db, err := data.ConnectAndMigrate(config.DBConnString)
	if err != nil {
		log.Fatal(err)
	}

	sched := scheduler.CreateSchedulerServer()

	go sched.StartSchedulerServer(config.SchedulerPort)

	tm := data.CreateTaskManager(config.TasksDir)

	c := controller.CreateAPIController(sched, db, tm)
	c.StartAPIServer(config.APIPort)
}
