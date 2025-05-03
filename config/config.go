package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() (*gorm.DB, error) {
	var err error
	db, err = InitializeDb()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
