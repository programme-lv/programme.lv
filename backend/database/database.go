package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connStr string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connStr), &gorm.Config{})
}
