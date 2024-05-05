package startup

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	tournamentDatabase "github.com/ArnaudLasnier/pingpong/internal/tournament/database"
	tournamentService "github.com/ArnaudLasnier/pingpong/internal/tournament/service"
	tournamentWeb "github.com/ArnaudLasnier/pingpong/internal/tournament/web"
	"github.com/aarondl/opt/null"
	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jonboulle/clockwork"
	"github.com/stephenafamo/bob"
)

type application struct {
	config            null.Val[Configuration]
	httpServer        *http.Server
	pgxPool           *pgxpool.Pool
	sqlDB             *sql.DB
	db                bob.DB
	clock             clockwork.Clock
	tournamentService *tournamentService.Service
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
	app.mustLoadConfigurationIfNotSet()
	app.mustSetupDatabaseConnections()
	app.mustRunDatabaseMigrations()
	app.mustSetupServices()
	app.mustStartServer()
}

type Configuration struct {
	Port        int    `env:"PORT" envDefault:"3000"`
	DatabaseURI string `env:"DATABASE_URI"`
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

func (app *application) mustSetupDatabaseConnections() {
	var err error
	ctx := context.Background()
	app.pgxPool, err = pgxpool.New(ctx, app.config.MustGet().DatabaseURI)
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

func (app *application) mustRunDatabaseMigrations() {
	var err error
	databaseName := app.pgxPool.Config().ConnConfig.Database
	mig, err := tournamentDatabase.NewMigrate(app.sqlDB, databaseName)
	if err != nil {
		panic(err)
	}
	err = mig.Up()
	if err != nil {
		panic(err)
	}
}

func (app *application) mustSetupServices() {
	app.clock = clockwork.NewFakeClock()
	app.tournamentService = tournamentService.NewService(app.db, app.clock)
}

func (app *application) mustStartServer() {
	address := ":" + strconv.Itoa(app.config.MustGet().Port)
	app.httpServer = &http.Server{
		Addr:    address,
		Handler: tournamentWeb.NewPingPongHandler(app.db, app.tournamentService),
	}
	err := app.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
