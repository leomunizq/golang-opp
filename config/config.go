package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	// db = gorm.Open("sqlite3", "test.db")
	// db.AutoMigrate(&schemas.Opening{})
	// return nil
	return errors.New("database connection failed")
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}