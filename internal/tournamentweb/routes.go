package tournamentweb

import (
	"log/slog"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/tournamentservice"
	"github.com/justinas/alice"
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
	router.Handle("/", http.RedirectHandler("/players", http.StatusMovedPermanently))
	router.Handle("/static/", h.staticHandler)
	router.Handle("/players", alice.New(MiddlewareHTML).ThenFunc(h.handleGetHomePage))
	router.Handle("/tournaments", alice.New(MiddlewareHTML).ThenFunc(h.handleGetTournamentsPage))
	router.Handle("GET /add-player-modal", alice.New(MiddlewareHTML).ThenFunc(h.handleGetCreatePlayerModal))
	router.Handle("POST /add-player-modal/form", alice.New(MiddlewareHTML).ThenFunc(h.handlePostPlayerCreationForm))
	router.ServeHTTP(w, r)
}
