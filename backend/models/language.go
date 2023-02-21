package models

import "time"

// Language is managed by GORM
type Language struct {
	ID        string    `json:"lang_id" gorm:"primaryKey"`
	UpdatedAt time.Time `json:"updated_time"`

	Name       string  `json:"name" gorm:"unique;not null"`
	Filename   string  `json:"filename" gorm:"not null"`
	CompileCmd *string `json:"compile_cmd"`
	ExecuteCmd string  `json:"execute_cmd" gorm:"not null"`
}
