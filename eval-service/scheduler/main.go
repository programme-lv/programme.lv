package main

import (
	"github.com/KrisjanisP/deikstra/service/scheduler/database"
)

func initDatabase(config WorkerConfig) {
	database.Connect(config.DBConnString)
	database.Migrate()
}

func main() {
	config := LoadAppConfig()

	initDatabase(config)
	go startWorkerServer(config)
	startAPIServer(config)
}
