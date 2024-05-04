package startup

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	pingpong_database "github.com/ArnaudLasnier/pingpong/internal/pingpong/database"
	pingpong_web "github.com/ArnaudLasnier/pingpong/internal/pingpong/web"
	"github.com/aarondl/opt/null"
	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stephenafamo/bob"
)

type application struct {
	config     null.Val[Configuration]
	httpServer *http.Server
	pgxPool    *pgxpool.Pool
	sqlDB      *sql.DB
	db         bob.DB
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
	mig, err := pingpong_database.NewMigrate(app.sqlDB, databaseName)
	if err != nil {
		panic(err)
	}
	err = mig.Up()
	if err != nil {
		panic(err)
	}
}

func (app *application) mustStartServer() {
	address := ":" + strconv.Itoa(app.config.MustGet().Port)
	app.httpServer = &http.Server{
		Addr:    address,
		Handler: pingpong_web.NewHandler(app.db),
	}
	err := app.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
