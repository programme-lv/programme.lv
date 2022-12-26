package api

import (
	"fmt"
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
}

func (c *Controller) registerAPIRoutes() {
	// tasks
	c.router.HandleFunc("/tasks/list", listTasks).Methods("GET")
	c.router.HandleFunc("/tasks/info/{task_id}", getTask).Methods("GET")
	c.router.HandleFunc("/tasks/create", createTask).Methods("POST")
	c.router.HandleFunc("/tasks/delete/{task_id}", deleteTask).Methods("DELETE")

	// submissions
	c.router.HandleFunc("/submissions/enqueue", c.enqueueSubmission).Methods("POST")
	c.router.HandleFunc("/submissions/list", c.listSubmissions).Methods("GET")
	c.router.HandleFunc("/submissions/info/{subm_id}", c.getSubmission).Methods("GET")
	c.router.HandleFunc("/submissions/subscribe", c.subscribeToResults).Methods("GET")

	// execute
	c.router.HandleFunc("/execute/enqueue", c.enqueueExecution).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/execute/communicate/{exec_id}", c.communicateWithExec)

	c.router.Use(mux.CORSMethodMiddleware(c.router))
}

func CreateAPIController(scheduler *scheduler.Scheduler, database *gorm.DB) *Controller {
	router := mux.NewRouter().StrictSlash(true)
	controller := Controller{
		scheduler: scheduler,
		router:    router,
		database:  database,
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
