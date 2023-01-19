package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KrisjanisP/deikstra/service/models"
)

func (c *Controller) enqueueSubmission(w http.ResponseWriter, r *http.Request) {
	var taskSubmReq models.TaskSubmBase
	err := json.NewDecoder(r.Body).Decode(&taskSubmReq)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	submission := models.TaskSubmission{
		CreatedAt:    time.Now(),
		UserId:       1,
		TaskSubmBase: taskSubmReq}
	c.database.Create(&submission)

	// echo back the submission
	resp, err := json.Marshal(submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	c.scheduler.EnqueueSubmission(submission)
}

func (c *Controller) listSubmissions(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var submissions []models.TaskSubmission
	result := c.database.Find(&submissions)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(submissions)
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

func (c *Controller) getSubmission(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *Controller) subscribeToResults(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
