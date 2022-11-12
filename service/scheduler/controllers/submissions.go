package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/scheduler/database"
	"github.com/KrisjanisP/deikstra/service/scheduler/models"
)

func EnqueueSubmission(w http.ResponseWriter, r *http.Request) {
	var submission models.TaskSubmission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	database.Instance.Create(&submission)
	resp, err := json.Marshal(submission)
	w.Write(resp)
	return
}

func ListSubmissions(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func GetSubmission(w http.ResponseWriter, r *http.Request) {
	// TODO
}
