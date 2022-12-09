package api

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/scheduler/logic"
	"github.com/gorilla/mux"
)

type APIController struct {
	scheduler *logic.Scheduler
	router    *mux.Router
}

func (c *APIController) registerAPIRoutes() {
	// tasks
	c.router.HandleFunc("/tasks/list", listTasks).Methods("GET")
	c.router.HandleFunc("/tasks/info/{task_id}", getTask).Methods("GET")
	c.router.HandleFunc("/tasks/create", createTask).Methods("POST")
	c.router.HandleFunc("/tasks/delete/{task_id}", deleteTask).Methods("DELETE")

	// submissions
	c.router.HandleFunc("/submissions/enqueue", c.enqueueSubmission).Methods("POST")
	c.router.HandleFunc("/submissions/list", listSubmissions).Methods("GET")
	c.router.HandleFunc("/submissions/info/{subm_id}", getSubmission).Methods("GET")

	// languages

	// execute
	c.router.HandleFunc("/execute/enqueue", c.enqueueExecution).Methods("POST")
	c.router.HandleFunc("/execute/list", listExecutions).Methods("GET")
	c.router.HandleFunc("/execute/info/{execution_id}", getExecution).Methods("GET")
}

func CreateAPIController(scheduler *logic.Scheduler) *APIController {
	router := mux.NewRouter().StrictSlash(true)
	controller := APIController{scheduler: scheduler, router: router}
	controller.registerAPIRoutes()
	return &controller
}

func (c *APIController) StartAPIServer(APIPort int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", APIPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("rest server listening at %v", lis.Addr())
	if err := http.Serve(lis, c.router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}