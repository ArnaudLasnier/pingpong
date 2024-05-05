package database

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations
var migrationsFS embed.FS

const migrationsDirName = "migrations"

func NewMigrate(db *sql.DB, databaseName string) (*migrate.Migrate, error) {
	var err error
	sourceDriver, err := iofs.New(migrationsFS, migrationsDirName)
	if err != nil {
		return nil, err
	}
	databaseDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	mig, err := migrate.NewWithInstance("iofs", sourceDriver, databaseName, databaseDriver)
	if err != nil {
		return nil, err
	}
	return mig, nil
}
