package config

import (
	"github.com/JF-Ar/gomarry/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {
	var err error
	logger = GetLogger("db")

	db, err = gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error opening database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.PingPong{})
	if err != nil {
		logger.ErrorF("Error migrating database: %v", err)
		return nil, err
	}

	return db, nil
}
