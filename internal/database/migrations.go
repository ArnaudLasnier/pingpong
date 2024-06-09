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

func NewMigrate(db *sql.DB, databaseName string, schemaName string) *migrate.Migrate {
	var err error
	sourceDriver, err := iofs.New(migrationsFS, migrationsDirName)
	if err != nil {
		panic(err)
	}
	databaseDriver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName:          databaseName,
		SchemaName:            "migrations",
		MigrationsTable:       schemaName,
		MultiStatementEnabled: true,
	})
	if err != nil {
		panic(err)
	}
	mig, err := migrate.NewWithInstance("iofs", sourceDriver, databaseName, databaseDriver)
	if err != nil {
		panic(err)
	}
	return mig
}
