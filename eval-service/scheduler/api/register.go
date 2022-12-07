package api

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/scheduler/logic"
	"github.com/gorilla/mux"
)

func registerAPIRoutes(router *mux.Router) {
	// tasks
	router.HandleFunc("/tasks/list", listTasks).Methods("GET")
	router.HandleFunc("/tasks/info/{task_id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/create", createTask).Methods("POST")
	router.HandleFunc("/tasks/delete/{task_id}", deleteTask).Methods("DELETE")

	// submissions
	router.HandleFunc("/submissions/enqueue", enqueueSubmission).Methods("POST")
	router.HandleFunc("/submissions/list", listSubmissions).Methods("GET")
	router.HandleFunc("/submissions/info/{subm_id}", getSubmission).Methods("GET")

	// languages

	// execute
	router.HandleFunc("/execute/enqueue", enqueueExecution).Methods("POST")
	router.HandleFunc("/execute/list", listExecutions).Methods("GET")
	router.HandleFunc("/execute/info/{execution_id}", getExecution).Methods("GET")
}

func CreateAPIRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	registerAPIRoutes(router)
	return router
}

func StartAPIServer(APIPort int, router *mux.Router, scheduler *logic.Scheduler) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", APIPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("rest server listening at %v", lis.Addr())
	if err := http.Serve(lis, router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
