package web

import (
	"log/slog"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/service"
	"github.com/justinas/alice"
	"github.com/stephenafamo/bob"
)

type handler struct {
	logger            *slog.Logger
	db                bob.Executor
	tournamentService *service.Service
	staticHandler     http.Handler
}

func NewHandler(logger *slog.Logger, db bob.Executor, tournamentService *service.Service) http.Handler {
	return &handler{
		logger:            logger,
		db:                db,
		tournamentService: tournamentService,
		staticHandler:     http.StripPrefix("/static", newStaticHandler()),
	}
}

func (handler *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Router
	router := http.NewServeMux()

	// Middleware Chains
	html := alice.New(MiddlewareHTML)
	htmx := html.Append(MiddlewareNoCache)

	// Static
	router.Handle("/static/", handler.staticHandler)

	// Players
	router.Handle("/players", htmx.ThenFunc(handler.players))
	router.Handle("GET /create-player-modal", htmx.ThenFunc(handler.createPlayerModal))
	router.Handle("POST /create-player-modal/form", htmx.ThenFunc(handler.createPlayerModalForm))

	// Tournaments
	router.Handle("/tournaments", htmx.ThenFunc(handler.tournaments))
	router.Handle("GET /create-tournament-modal", htmx.ThenFunc(handler.createTournamentModal))
	router.Handle("POST /create-tournament-modal/form", htmx.ThenFunc(handler.createTournamentModalForm))
	router.Handle("GET /add-participant-modal", htmx.ThenFunc(handler.addParticipantModal))
	router.Handle("POST /add-participant-modal/form", htmx.ThenFunc(handler.addParticipantModalForm))
	router.Handle("POST /add-participant-to-tournament/{tournamentID}", htmx.ThenFunc(handler.addParticipants))

	router.ServeHTTP(w, r)
}
