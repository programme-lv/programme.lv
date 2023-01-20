package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Role     string `json:"role" gorm:"not null"`
}
