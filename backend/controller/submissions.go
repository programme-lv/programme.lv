package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/models"
)

func (c *Controller) enqueueSubmission(w http.ResponseWriter, r *http.Request) {
	var submission models.TaskSubmission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	c.database.Create(&submission)
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
	} // echo back the submission
	c.scheduler.EnqueueSubmission(submission)
}

func (c *Controller) listSubmissions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *Controller) getSubmission(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *Controller) subscribeToResults(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
