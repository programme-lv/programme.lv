package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/models"
)

// c.router.HandleFunc("/tasks/list", c.listTasks).Methods("GET", "OPTIONS")
func (c *Controller) listTasks(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	tasks, err := c.taskFS.GetTaskList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send the response
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// c.router.HandleFunc("/tasks/view/{task_code}", c.getTask).Methods("GET", "OPTIONS")
func (c *Controller) getTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	taskCode := mux.Vars(r)["task_code"]
	task, err := c.taskFS.GetTaskWithStatements(taskCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send the response
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// c.router.HandleFunc("/tasks/statement/{task_code}/{filename}", c.getPDFStatement).Methods("GET", "OPTIONS")
func (c *Controller) getPDFStatement(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	w.Header().Set("Content-Type", "application/pdf")

	taskCode := mux.Vars(r)["task_code"]
	filename := mux.Vars(r)["filename"]
	statement, err := c.taskFS.GetTaskPDFStatementBytes(taskCode, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = w.Write(statement)
}

func (c *Controller) createTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	err := r.ParseMultipartForm(50 * (1 << 20)) // ~ 50 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mForm := r.MultipartForm
	for k := range mForm.File {
		file, _, err := r.FormFile(k)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = c.taskFS.CreateTaskVersion(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(200)
}

func (c *Controller) deleteTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	log.Println(r.Body)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	// decode the request
	var task models.Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&task)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.database.Delete(&task).Error
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("deleted code: ", task.Code)
	log.Println("deleted name: ", task.Name)

	// echo back the task
	resp, err := json.Marshal(task)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// send the response
	_, err = w.Write(resp)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
