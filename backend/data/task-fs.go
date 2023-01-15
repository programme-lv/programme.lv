package data

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/utils"
)

type TaskFS struct {
	TasksDir string `mapstructure:"tasks_folder"`
}

func CreateTaskManager(tasksDir string) *TaskFS {
	return &TaskFS{TasksDir: tasksDir}
}

func (tfs *TaskFS) ReadFilenameTask(filename string) (models.Task, error) {
	res := models.Task{}

	problemTOMLPath := filepath.Join(tfs.TasksDir, filename, "problem.toml")
	problemTOMLBytes, err := os.ReadFile(problemTOMLPath)
	if err != nil {
		return res, err
	}

	_, err = toml.Decode(string(problemTOMLBytes), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (tfs *TaskFS) GetTaskVersionPath(code string, version int) string {
	filename := code + "V" + strconv.Itoa(version)
	return filepath.Join(tfs.TasksDir, filename)
}

func (tfs *TaskFS) GetTaskList() ([]models.Task, error) {
	files, err := os.ReadDir(tfs.TasksDir)
	if err != nil {
		return nil, err
	}
	tasks := make(map[string]models.Task)
	for _, file := range files {
		task, err := tfs.ReadFilenameTask(file.Name())
		if err != nil {
			return nil, err
		}
		val, exists := tasks[task.Code]
		if !exists {
			tasks[task.Code] = task
		} else if task.Version > val.Version {
			tasks[task.Code] = task
		}
	}
	res := make([]models.Task, 0)
	for _, v := range tasks {
		log.Println(v)
		res = append(res, v)
	}
	return res, nil
}

// CreateTaskVersion creates the task, validates it, names it
func (tfs *TaskFS) CreateTaskVersion(taskFile multipart.File) error {
	dirPath := filepath.Join("/tmp", "deikstra")
	_ = os.MkdirAll(dirPath, os.ModePerm)
	tmpDir, _ := os.MkdirTemp(dirPath, "")
	downPath := filepath.Join(tmpDir, "download.zip")

	err := utils.SaveMultiPartFile(taskFile, downPath)
	if err != nil {
		return err
	}

	decompPath := filepath.Join(tmpDir, "decompressed")
	err = utils.DecompressZIP(downPath, decompPath)
	if err != nil {
		return err
	}

	problemTOMLBytes, err := os.ReadFile(filepath.Join(decompPath, "problem.toml"))
	if err != nil {
		return err
	}

	problem := models.Task{}
	_, err = toml.Decode(string(problemTOMLBytes), &problem)
	if err != nil {
		return err
	}
	problem.CreatedAt = time.Now()
	log.Printf("received %v", problem)

	taskFileName := problem.Code + "V" + strconv.Itoa(problem.Version)
	taskPath := filepath.Join(tfs.TasksDir, taskFileName)

	if _, err := os.Stat(taskPath); err == nil {
		return fmt.Errorf("task %v (version %v) already exists", problem.Name, problem.Version)
	}

	cmd := exec.Command("mv", decompPath, taskPath)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
