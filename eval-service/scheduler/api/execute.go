package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/scheduler/data"
)

func (c *APIController) enqueueExecution(w http.ResponseWriter, r *http.Request) {
	var submission data.ExecSubmission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write(resp) // echo back the submission
	c.scheduler.EnqueueExecution(submission)
}
