package main

import (
	"deikstra-service/controllers"
	"deikstra-service/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
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

func main() {
	// load configurations from config.json using Viper
	LoadAppConfig()

	// initialize database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// register routes
	RegisterProductRoutes(router)

	// start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
