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

type TaskWithStatements struct {
	Task
	MDStatements  []MarkdownStatement `json:"md_statements"`
	PDFStatements []PDFStatement      `json:"pdf_statements"`
}
