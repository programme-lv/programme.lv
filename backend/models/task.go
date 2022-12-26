package models

import "time"

type Task struct {
	Code      string `json:"task_code" gorm:"primarykey"`
	Name      string `json:"task_name" gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskStatement struct {
	TaskCode string `json:"task_code"`
	Desc     string `json:"statement_desc"`
	Input    string `json:"statement_input"`
	Output   string `json:"statement_output"`
	Notes    string `json:"statement_notes"`
	Scoring  string `json:"statement_scoring"`
}
