package dbgen

import (
	"context"
	"database/sql"
	"time"

	"github.com/ArnaudLasnier/pingpong/internal/database"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lib/pq"
	"github.com/stephenafamo/bob/gen"
	helpers "github.com/stephenafamo/bob/gen/bobgen-helpers"
	"github.com/stephenafamo/bob/gen/bobgen-psql/driver"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dummyDatabaseName     = "dummy_database"
	dummyDatabaseUser     = "dummy_user"
	dummyDatabasePassword = "dummy_password"
)

const outputPath = "models"

func Run() {
	tempDatabaseURI := setupTempPostgres(dummyDatabaseName, dummyDatabaseUser, dummyDatabasePassword)
	db := openDB(tempDatabaseURI)
	migrateDB(db)
	generateModels(tempDatabaseURI, outputPath)
}

func setupTempPostgres(dbName, dbUser, dbPassword string) (dbURI string) {
	ctx := context.TODO()
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}
	databaseURI, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}
	return databaseURI
}

func openDB(databaseURI string) *sql.DB {
	var err error
	db, err := sql.Open("pgx", databaseURI)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func migrateDB(db *sql.DB) {
	var err error
	mig := database.NewMigrate(db, dummyDatabaseName)
	err = mig.Up()
	if err != nil {
		panic(err)
	}
}

func generateModels(databaseURI string, outputPath string) {
	var err error
	ctx := context.TODO()
	driverConfig := driver.Config{
		Dsn:          databaseURI,
		Schemas:      pq.StringArray{database.SchemaName},
		SharedSchema: database.SchemaName,
		Only:         nil,
		Except:       nil,
		Concurrency:  10,
		UUIDPkg:      "google",
		Output:       outputPath,
		Pkgname:      "models",
		NoFactory:    false,
	}
	driver := driver.New(driverConfig)
	config := gen.Config{
		Generator: "the local DBGEN tool",
	}
	outputs := helpers.DefaultOutputs(driverConfig.Output, driverConfig.Pkgname, config.NoFactory, nil)
	state := &gen.State{
		Config:  config,
		Outputs: outputs,
	}
	err = gen.Run(ctx, state, driver)
	if err != nil {
		panic(err)
	}
}
