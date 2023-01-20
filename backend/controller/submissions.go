package controller

import (
	"encoding/json"
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (c *Controller) enqueueSubmission(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var taskSubmReq models.TaskSubmBase
	err := json.NewDecoder(r.Body).Decode(&taskSubmReq)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if taskSubmReq.TaskCode == "" {
		log.Printf("HTTP %s", "task_code is required")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if taskSubmReq.LangCode == "" {
		log.Printf("HTTP %s", "lang_code is required")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if taskSubmReq.SubmSrcCode == "" {
		log.Printf("HTTP %s", "subm_src_code is required")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	submission := models.TaskSubmission{
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
	result := c.database.Order("created_at desc").Find(&submissions)
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

func (c *Controller) getSubmission(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	submissionId := mux.Vars(r)["subm_id"]
	var submission models.TaskSubmission
	result := c.database.First(&submission, submissionId)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(submission)
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

func (c *Controller) subscribeToResults(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
