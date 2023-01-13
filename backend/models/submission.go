package models

import "time"

type TaskSubmission struct {
	ID        uint      `json:"submission_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UserId    string    `json:"user_id"`
	TaskCode  string    `json:"task_code"`
	LangCode  string    `json:"lang_code"`
}

type ExecSubmission struct {
	ID        uint      `json:"submission_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UserId    string    `json:"user_id"`
	UserCode  string    `json:"user_code"`
	LangCode  string    `json:"lang_code"`
	StdInput  string    `json:"std_input"`
}
