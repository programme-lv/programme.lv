package main

import (
	"github.com/KrisjanisP/deikstra/service/scheduler/api"
	"github.com/KrisjanisP/deikstra/service/scheduler/data"
	"github.com/KrisjanisP/deikstra/service/scheduler/logic"
)

func initDatabase(DBConnString string) {
	data.Connect(DBConnString)
	data.Migrate()
}

func main() {
	config := LoadAppConfig()
	initDatabase(config.DBConnString)

	apiRouter := api.CreateAPIRouter()
	server, scheduler := logic.CreateSchedulerServer()

	go logic.StartSchedulerServer(config.SchedulerPort, server)

	api.StartAPIServer(config.APIPort, apiRouter, scheduler)
}
