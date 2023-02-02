package models

import (
	"time"
)

type TaskSubmission struct {
	ID        uint64    `json:"subm_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	UserId uint64 `json:"user_id" gorm:"not null"`
	User   User   `json:"user" gorm:"foreignKey:UserId"`

	TaskCode string `json:"task_code" gorm:"not null"`
	Task     Task   `json:"task" gorm:"foreignKey:TaskCode"`

	LanguageId string   `json:"lang_id" gorm:"not null"`
	Language   Language `json:"language" gorm:"foreignKey:LanguageId"`

	SrcCode string `json:"src_code" gorm:"not null"`

	TaskSubmJobs []TaskSubmEvaluation `json:"task_subm_jobs"`
}

type TaskSubmEvaluation struct {
	ID        uint64    `json:"subm_job_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	TaskSubmissionId uint64         `json:"subm_id" gorm:"not null"`
	TaskSubmission   TaskSubmission `json:"task_submission" gorm:"foreignKey:TaskSubmissionId"`

	Status string `json:"status" gorm:"not null"`
	Score  int    `json:"score"`
}

type TaskSubmEvalTest struct {
	ID        uint64    `json:"subm_job_test_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	TaskSubmJobId uint64             `json:"subm_job_id" gorm:"not null"`
	TaskSubmJob   TaskSubmEvaluation `json:"task_subm_job" gorm:"foreignKey:TaskSubmJobId"`

	TestId uint64   `json:"test_id" gorm:"not null"`
	Test   TaskTest `json:"test" gorm:"foreignKey:TestId"`

	Time   uint64 `json:"time"`
	Memory uint64 `json:"memory"`

	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`

	Status string `json:"status" gorm:"not null"`
	Score  int    `json:"score"`
}

type ExecSubmission struct {
	ID        uint64    `json:"exec_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	UserId uint64 `json:"user_id"`
	User   User   `json:"user"`

	LanguageId string   `json:"lang_id"`
	Language   Language `json:"language"`

	SrcCode  string `json:"src_code"`
	StdInput string `json:"std_input"`
}
