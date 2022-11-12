package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/scheduler/controllers"
	"github.com/gorilla/mux"
)

func registerAPIRoutes(router *mux.Router) {
	// tasks
	router.HandleFunc("/tasks/list", controllers.ListTasks).Methods("GET")
	router.HandleFunc("/tasks/info/{task_id}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/create", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/delete/{task_id}", controllers.DeleteTask).Methods("DELETE")

	// submissions
	router.HandleFunc("/submissions/enqueue", controllers.EnqueueSubmission).Methods("POST")
	router.HandleFunc("/submissions/list", controllers.ListSubmissions).Methods("GET")
	router.HandleFunc("/submissions/info/{subm_id}", controllers.GetSubmission).Methods("GET")

	// languages

	// execute
	router.HandleFunc("/execute/enqueue", controllers.EnqueueExecution).Methods("POST")
	router.HandleFunc("/execute/list", controllers.ListExecutions).Methods("GET")
	router.HandleFunc("/execute/info/{execution_id}", controllers.GetExecution).Methods("GET")
}

func startAPIServer(config WorkerConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.APIPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	router := mux.NewRouter().StrictSlash(true)
	registerAPIRoutes(router)
	log.Printf("rest server listening at %v", lis.Addr())
	if err := http.Serve(lis, router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
