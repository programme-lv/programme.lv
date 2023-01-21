package controller

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/utils"
	"log"
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

		type subtaskTOML struct {
			Name    string
			Score   int
			Pattern string
		}

		type taskTOML struct {
			Code      string        `json:"code"`
			Name      string        `json:"name"`
			Version   int           `json:"version"`
			Author    string        `json:"author"`
			Tags      []string      `json:"tags"`
			Type      string        `json:"type"`
			TimeLim   float32       `json:"time_lim" toml:"time_lim"`
			MemLim    int           `json:"mem_lim" toml:"mem_lim"`
			Subtasks  []subtaskTOML `json:"subtasks"`
			CreatedAt time.Time     `json:"created_time"`
		}
		problem := taskTOML{}
		_, err = toml.Decode(string(problemTOMLBytes), &problem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(problem)
	}

	w.WriteHeader(200)
}
