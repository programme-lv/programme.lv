package controller

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/data"

	"github.com/KrisjanisP/deikstra/service/scheduler"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

type Controller struct {
	scheduler *scheduler.Scheduler
	database  *gorm.DB
	router    *mux.Router
	taskFS    *data.TaskFS
}

func CreateAPIController(scheduler *scheduler.Scheduler, database *gorm.DB, taskManager *data.TaskFS) *Controller {
	router := mux.NewRouter().StrictSlash(true)
	controller := Controller{
		scheduler: scheduler,
		router:    router,
		database:  database,
		taskFS:    taskManager,
	}
	controller.registerAPIRoutes()
	return &controller
}

func (c *Controller) StartAPIServer(APIPort int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", APIPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("rest server listening at %v", lis.Addr())
	if err := http.Serve(lis, c.router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
