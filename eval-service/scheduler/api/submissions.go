package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/scheduler/data"
)

func (c *APIController) enqueueSubmission(w http.ResponseWriter, r *http.Request) {
	var submission data.TaskSubmission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	data.Instance.Create(&submission)
	resp, err := json.Marshal(submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write(resp)
	c.scheduler.TaskQueue <- submission
}

func (c *APIController) listSubmissions(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *APIController) getSubmission(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (c *APIController) subscribeToResults(w http.ResponseWriter, r *http.Request) {
	// TODO
}
