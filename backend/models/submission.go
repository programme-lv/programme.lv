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

	LanguageId string   `json:"lang_id" gorm:"not null"`
	Language   Language `json:"language" gorm:"foreignKey:LanguageId"`

	SrcCode string `json:"src_code" gorm:"not null"`

	TaskSubmJobs []TaskSubmJob `json:"task_subm_jobs"`
}

type TaskSubmJob struct {
	ID        uint64    `json:"subm_job_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	TaskSubmissionId uint64         `json:"subm_id" gorm:"not null"`
	TaskSubmission   TaskSubmission `json:"task_submission" gorm:"foreignKey:TaskSubmissionId"`

	TaskVersion int `json:"task_version" gorm:"not null"`

	Status string `json:"status" gorm:"not null"`
	Score  int    `json:"score"`
}

type TaskSubmJobTest struct {
	ID        uint64    `json:"subm_job_test_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	TaskSubmJobId uint64      `json:"subm_job_id" gorm:"not null"`
	TaskSubmJob   TaskSubmJob `json:"task_subm_job" gorm:"foreignKey:TaskSubmJobId"`

	TestId uint64 `json:"test_id" gorm:"not null"`
	Test   Test   `json:"test" gorm:"foreignKey:TestId"`

	Output string `json:"output"`
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
