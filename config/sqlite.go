package config

import (
	"os"

	"github.com/leomunizq/golang-opp/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating a new one")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errf("error creating db directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errf("error creating db file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("sql opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errf("auto migration error: %v", err)
		return nil, err
	}

	logger.Info("Database connection successful")
	return db, nil

}
