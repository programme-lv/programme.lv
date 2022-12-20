package main

import (
	"github.com/KrisjanisP/deikstra/service/api"
	"github.com/KrisjanisP/deikstra/service/data"
	"github.com/KrisjanisP/deikstra/service/scheduler"
)

func initDatabase(DBConnString string) {
	data.Connect(DBConnString)
	data.Migrate()
}

func main() {
	config := LoadAppConfig()
	initDatabase(config.DBConnString)

	scheduler := scheduler.CreateSchedulerServer()

	go scheduler.StartSchedulerServer(config.SchedulerPort)

	apiController := api.CreateAPIController(scheduler)
	apiController.StartAPIServer(config.APIPort)
}
