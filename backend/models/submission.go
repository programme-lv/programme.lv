package models

import "time"

type TaskSubmBase struct {
	TaskCode    string `json:"task_code"`
	LangCode    string `json:"lang_code"`
	SubmSrcCode string `json:"subm_src_code"`
}

// TaskSubmission is managed by GORM
type TaskSubmission struct {
	ID        uint      `json:"submission_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UserId    int32     `json:"user_id"`
	TaskSubmBase
}

type ExecSubmission struct {
	ID          uint      `json:"submission_id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_time"`
	UserId      int32     `json:"user_id"`
	SubmSrcCode string    `json:"subm_src_code"`
	LangCode    string    `json:"lang_code"`
	StdInput    string    `json:"std_input"`
}
