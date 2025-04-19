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

	db, err = openDb()
	if err != nil {
		logger.ErrorF("Error opening database: %v", err)
		return nil, err
	}
	m, err := migrate.New(
		"file://infra/database/migrations",
		buildMigrateURL(),
	)
	if err != nil {
		logger.ErrorF("Error creating migrations: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.ErrorF("Cannot run migrations: %v", err)
	}

	return db, nil
}

func openDb() (*gorm.DB, error) {
	var err error

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: defineDNS(),
	}), &gorm.Config{})

	return db, err
}

func defineDNS() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v TimeZone=America/Sao_Paulo",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), "1234",
		os.Getenv("POSTGRES_DB"), "5432")
}

func buildMigrateURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
}
