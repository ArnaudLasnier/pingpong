package web

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ArnaudLasnier/pingpong/internal/service"
	"github.com/ArnaudLasnier/pingpong/internal/tools/dbgen"
	"github.com/jonboulle/clockwork"
	"github.com/stephenafamo/bob"
)

func NewTestWebServer() *webServer {
	logger := mustSetupLogger()
	tempDatabaseURI := dbgen.SetupTempPostgres()
	fmt.Printf("temp connection string: %s\n\n", tempDatabaseURI)
	sqlDB := dbgen.OpenDB(tempDatabaseURI)
	dbgen.MigrateDB(sqlDB)
	db := bob.NewDB(sqlDB)
	clock := clockwork.NewFakeClock()
	service := service.NewService(db, clock)
	return NewWebServer(logger, db, service)
}

func mustSetupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelError)
	return logger
}
