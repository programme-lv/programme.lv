package controller

import (
	"encoding/json"
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/utils"
	"github.com/gorilla/mux"
	"net/http"
)

// c.router.HandleFunc("/tasks/list", c.listTasks).Methods("GET", "OPTIONS")
func (c *Controller) listTasks(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var tasks []models.Task
	err := c.database.Model(&models.Task{}).Preload("Tags").Find(&tasks).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(tasks)
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

// c.router.HandleFunc("/tasks/view/{task_ir}", c.getTask).Methods("GET", "OPTIONS")
func (c *Controller) getTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	taskId := mux.Vars(r)["task_id"]

	var task models.Task
	task.ID = taskId
	err := c.database.Model(&task).Preload("MDStatements.Examples").Preload("Tags").Take(&task).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *Controller) importTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	err := r.ParseMultipartForm(50 * (1 << 20)) // ~ 50 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}

	mForm := r.MultipartForm
	for k := range mForm.File {
		file, _, err := r.FormFile(k)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task, err := utils.ParseTaskFile(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tx := c.database.Begin()
		// soft delete previous tests
		err = tx.Where("task_id", task.ID).Delete(&models.TaskTest{}).Error
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// update task values
		err = tx.Unscoped().Updates(task).Error
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// remove soft delete from task
		err = tx.Unscoped().Model(task).Where("id", task.ID).Update("deleted_at", nil).Error
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tx.Commit()
	}

	w.WriteHeader(200)
}

func (c *Controller) deleteTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var req struct {
		TaskId string `json:"task_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.database.Delete(&models.Task{ID: req.TaskId})
}
