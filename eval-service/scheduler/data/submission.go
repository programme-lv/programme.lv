package data

import "time"

type TaskSubmission struct {
	ID        uint      `json:"submission_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	TaskName  string    `json:"task_name"`
	UserCode  string    `json:"user_code"`
	LangId    string    `json:"lang_id"`
}

type ExecSubmission struct {
	UserCode string `json:"user_code"`
	LangId   string `json:"lang_id"`
}
