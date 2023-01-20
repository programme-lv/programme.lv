package models

import "time"

// Language is managed by GORM
type Language struct {
	ID        string    `json:"lang_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Name string `json:"name" gorm:"unique;not null"`
}
