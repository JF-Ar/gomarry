package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {
	var err error
	logger = GetLogger("db")

	dsn := defineDSN()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error opening database: %v", err)
		return nil, err
	}

	migrateURL := buildMigrateURL()
	m, err := migrate.New(
		"file://infra/database/migrations",
		migrateURL,
	)
	if err != nil {
		logger.ErrorF("Error creating migration instance: %v", err)
		return db, err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.ErrorF("Error applying migrations: %v", err)
		return db, err
	}

	return db, nil
}

func defineDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=America/Sao_Paulo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		"5432",
	)
}

func buildMigrateURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		"5432",
		os.Getenv("POSTGRES_DB"),
	)
}
