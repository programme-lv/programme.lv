package controller

import (
	"encoding/json"
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/gorilla/mux"
	"net/http"
)

func (c *Controller) enqueueSubmission(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var taskSubmReq struct {
		TaskCode   string `json:"task_code" validate:"required"`
		SrcCode    string `json:"src_code" validate:"required"`
		LanguageId string `json:"lang_id" validate:"required"`
	}

	err := json.NewDecoder(r.Body).Decode(&taskSubmReq)
	if err != nil {
		c.errorLogger.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = c.validate.Struct(taskSubmReq)
	if err != nil {
		c.errorLogger.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	submission := models.TaskSubmission{
		UserId:     1,
		TaskCode:   taskSubmReq.TaskCode,
		SrcCode:    taskSubmReq.SrcCode,
		LanguageId: taskSubmReq.LanguageId,
	}

	err = c.database.Create(&submission).Error
	if err != nil {
		c.errorLogger.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	c.infoLogger.Printf("HTTP created submission %v for task %v", submission.ID, submission.TaskCode)

	err = c.scheduler.EnqueueSubmission(&submission)
	if err != nil {
		c.errorLogger.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(submission)
	if err != nil {
		c.errorLogger.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		c.errorLogger.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

func (c *Controller) listSubmissions(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var submissions []models.TaskSubmission
	err := c.database.Model(&models.TaskSubmission{}).Preload("TaskSubmJobs").Order("created_at desc").Find(&submissions).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type submWithStatus struct {
		models.TaskSubmission
		Status string `json:"status"`
	}

	var submsWithStatus []submWithStatus
	for _, subm := range submissions {
		status := ""
		for _, job := range subm.TaskSubmJobs {
			status = job.Status
		}
		submWithStatus := submWithStatus{
			TaskSubmission: subm,
			Status:         status,
		}
		submWithStatus.TaskSubmJobs = nil // don't send data that isn't required
		submsWithStatus = append(submsWithStatus, submWithStatus)
	}

	resp, err := json.Marshal(submsWithStatus)
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
