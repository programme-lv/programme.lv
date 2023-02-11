package utils

import (
	"github.com/BurntSushi/toml"
	"github.com/KrisjanisP/deikstra/service/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func ParseTaskFile(file multipart.File) (task *models.Task, err error) {
	task = &models.Task{}

	var downPath string
	downPath, err = DownloadMultiPartFile(file, "task.zip")
	if err != nil {
		return
	}

	var unzippedPath string
	unzippedPath, err = DecompressZIP(downPath)
	if err != nil {
		return
	}

	var taskToml *TaskToml
	taskToml, err = ParseTaskToml(filepath.Join(unzippedPath, "problem.toml"))
	if err != nil {
		return
	}

	task.ID = taskToml.Code
	task.Name = taskToml.Name
	task.Type = taskToml.Type
	task.Author = taskToml.Author
	task.TimeLim = taskToml.TimeLim
	task.MemLim = taskToml.MemLim

	task.Tags = make([]*models.TaskTag, len(taskToml.Tags))
	for i, tag := range taskToml.Tags {
		task.Tags[i] = &models.TaskTag{ID: tag}
	}

	task.Tests, err = ParseTaskTests(filepath.Join(unzippedPath, "tests"))
	if err != nil {
		return
	}

	task.MDStatements, err = ParseTaskMDStatements(filepath.Join(unzippedPath, "statements"))
	return
}

type SubtaskTOML struct {
	Points  uint32 `json:"points"`
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
	TimeLim   uint32        `json:"time_lim" toml:"time_lim"`
	MemLim    uint32        `json:"mem_lim" toml:"mem_lim"`
	Subtasks  []SubtaskTOML `json:"subtasks"`
	CreatedAt time.Time     `json:"created_time"`
}

func ParseTaskToml(tomlPath string) (taskToml *TaskToml, err error) {
	var tomlBytes []byte
	tomlBytes, err = os.ReadFile(tomlPath)
	if err != nil {
		return
	}
	_, err = toml.Decode(string(tomlBytes), &taskToml)
	return
}

func ParseTaskTests(testsPath string) (tests []*models.TaskTest, err error) {
	testDir, err := os.ReadDir(testsPath)
	var testNames = make(map[string]bool)

	// read all test names into testNames
	for _, test := range testDir {
		filename := test.Name()
		nameNoExt := filename[:len(filename)-len(filepath.Ext(filename))]
		testNames[nameNoExt] = true
	}

	// loop through test names
	for testName := range testNames {
		inPath := filepath.Join(testsPath, testName+".in")
		ansPath := filepath.Join(testsPath, testName+".ans")
		inBytes, _ := os.ReadFile(inPath)
		ansBytes, _ := os.ReadFile(ansPath)
		tests = append(tests, &models.TaskTest{
			Name:   testName,
			Input:  string(inBytes),
			Answer: string(ansBytes),
		})
	}

	return
}

func ParseTaskMDStatements(statementsPath string) (statements []*models.MarkdownStatement, err error) {
	statements = make([]*models.MarkdownStatement, 0)

	mdDirEntries, err := os.ReadDir(statementsPath)
	if err != nil {
		return
	}

	for _, mdDirEntry := range mdDirEntries {
		if !mdDirEntry.IsDir() {
			continue
		}
		var mdStatement *models.MarkdownStatement
		mdStatement, err = ParseMDStatement(filepath.Join(statementsPath, mdDirEntry.Name()))
		if err != nil {
			return
		}

		statements = append(statements, mdStatement)
	}

	return
}

func ParseMDStatement(statementPath string) (statement *models.MarkdownStatement, err error) {
	statement = &models.MarkdownStatement{}
	statement.Examples = make([]*models.MDStatementExample, 0)

	statement.Name = filepath.Base(statementPath)

	var entries []os.DirEntry
	entries, err = os.ReadDir(statementPath)
	if err != nil {
		return
	}

	for _, entry := range entries {

		if entry.IsDir() {
			var example *models.MDStatementExample
			example, err = ParseMDStatementExample(filepath.Join(statementPath, entry.Name()))
			if err != nil {
				return
			}
			statement.Examples = append(statement.Examples, example)
		} else {
			var fileBytes []byte
			fileBytes, err = os.ReadFile(filepath.Join(statementPath, entry.Name()))
			if err != nil {
				return
			}
			fileStr := string(fileBytes)
			switch entry.Name() {
			case "description.md":
				statement.Desc = fileStr
			case "input.md":
				statement.Input = fileStr
			case "output.md":
				statement.Output = fileStr
			case "scoring.md":
				statement.Scoring = fileStr
			case "notes.md":
				statement.Notes = fileStr
			}

		}
	}

	return
}

func ParseMDStatementExample(examplePath string) (example *models.MDStatementExample, err error) {
	example = &models.MDStatementExample{}

	var dirEntries []os.DirEntry
	dirEntries, err = os.ReadDir(examplePath)
	if err != nil {
		return
	}

	exampleNames := make(map[string]bool)
	for _, entry := range dirEntries {
		nameNoExt := entry.Name()[:len(entry.Name())-len(filepath.Ext(entry.Name()))]
		exampleNames[nameNoExt] = true
	}

	for name := range exampleNames {
		exampleInPath := filepath.Join(examplePath, name+".in")
		exampleOutPath := filepath.Join(examplePath, name+".out")

		var exampleInBytes, exampleOutBytes []byte

		exampleInBytes, err = os.ReadFile(exampleInPath)
		if err != nil {
			return
		}

		exampleOutBytes, err = os.ReadFile(exampleOutPath)
		if err != nil {
			return
		}

		example.Input = string(exampleInBytes)
		example.Output = string(exampleOutBytes)
	}

	return
}
