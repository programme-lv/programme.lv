package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/models"
)

func listTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *Controller) createTask(w http.ResponseWriter, r *http.Request) {
	// CORS
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

	// process the request
	if len(task.Code) == 0 || len(task.Name) == 0 {
		err = fmt.Errorf("neither task_code nor task_name can be empty")
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.database.Create(&task).Error
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("code: ", task.Code)
	log.Println("name: ", task.Name)

	// echo back the task
	resp, err := json.Marshal(task)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
