package tournamentweb

import (
	"log/slog"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/tournamentservice"
	"github.com/stephenafamo/bob"
)

type handler struct {
	logger            *slog.Logger
	db                bob.Executor
	tournamentService *tournamentservice.Service
	staticHandler     http.Handler
}

func NewHandler(logger *slog.Logger, db bob.Executor, tournamentService *tournamentservice.Service) http.Handler {
	return &handler{
		logger:            logger,
		db:                db,
		tournamentService: tournamentService,
		staticHandler:     http.StripPrefix("/static", newStaticHandler()),
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/static/", h.staticHandler)
	router.HandleFunc("/", h.handleGetHomePage)
	router.HandleFunc("GET /add-player-modal", h.handleGetCreatePlayerModal)
	router.HandleFunc("POST /add-player-modal", h.handlePostCreatePlayerModal)
	router.ServeHTTP(w, r)
}

func todoPanic(v any) {
	panic(v)
}
