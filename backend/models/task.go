package models

import "time"

type Task struct {
	ID        string    `json:"task_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name   string `json:"name" gorm:"not null;unique"`
	Author string `json:"author"`
	Tags   []Tag  `json:"tags" gorm:"many2many:task_tags"`

	TimeLim uint32 `json:"time_lim" toml:"time_lim"`
	MemLim  uint32 `json:"mem_lim" toml:"mem_lim"`

	Tests    []Test    `json:"tests"`
	Subtasks []Subtask `json:"subtasks"`

	MDStatements  []MarkdownStatement `json:"md_statements"`
	PDFStatements []PDFStatement      `json:"pdf_statements"`
}

type Test struct {
	ID        uint64    `json:"test_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	TaskID string `json:"task_id" gorm:"not null"`
	Task   Task   `json:"task" gorm:"foreignKey:TaskID"`

	Input  string `json:"input"`
	Answer string `json:"answer"`

	Subtasks []Subtask `json:"subtasks" gorm:"many2many:subtask_tests"`
}

type Subtask struct {
	ID        uint64    `json:"test_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	TaskID string `json:"task_id" gorm:"not null"`
	Task   Task   `json:"task" gorm:"foreignKey:TaskID"`

	Score int `json:"score" go:"not null"`

	Tests []Test `json:"tests"  gorm:"many2many:subtask_tests"`
}

type Tag struct {
	ID        uint64    `json:"tag_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name  string `json:"name" gorm:"unique;not null"`
	Color string `json:"color"`

	Tasks []Task `json:"tasks" gorm:"many2many:task_tags"`
}
