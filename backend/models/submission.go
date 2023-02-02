package models

import (
	"time"
)

type TaskSubmission struct {
	ID        uint64    `json:"subm_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	UserId uint64 `json:"user_id" gorm:"not null"`
	User   *User  `json:"user,omitempty" gorm:"foreignKey:UserId"`

	TaskCode string `json:"task_code" gorm:"not null"`
	Task     *Task  `json:"task,omitempty" gorm:"foreignKey:TaskCode"`

	LanguageId string    `json:"lang_id" gorm:"not null"`
	Language   *Language `json:"language,omitempty" gorm:"foreignKey:LanguageId"`

	SrcCode string `json:"src_code" gorm:"not null"`

	TaskSubmEvals []*TaskSubmEvaluation `json:"task_subm_evals"`
}

type TaskSubmEvaluation struct {
	ID        uint64    `json:"subm_eval_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	TaskSubmissionId uint64          `json:"subm_id" gorm:"not null"`
	TaskSubmission   *TaskSubmission `json:"task_submission,omitempty" gorm:"foreignKey:TaskSubmissionId"`

	CompilationStdout string `json:"compilation_stdout"`
	CompilationStderr string `json:"compilation_stderr"`

	MaximumTime   uint64 `json:"maximum_time"`
	MaximumMemory uint64 `json:"maximum_memory"`
	TotalTime     uint64 `json:"total_time"`
	TotalMemory   uint64 `json:"total_memory"`

	Status string `json:"status" gorm:"not null"`
	Score  int    `json:"score"`

	TaskSubmEvalTests []*TaskSubmEvalTest `json:"task_subm_job_tests"`
}

type TaskSubmEvalTest struct {
	ID        uint64    `json:"subm_eval_test_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	TaskSubmEvaluationId uint64              `json:"subm_eval_id" gorm:"not null"`
	TaskSubmEvaluation   *TaskSubmEvaluation `json:"task_subm_eval,omitempty" gorm:"foreignKey:TaskSubmEvaluationId"`

	TestId uint64    `json:"test_id" gorm:"not null"`
	Test   *TaskTest `json:"test,omitempty" gorm:"foreignKey:TestId"`

	Time   uint64 `json:"time"`
	Memory uint64 `json:"memory"`

	Stdout string `json:"stdout,omitempty"`
	Stderr string `json:"stderr,omitempty"`

	Status string `json:"status" gorm:"not null"`
	Score  int    `json:"score"`
}

type ExecSubmission struct {
	ID        uint64    `json:"exec_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	UserId uint64 `json:"user_id"`
	User   *User  `json:"user"`

	LanguageId string    `json:"lang_id"`
	Language   *Language `json:"language"`

	SrcCode  string `json:"src_code"`
	StdInput string `json:"std_input"`
}
