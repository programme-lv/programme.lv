package models

import "time"

type Task struct {
	ID        string    `json:"task_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name   string    `json:"name" gorm:"not null;unique"`
	Author string    `json:"author"`
	Tags   []TaskTag `json:"tags" gorm:"many2many:tasks_tags;"`

	Type     string   `json:"type" gorm:"not null"`
	TaskType TaskType `gorm:"foreignKey:Type"`

	TimeLim uint32 `json:"time_lim" toml:"time_lim"`
	MemLim  uint32 `json:"mem_lim" toml:"mem_lim"`

	Tests    []TaskTest    `json:"tests"`
	Subtasks []TaskSubtask `json:"subtasks"`

	MDStatements  []MarkdownStatement `json:"md_statements"`
	PDFStatements []PDFStatement      `json:"pdf_statements"`
}

type TaskType struct {
	ID          string `json:"type_id" gorm:"primaryKey"`
	Description string `json:"description"`
}

type TaskTest struct {
	ID        uint64    `json:"test_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name string `json:"name" gorm:"not null"`

	TaskID string `json:"task_id" gorm:"not null"`
	Task   Task   `json:"task" gorm:"foreignKey:TaskID"`

	Input  string `json:"input"`
	Answer string `json:"answer"`

	Subtasks []TaskSubtask `json:"subtasks" gorm:"many2many:subtasks_tests"`
}

type TaskSubtask struct {
	ID        uint64    `json:"test_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name    string `json:"name" gorm:"not null"`
	Pattern string `json:"pattern" gorm:"not null"`

	TaskID string `json:"task_id" gorm:"not null"`
	Task   Task   `json:"task" gorm:"foreignKey:TaskID"`

	Score int `json:"score" go:"not null"`

	Tests []TaskTest `json:"tests"  gorm:"many2many:subtasks_tests"`
}

type TaskTag struct {
	ID        uint64    `json:"tag_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name  string `json:"name" gorm:"unique;not null"`
	Color string `json:"color"`

	Tasks []Task `json:"tasks" gorm:"many2many:tasks_tags"`
}
