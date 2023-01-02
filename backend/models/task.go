package models

import "time"

type Subtask struct {
	Name    string
	Score   int
	Pattern string
}

type Task struct {
	Code      string    `json:"code" gorm:"primarykey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	Version   int       `json:"version" gorm:"not null"`
	Author    string    `json:"author"`
	Tags      []string  `json:"tags"`
	Type      string    `json:"type"`
	TimeLim   float32   `json:"time_lim" toml:"time_lim"`
	MemLim    int       `json:"mem_lim" toml:"mem_lim"`
	Subtasks  []Subtask `json:"subtasks"`
	CreatedAt time.Time `json:"created"`
}

type TaskStatement struct {
	TaskCode string `json:"task_code"`
	Desc     string `json:"desc"`
	Input    string `json:"input"`
	Output   string `json:"output"`
	Notes    string `json:"notes"`
	Scoring  string `json:"scoring"`
}
