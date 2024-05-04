//go:build dev

package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/ArnaudLasnier/pingpong/internal/startup"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const port = 3000

const (
	databaseName     = "pingpongdb"
	databaseUser     = "postgres"
	databasePassword = "password"
)

func main() {
	ctx := context.Background()
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16"),
		postgres.WithDatabase(databaseName),
		postgres.WithUsername(databaseUser),
		postgres.WithPassword(databasePassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
	databaseURI, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}
	config := startup.Configuration{
		Port:        port,
		DatabaseURI: databaseURI,
	}
	slog.Info("development configuration set up", "port", port, "databaseURI", databaseURI)
	app := startup.NewApplicationWithConfiguration(config)
	app.Start()
}
