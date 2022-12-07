package main

import (
	"github.com/KrisjanisP/deikstra/service/scheduler/database"
	"google.golang.org/grpc"
)

func initDatabase(config SchedulerConfig) {
	database.Connect(config.DBConnString)
	database.Migrate()
}

func main() {
	config := LoadAppConfig()
	initDatabase(config)

	server := grpc.NewServer()

	go startSchedulerServer(config)

	startAPIServer(config)
}
