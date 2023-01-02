package database

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/KrisjanisP/deikstra/service/utils"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type TaskManager struct {
	TasksDir string `mapstructure:"tasks_folder"`
}

func CreateTaskManager(tasksDir string) *TaskManager {
	return &TaskManager{TasksDir: tasksDir}
}

type Subtask struct {
	Name    string
	Score   int
	Pattern string
}

type ProblemTOML struct {
	Code     string
	Name     string
	Version  int
	Author   string
	Tags     []string
	Type     string
	TimeLim  float32 `toml:"time_lim"`
	MemLim   int     `toml:"mem_lim"`
	Subtasks []Subtask
}

// CreateTask creates the task, validates it, names it
func (tm *TaskManager) CreateTask(taskFile multipart.File) error {
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

	problem := ProblemTOML{}
	_, err = toml.Decode(string(problemTOMLBytes), &problem)
	if err != nil {
		return err
	}

	taskFileName := problem.Name + "V" + strconv.Itoa(problem.Version)
	taskPath := filepath.Join(tm.TasksDir, taskFileName)

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
