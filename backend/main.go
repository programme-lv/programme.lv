package main

import (
	"github.com/KrisjanisP/deikstra/service/config"
	"log"

	"github.com/KrisjanisP/deikstra/service/controller"
	"github.com/KrisjanisP/deikstra/service/database"
	"github.com/KrisjanisP/deikstra/service/scheduler"
)

func main() {
	conf := config.LoadAppConfig()

	db, err := database.Connect(conf.DBConnString)
	if err != nil {
		log.Fatal(err)
	}

	sched := scheduler.NewScheduler(db)

	go sched.StartSchedulerServer(conf.SchedulerPort)

	c := controller.CreateAPIController(sched, db)
	c.StartAPIServer(conf.APIPort)
}
