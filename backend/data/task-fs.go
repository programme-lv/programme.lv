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

// readTaskByFilename reads task from problem.toml file
func (tfs *TaskFS) readTaskByFilename(filename string) (models.Task, error) {
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

// GetTaskNewestVersion returns the newest version of the task
func (tfs *TaskFS) getTaskNewestVersion(taskCode string) (models.Task, error) {
	res := models.Task{}
	files, err := os.ReadDir(tfs.TasksDir)
	if err != nil {
		return res, err
	}
	for _, file := range files {
		task, err := tfs.readTaskByFilename(file.Name())
		if err != nil {
			return res, err
		}
		if task.Code == taskCode {
			if task.Version > res.Version {
				res = task
			}
		}
	}
	return res, nil
}

// getTaskPath returns the path to the task newest version
func (tfs *TaskFS) getTaskPath(taskCode string) (string, error) {
	task, err := tfs.getTaskNewestVersion(taskCode)
	if err != nil {
		return "", err
	}
	return filepath.Join(tfs.TasksDir, task.Code+"V"+strconv.Itoa(task.Version)), nil
}

// GetTaskMDStatements returns task (newest version) markdown statements
func (tfs *TaskFS) GetTaskMDStatements(taskCode string) ([]models.MarkdownStatement, error) {
	taskPath, err := tfs.getTaskPath(taskCode)
	if err != nil {
		return nil, err
	}
	statementsPath := filepath.Join(taskPath, "statements")
	statementEntries, err := os.ReadDir(statementsPath)
	if err != nil {
		return nil, err
	}
	statements := make([]models.MarkdownStatement, 0)
	for _, statementDirEntry := range statementEntries {
		if !statementDirEntry.IsDir() { // markdown statements are in directories
			continue
		}
		statement := models.MarkdownStatement{}

		statement.Name = statementDirEntry.Name()

		descPath := filepath.Join(statementsPath, statementDirEntry.Name(), "description.md")
		description, err := os.ReadFile(descPath)
		if err == nil {
			statement.Desc = string(description)
		}

		inputPath := filepath.Join(statementsPath, statementDirEntry.Name(), "input.md")
		input, err := os.ReadFile(inputPath)
		if err == nil {
			statement.Input = string(input)
		}

		outputPath := filepath.Join(statementsPath, statementDirEntry.Name(), "output.md")
		output, err := os.ReadFile(outputPath)
		if err == nil {
			statement.Output = string(output)
		}

		scoringPath := filepath.Join(statementsPath, statementDirEntry.Name(), "scoring.md")
		scoring, err := os.ReadFile(scoringPath)
		if err == nil {
			statement.Scoring = string(scoring)
		}

		examplesPath := filepath.Join(statementsPath, statementDirEntry.Name(), "examples")
		examplesEntries, err := os.ReadDir(examplesPath)
		if err == nil {
			statement.Examples = make([]models.MDSTatementExample, 0)
			exampleNames := make(map[string]bool)
			for _, exampleDirEntry := range examplesEntries {
				if exampleDirEntry.IsDir() { // examples are not in directories
					continue
				}
				name := exampleDirEntry.Name()
				name = name[0 : len(name)-len(filepath.Ext(name))]
				exampleNames[name] = true
			}
			for name := range exampleNames {
				example := models.MDSTatementExample{}
				InputPath := filepath.Join(examplesPath, name+".in")
				OutputPath := filepath.Join(examplesPath, name+".out")
				exampleInput, err := os.ReadFile(InputPath)
				if err == nil {
					example.Input = string(exampleInput)
				}
				exampleOutput, err := os.ReadFile(OutputPath)
				if err == nil {
					example.Output = string(exampleOutput)
				}
				statement.Examples = append(statement.Examples, example)
			}
		}
		statements = append(statements, statement)
	}
	return statements, nil

}

// GetTaskPDFStatements returns task (newest version) pdf statements
func (tfs *TaskFS) GetTaskPDFStatements(taskCode string) ([]models.PDFStatement, error) {
	taskPath, err := tfs.getTaskPath(taskCode)
	if err != nil {
		return nil, err
	}
	statementsPath := filepath.Join(taskPath, "statements")
	statementEntries, err := os.ReadDir(statementsPath)
	if err != nil {
		return nil, err
	}
	statements := make([]models.PDFStatement, 0)
	for _, statementDirEntry := range statementEntries {
		if statementDirEntry.IsDir() { // pdf statements are not in directories
			continue
		}
		statement := models.PDFStatement{}
		statement.Name = statementDirEntry.Name()
		statements = append(statements, statement)
	}
	return statements, nil
}

// GetTaskList returns task newest versions
func (tfs *TaskFS) GetTaskList() ([]models.Task, error) {
	files, err := os.ReadDir(tfs.TasksDir)
	if err != nil {
		return nil, err
	}
	tasks := make(map[string]models.Task)
	for _, file := range files {
		task, err := tfs.readTaskByFilename(file.Name())
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

// GetTaskWithStatements returns task (newest version) with statements
func (tfs *TaskFS) GetTaskWithStatements(taskCode string) (models.TaskWithStatements, error) {
	res := models.TaskWithStatements{}
	task, err := tfs.getTaskNewestVersion(taskCode)
	if err != nil {
		return res, err
	}
	res.Task = task
	res.MDStatements, err = tfs.GetTaskMDStatements(taskCode)
	res.PDFStatements, err = tfs.GetTaskPDFStatements(taskCode)
	if err != nil {
		return res, err
	}
	return res, nil
}

// GetTaskPDFStatementBytes returns task (newest version) pdf statement
func (tfs *TaskFS) GetTaskPDFStatementBytes(taskCode string, filename string) ([]byte, error) {
	// TODO: ensure that the filename is bounded to task statements directory
	taskPath, err := tfs.getTaskPath(taskCode)
	if err != nil {
		return nil, err
	}
	statementsPath := filepath.Join(taskPath, "statements")
	statementPath := filepath.Join(statementsPath, filename)
	return os.ReadFile(statementPath)
}
