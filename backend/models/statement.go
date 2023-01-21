package models

import (
	"time"
)

type MarkdownStatement struct {
	ID        uint64    `json:"statement_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name    string `json:"name" gorm:"not null"`
	Desc    string `json:"desc"`
	Input   string `json:"input"`
	Output  string `json:"output"`
	Notes   string `json:"notes"`
	Scoring string `json:"scoring"`

	Examples []MDSTatementExample `json:"examples"`

	TaskID string `json:"task_id" gorm:"not null"`
	Task   Task   `json:"task" gorm:"foreignKey:TaskID"`
}

type MDSTatementExample struct {
	ID        uint64    `json:"example_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Input  string `json:"input"`
	Output string `json:"output"`

	MarkdownStatementID uint64            `json:"statement_id" gorm:"not null"`
	MarkdownStatement   MarkdownStatement `json:"statement" gorm:"foreignKey:MarkdownStatementID"`
}

type PDFStatement struct {
	ID        uint64    `json:"pdf_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name string `json:"filename"`

	TaskID string `json:"task_id" gorm:"not null"`
	Task   Task   `json:"task" gorm:"foreignKey:TaskID"`
}
