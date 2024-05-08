package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/urfave/cli/v2"
)

const (
	databaseName     = "pingpongdb"
	databaseUser     = "postgres"
	databasePassword = "password"
)

func main() {
	var err error
	ctx := context.Background()
	cmd := &cli.App{
		Name: "localdeps",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "env-file",
				Required: true,
			},
		},
		Action: func(cliCtx *cli.Context) error {
			var err error
			envFilePath := cliCtx.String("env-file")
			envFilePath, err = filepath.Abs(envFilePath)
			if err != nil {
				return err
			}
			envFile, err := os.Create(envFilePath)
			if err != nil {
				return err
			}
			envConfig := setupDependencies(ctx)
			envConfig.Dump(envFile)
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			<-c
			return nil
		},
	}
	err = cmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

// Map that holds environment variables names as keys and their values as values.
type environmentConfiguration map[string]string

func (envConfig environmentConfiguration) Dump(w io.Writer) {
	for envVariableKey, envVariableValue := range envConfig {
		w.Write([]byte(fmt.Sprintf("export %s=%s\n", envVariableKey, envVariableValue)))
	}
}

func setupDependencies(ctx context.Context) environmentConfiguration {
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
	databaseURI, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}
	return environmentConfiguration{
		"DATABASE_URI": databaseURI,
	}
}
