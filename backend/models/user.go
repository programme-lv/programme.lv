package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	Username  string `json:"username" gorm:"unique;not null"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`

	Admin bool `json:"admin" gorm:"default:false"`
}
