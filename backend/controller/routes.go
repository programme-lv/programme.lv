package controller

import "github.com/gorilla/mux"

func (c *Controller) registerAPIRoutes() {
	// users
	c.router.HandleFunc("/users/list", c.listUsers).Methods("GET", "OPTIONS")
	c.router.HandleFunc("/users/view/{user_id}", c.getUser).Methods("GET", "OPTIONS")
	c.router.HandleFunc("/users/register", c.createUser).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/users/login", c.loginUser).Methods("POST", "OPTIONS")
	//c.router.HandleFunc("/users/delete/{user_id}", c.deleteUser).Methods("DELETE", "OPTIONS")

	// tasks
	c.router.HandleFunc("/tasks/list", c.listTasks).Methods("GET", "OPTIONS")
	c.router.HandleFunc("/tasks/view/{task_id}", c.getTask).Methods("GET", "OPTIONS")
	//c.router.HandleFunc("/tasks/create", c.createTask).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/tasks/import", c.importTask).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/tasks/delete/{task_id}", c.deleteTask).Methods("DELETE", "OPTIONS")

	// submissions
	c.router.HandleFunc("/submissions/enqueue", c.enqueueSubmission).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/submissions/list", c.listSubmissions).Methods("GET")
	c.router.HandleFunc("/submissions/view/{subm_id}", c.getSubmission).Methods("GET")
	c.router.HandleFunc("/submissions/subscribe", c.subscribeToResults).Methods("GET")

	// execution
	c.router.HandleFunc("/execution/enqueue", c.enqueueExecution).Methods("POST", "OPTIONS")
	c.router.HandleFunc("/execution/communicate/{exec_id}", c.communicateWithExec)

	// languages
	c.router.HandleFunc("/languages/list", c.listLanguages).Methods("GET", "OPTIONS")

	c.router.Use(mux.CORSMethodMiddleware(c.router))
}
