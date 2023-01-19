package controller

import "github.com/gorilla/mux"

func (c *Controller) registerAPIRoutes() {
	// tasks
	c.router.HandleFunc("/tasks/list", c.listTasks).Methods("GET", "OPTIONS")
	c.router.HandleFunc("/tasks/view/{task_code}", c.getTask).Methods("GET", "OPTIONS")
	c.router.HandleFunc("/tasks/create", c.createTask).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/tasks/delete/{task_code}", c.deleteTask).Methods("DELETE", "OPTIONS")
	c.router.HandleFunc("/tasks/statement/{task_code}/{filename}", c.getPDFStatement).Methods("GET", "OPTIONS")

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