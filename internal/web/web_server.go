package web

import (
	"log/slog"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/service"
	"github.com/stephenafamo/bob"
)

type webServer struct {
	logger        *slog.Logger
	db            bob.Executor
	service       *service.Service
	staticHandler http.Handler
}

func NewWebServer(logger *slog.Logger, db bob.Executor, tournamentService *service.Service) *webServer {
	return &webServer{
		logger:        logger,
		db:            db,
		service:       tournamentService,
		staticHandler: http.StripPrefix("/static", newStaticHandler()),
	}
}
