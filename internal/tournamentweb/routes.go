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
	html := alice.New(MiddlewareHTML)
	htmx := html.Append(MiddlewareNoCache)
	router := http.NewServeMux()
	// router.Handle("/", http.RedirectHandler("/players", http.StatusMovedPermanently))
	router.Handle("/static/", h.staticHandler)
	router.Handle("/players", htmx.ThenFunc(h.handleGetPlayersPage))
	router.Handle("GET /add-player-modal", htmx.ThenFunc(h.handleGetCreatePlayerModal))
	router.Handle("POST /add-player-modal/form", htmx.ThenFunc(h.handlePostPlayerCreationForm))
	router.Handle("/tournaments", htmx.ThenFunc(h.handleGetTournamentsPage))
	router.Handle("GET /create-tournament-modal", htmx.ThenFunc(h.handleGetCreateTournamentModal))
	router.Handle("POST /create-tournament-modal/form", htmx.ThenFunc(h.handlePostTournamentCreationForm))
	router.Handle("/tailwind-test", htmx.ThenFunc(handleTailwindTest))
	router.Handle("/tailwind-test/button", htmx.ThenFunc(tailwindTestButtonResult))
	router.ServeHTTP(w, r)
}
