package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/scheduler"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

type Controller struct {
	scheduler *scheduler.Scheduler
	database  *gorm.DB
	router    *mux.Router
	validate  *validator.Validate
}

func CreateAPIController(scheduler *scheduler.Scheduler, database *gorm.DB) *Controller {
	router := mux.NewRouter().StrictSlash(true)
	controller := Controller{
		scheduler: scheduler,
		router:    router,
		database:  database,
		validate:  validator.New(),
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
