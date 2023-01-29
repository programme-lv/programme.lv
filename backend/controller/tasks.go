package controller

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/utils"
	"github.com/gorilla/mux"
	"log"
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

		// PARSE TESTS
		testDir, err := os.ReadDir(filepath.Join(decompPath, "tests"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var tests []models.TaskTest
		var testNames = make(map[string]bool)
		// read all test names into testNames
		for _, test := range testDir {
			// remove extension from test.Name()
			testName := test.Name()[:len(test.Name())-len(filepath.Ext(test.Name()))]
			testNames[testName] = true
		}
		// loop through test names
		for testName := range testNames {
			inPath := filepath.Join(decompPath, "tests", testName+".in")
			ansPath := filepath.Join(decompPath, "tests", testName+".ans")
			inBytes, _ := os.ReadFile(inPath)
			ansBytes, _ := os.ReadFile(ansPath)
			tests = append(tests, models.TaskTest{
				Name:   testName,
				Input:  string(inBytes),
				Answer: string(ansBytes),
			})
		}

		// PARSE MARKDOWN STATEMENTS
		var mdStatements []models.MarkdownStatement
		mdDir := filepath.Join(decompPath, "statements")
		mdDirEntries, err := os.ReadDir(mdDir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, mdDirEntry := range mdDirEntries {
			if !mdDirEntry.IsDir() {
				continue
			}
			mdDirEntryPath := filepath.Join(mdDir, mdDirEntry.Name())
			mdDirEntryEntries, err := os.ReadDir(mdDirEntryPath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var mdStatement models.MarkdownStatement
			mdStatement.Name = mdDirEntry.Name()
			for _, mdDirEntryEntry := range mdDirEntryEntries {
				if mdDirEntryEntry.IsDir() {
					var examples = make([]models.MDStatementExample, 0)
					examplesDirPath := filepath.Join(mdDirEntryPath, mdDirEntryEntry.Name())
					examplesDirEntries, err := os.ReadDir(examplesDirPath)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					exampleNames := make(map[string]bool)
					for _, examplesDirEntry := range examplesDirEntries {
						exampleName := examplesDirEntry.Name()[:len(examplesDirEntry.Name())-len(filepath.Ext(examplesDirEntry.Name()))]
						exampleNames[exampleName] = true
					}
					for exampleName := range exampleNames {
						exampleInPath := filepath.Join(examplesDirPath, exampleName+".in")
						exampleOutPath := filepath.Join(examplesDirPath, exampleName+".out")
						exampleInBytes, err := os.ReadFile(exampleInPath)
						if err != nil {
							http.Error(w, err.Error(), http.StatusBadRequest)
							return
						}
						exampleOutBytes, err := os.ReadFile(exampleOutPath)
						if err != nil {
							http.Error(w, err.Error(), http.StatusBadRequest)
							return
						}
						examples = append(examples, models.MDStatementExample{
							Input:  string(exampleInBytes),
							Output: string(exampleOutBytes),
						})
					}
					mdStatement.Examples = examples
					continue
				}
				entryContent, err := os.ReadFile(filepath.Join(mdDirEntryPath, mdDirEntryEntry.Name()))
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				switch mdDirEntryEntry.Name() {
				case "description.md":
					mdStatement.Desc = string(entryContent)
				case "input.md":
					mdStatement.Input = string(entryContent)
				case "output.md":
					mdStatement.Output = string(entryContent)
				case "scoring.md":
					mdStatement.Scoring = string(entryContent)
				case "notes.md":
					mdStatement.Notes = string(entryContent)
				}
			}
			mdStatements = append(mdStatements, mdStatement)
		}

		task := models.Task{
			ID: taskTOML.Code,

			Name:   taskTOML.Name,
			Author: taskTOML.Author,
			Type:   taskTOML.Type,

			TimeLim: uint32(math.Round(taskTOML.TimeLim * 1000)),
			MemLim:  taskTOML.MemLim,

			Tests: tests,

			MDStatements: mdStatements,
		}

		tx := c.database.Begin()
		var tags []models.TaskTag
		for _, tagName := range taskTOML.Tags {
			var tag models.TaskTag
			tag.Name = tagName
			err = tx.FirstOrCreate(&tag, tag).Error
			log.Println(tag, tagName)
			if err != nil {
				tx.Rollback()
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			tags = append(tags, tag)
		}
		task.Tags = tags
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

	tx := c.database.Begin()

	// DELETE MARKDOWN STATEMENT EXAMPLES
	err = tx.Exec("DELETE FROM md_statement_examples WHERE markdown_statement_id in (SELECT id FROM markdown_statements WHERE task_id = ?)", task.ID).Error
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// DELETE MARKDOWN STATEMENTS
	err = tx.Where("task_id = ?", task.ID).Delete(&models.MarkdownStatement{}).Error
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// DELETE TESTS
	err = tx.Where("task_id = ?", task.ID).Delete(&models.TaskTest{}).Error
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// CLEAR TAGS
	err = tx.Model(&task).Association("Tags").Clear()
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// DELETE TASK ITSELF
	err = tx.Delete(&task).Error
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx.Commit()
}
