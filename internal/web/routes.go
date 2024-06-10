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
	htmlPage := alice.New(HTMLContentMiddleware)
	htmxFragment := htmlPage

	// Static
	router.Handle("/static/", handler.staticHandler)

	// Home Page
	router.Handle("/", htmlPage.ThenFunc(handler.tournaments))

	// Get Players
	router.Handle("GET /players", htmlPage.ThenFunc(handler.players))

	// Create Player
	router.Handle("GET /create-player-modal", htmxFragment.ThenFunc(handler.createPlayerModalHandler))
	router.Handle("POST /create-player-modal/form", htmxFragment.ThenFunc(handler.createPlayerFormHandler))

	// Get Tournaments
	router.Handle("GET /tournaments", htmlPage.ThenFunc(handler.tournaments))
	router.Handle("GET /tournaments/"+tournamentID.Segment(), htmlPage.ThenFunc(handler.tournament))

	// Create Tournament
	router.Handle("GET /create-tournament-modal", htmxFragment.ThenFunc(handler.createTournamentModalHandler))
	router.Handle("POST /create-tournament-modal/form", htmxFragment.ThenFunc(handler.createTournamentFormHandler))

	// Add Participant
	router.Handle("GET /add-participant-modal", htmxFragment.ThenFunc(handler.addParticipantModalHandler))
	router.Handle("POST /add-participant-modal/form", htmxFragment.ThenFunc(handler.addParticipantFormHandler))
	router.Handle("POST /add-participant-to-tournament/"+tournamentID.Segment(), htmxFragment.ThenFunc(handler.addParticipants))

	router.ServeHTTP(w, r)
}
