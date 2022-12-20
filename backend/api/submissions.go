package api

import (
	"encoding/json"
	data2 "github.com/KrisjanisP/deikstra/service/data"
	"log"
	"net/http"
)

func (c *Controller) enqueueSubmission(w http.ResponseWriter, r *http.Request) {
	var submission data2.TaskSubmission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	data2.Instance.Create(&submission)
	resp, err := json.Marshal(submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write(resp) // echo back the submission
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
