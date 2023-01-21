package controller

import (
	"net/http"
)

// c.router.HandleFunc("/tasks/list", c.listTasks).Methods("GET", "OPTIONS")
func (c *Controller) listTasks(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

// c.router.HandleFunc("/tasks/view/{task_code}", c.getTask).Methods("GET", "OPTIONS")
func (c *Controller) getTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func (c *Controller) createTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}
