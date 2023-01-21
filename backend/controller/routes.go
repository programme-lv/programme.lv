package controller

import "github.com/gorilla/mux"

func (c *Controller) registerAPIRoutes() {
	// tasks
	c.router.HandleFunc("/tasks/list", c.listTasks).Methods("GET", "OPTIONS")
	c.router.HandleFunc("/tasks/view/{task_code}", c.getTask).Methods("GET", "OPTIONS")
	c.router.HandleFunc("/tasks/create", c.createTask).Methods("POST", "OPTIONS")

	// submissions
	c.router.HandleFunc("/submissions/enqueue", c.enqueueSubmission).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/submissions/list", c.listSubmissions).Methods("GET")
	c.router.HandleFunc("/submissions/view/{subm_id}", c.getSubmission).Methods("GET")
	c.router.HandleFunc("/submissions/subscribe", c.subscribeToResults).Methods("GET")

	// execute
	c.router.HandleFunc("/execute/enqueue", c.enqueueExecution).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/execute/communicate/{exec_id}", c.communicateWithExec)

	c.router.Use(mux.CORSMethodMiddleware(c.router))
}
