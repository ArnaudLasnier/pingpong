package startup

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"

	tournamentdatabase "github.com/ArnaudLasnier/pingpong/internal/tournamentdatabase"
	tournamentservice "github.com/ArnaudLasnier/pingpong/internal/tournamentservice"
	tournamentweb "github.com/ArnaudLasnier/pingpong/internal/tournamentweb"
	"github.com/aarondl/opt/null"
	"github.com/caarlos0/env/v11"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jonboulle/clockwork"
	"github.com/stephenafamo/bob"
)

type DatabaseSchema string

const (
	TournamentDatabaseSchema DatabaseSchema = "tournament"
)

type application struct {
	logger            *slog.Logger
	config            null.Val[Configuration]
	httpServer        *http.Server
	pgxPool           *pgxpool.Pool
	sqlDB             *sql.DB
	db                bob.DB
	migrationRunners  []*migrate.Migrate
	clock             clockwork.Clock
	tournamentService *tournamentservice.Service
}

func NewApplication() *application {
	return &application{}
}

func NewApplicationWithConfiguration(config Configuration) *application {
	return &application{
		config: null.From(config),
	}
}

func (app *application) Start() {
	app.mustSetupLogger()
	app.mustLoadConfigurationIfNotSet()
	app.mustSetupDatabaseConnectionPool()
	app.mustSetupDatabaseMigrations()
	app.mustRunDatabaseMigrations()
	app.mustSetupServices()
	app.logger.Info("application started successfully")
	app.logger.Info("starting server", slog.Int("port", app.config.MustGet().Port))
	app.mustStartServer()
}

type Configuration struct {
	Port        int    `env:"PORT" envDefault:"3000"`
	DatabaseURI string `env:"DATABASE_URI"`
}

func (app *application) mustSetupLogger() {
	app.logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(app.logger)
	slog.SetLogLoggerLevel(slog.LevelError)
}

func (app *application) mustLoadConfigurationIfNotSet() {
	if app.config.IsSet() {
		return
	}
	config := Configuration{}
	err := env.Parse(&config)
	if err != nil {
		panic(err)
	}
	app.config = null.From(config)
}

func (app *application) mustSetupDatabaseConnectionPool() {
	var err error
	ctx := context.Background()
	databaseURI := app.config.MustGet().DatabaseURI
	parsedDatabaseURI, err := url.Parse(databaseURI)
	if err != nil {
		panic(err)
	}
	connString := parsedDatabaseURI.String()
	app.pgxPool, err = pgxpool.New(ctx, connString)
	if err != nil {
		panic(err)
	}
	err = app.pgxPool.Ping(ctx)
	if err != nil {
		panic(err)
	}
	app.sqlDB = stdlib.OpenDBFromPool(app.pgxPool)
	app.db = bob.NewDB(app.sqlDB)
}

func (app *application) mustSetupDatabaseMigrations() {
	_, err := app.pgxPool.Exec(context.TODO(), "CREATE SCHEMA IF NOT EXISTS migrations")
	if err != nil {
		panic(err)
	}
	databaseName := app.pgxPool.Config().ConnConfig.Database
	app.migrationRunners = append(
		app.migrationRunners,
		tournamentdatabase.NewMigrate(app.sqlDB, databaseName, "tournament"),
	)
}

func (app *application) mustRunDatabaseMigrations() {
	for _, mig := range app.migrationRunners {
		err := mig.Up()
		if errors.Is(err, migrate.ErrNoChange) {
			return
		}
		if err != nil {
			panic(err)
		}
	}
}

func (app *application) mustSetupServices() {
	app.clock = clockwork.NewFakeClock()
	app.tournamentService = tournamentservice.NewService(app.db, app.clock)
}

func (app *application) mustStartServer() {
	address := ":" + strconv.Itoa(app.config.MustGet().Port)
	app.httpServer = &http.Server{
		Addr:    address,
		Handler: tournamentweb.NewHandler(app.logger, app.db, app.tournamentService),
	}
	err := app.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
