package devdb

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	databaseName     = "pingpongdb"
	databaseUser     = "postgres"
	databasePassword = "password"
)

const envFilePath string = "./.envrc"

func Run() {
	envFile := mustCreateEnvFile(envFilePath)
	envConfig := setupPostgresAndReturnConfig()
	envConfig.Dump(envFile)
	waitForSigInt() // this is a blocking call!
}

func mustCreateEnvFile(envFilePath string) *os.File {
	envFile, err := os.Create(envFilePath)
	if err != nil {
		panic(err)
	}
	return envFile
}

func waitForSigInt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

// Map that holds environment variables names as keys and their values as values.
type environmentConfiguration map[string]string

func (envConfig environmentConfiguration) Dump(w io.Writer) {
	for envVariableKey, envVariableValue := range envConfig {
		w.Write([]byte(fmt.Sprintf("export %s=%s\n", envVariableKey, envVariableValue)))
	}
}

func setupPostgresAndReturnConfig() environmentConfiguration {
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
	databaseURI, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}
	return environmentConfiguration{
		"DATABASE_URI": databaseURI,
	}
}
