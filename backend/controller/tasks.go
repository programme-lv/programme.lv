package controller

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/utils"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"time"
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
	err := c.database.Find(&tasks).Error
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

// c.router.HandleFunc("/tasks/view/{task_code}", c.getTask).Methods("GET", "OPTIONS")
func (c *Controller) getTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func (c *Controller) createTask(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	err := r.ParseMultipartForm(50 * (1 << 20)) // ~ 50 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mForm := r.MultipartForm
	for k := range mForm.File {
		file, _, err := r.FormFile(k)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// CREATE TMP DIR FOR DOWNLOADING FILES
		tmpDir := filepath.Join("/tmp", "programme")
		_ = os.MkdirAll(tmpDir, os.ModePerm)
		tmpDir, _ = os.MkdirTemp(tmpDir, "")
		downPath := filepath.Join(tmpDir, "download.zip")

		// DOWNLOAD FILE
		err = utils.SaveMultiPartFile(file, downPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// UNZIP FILE
		decompPath := filepath.Join(tmpDir, "decompressed")
		err = utils.DecompressZIP(downPath, decompPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// PARSE PROBLEM.TOML FILE
		problemTOMLBytes, err := os.ReadFile(filepath.Join(decompPath, "problem.toml"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		type SubtaskTOML struct {
			Name    string
			Score   int
			Pattern string
		}
		type TaskToml struct {
			Code      string        `json:"code"`
			Name      string        `json:"name"`
			Version   int           `json:"version"`
			Author    string        `json:"author"`
			Tags      []string      `json:"tags"`
			Type      string        `json:"type"`
			TimeLim   float64       `json:"time_lim" toml:"time_lim"`
			MemLim    uint32        `json:"mem_lim" toml:"mem_lim"`
			Subtasks  []SubtaskTOML `json:"subtasks"`
			CreatedAt time.Time     `json:"created_time"`
		}
		taskTOML := TaskToml{}
		_, err = toml.Decode(string(problemTOMLBytes), &taskTOML)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		os.ReadDir(filepath.Join(decompPath, "testing-dst"))

		tx := c.database.Begin()
		var tags []models.Tag
		for _, tag := range taskTOML.Tags {
			tags = append(tags, models.Tag{
				Name: tag,
			})
		}
		task := models.Task{
			ID: taskTOML.Code,

			Name:   taskTOML.Name,
			Author: taskTOML.Author,

			TimeLim: uint32(math.Round(taskTOML.TimeLim * 1000)),
			MemLim:  taskTOML.MemLim,
		}

		err = tx.Create(&task).Error
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

	var task models.Task
	task.ID = req.TaskId
	err = c.database.Delete(&task).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
