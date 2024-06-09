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

	"github.com/ArnaudLasnier/pingpong/internal/database"
	"github.com/ArnaudLasnier/pingpong/internal/service"
	"github.com/ArnaudLasnier/pingpong/internal/web"
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
	config            Configuration
	httpServer        *http.Server
	pgxPool           *pgxpool.Pool
	sqlDB             *sql.DB
	db                bob.DB
	migrate           *migrate.Migrate
	clock             clockwork.Clock
	tournamentService *service.Service
}

func NewApplication() *application {
	return &application{}
}

func (app *application) Start() {
	app.mustSetupLogger()
	app.mustLoadConfiguration()
	app.mustSetupDatabaseConnectionPool()
	app.mustSetupDatabaseMigrations()
	app.mustRunDatabaseMigrations()
	app.mustSetupServices()
	app.logger.Info("application started successfully")
	app.logger.Info("starting server", slog.Int("port", app.config.Port))
	app.mustStartServer()
}

type Configuration struct {
	Port        int    `env:"PORT" envDefault:"3000"`
	DatabaseURI string `env:"DATABASE_URI"`
}

func (app *application) mustSetupLogger() {
	app.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(app.logger)
	slog.SetLogLoggerLevel(slog.LevelError)
}

func (app *application) mustLoadConfiguration() {
	err := env.Parse(&app.config)
	if err != nil {
		panic(err)
	}
}

func (app *application) mustSetupDatabaseConnectionPool() {
	var err error
	ctx := context.Background()
	databaseURI := app.config.DatabaseURI
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
	databaseName := app.pgxPool.Config().ConnConfig.Database
	app.migrate = database.NewMigrate(app.sqlDB, databaseName)
}

func (app *application) mustRunDatabaseMigrations() {
	err := app.migrate.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return
	}
	if err != nil {
		panic(err)
	}
}

func (app *application) mustSetupServices() {
	app.clock = clockwork.NewFakeClock()
	app.tournamentService = service.NewService(app.db, app.clock)
}

func (app *application) mustStartServer() {
	address := ":" + strconv.Itoa(app.config.Port)
	app.httpServer = &http.Server{
		Addr:    address,
		Handler: web.NewHandler(app.logger, app.db, app.tournamentService),
	}
	err := app.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
