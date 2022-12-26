package api

import (
	"encoding/json"
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
	log.Println(r.Body)

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
	err = c.database.Create(&task).Error
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("id: ", task.ID)
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
