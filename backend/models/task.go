package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Code string `json:"task_code" gorm:"unique;not null"`
	Name string `json:"task_name" gorm:"unique;not null"`
}
