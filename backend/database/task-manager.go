package database

import (
	"github.com/BurntSushi/toml"
	"github.com/KrisjanisP/deikstra/service/utils"
	"mime/multipart"
	"os"
	"path/filepath"
)

type TaskManager struct {
}

func CreateTaskManager() *TaskManager {
	return &TaskManager{}
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
		return nil
	}

	return nil
}
