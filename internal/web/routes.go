package web

import (
	"log/slog"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/service"
	"github.com/justinas/alice"
	"github.com/stephenafamo/bob"
)

type webServer struct {
	logger        *slog.Logger
	db            bob.Executor
	service       *service.Service
	staticHandler http.Handler
}

func NewWebServer(logger *slog.Logger, db bob.Executor, tournamentService *service.Service) http.Handler {
	return &webServer{
		logger:        logger,
		db:            db,
		service:       tournamentService,
		staticHandler: http.StripPrefix("/static", newStaticHandler()),
	}
}

func (server *webServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Router
	router := http.NewServeMux()

	// Middleware Chains
	htmlPage := alice.New(HTMLContentMiddleware)
	htmxFragment := htmlPage

	// Static
	router.Handle("/static/", server.staticHandler)

	// HTML Pages
	router.Handle("/", htmlPage.ThenFunc(server.tournamentsHandlerFunc))
	router.Handle("GET /players", htmlPage.ThenFunc(server.playersHandlerFunc))
	router.Handle("GET /tournaments", htmlPage.ThenFunc(server.tournamentsHandlerFunc))
	router.Handle("GET /tournaments/"+tournamentID.DynamicSegment(), htmlPage.ThenFunc(server.tournamentHandlerFunc))

	// HTMX Fragments
	router.Handle(createPlayerModalResource.GetEndpoint(), htmxFragment.ThenFunc(server.createPlayerModalHandlerFunc))
	router.Handle(createPlayerFormResource.PostEndpoint(), htmxFragment.ThenFunc(server.createPlayerFormHandlerFunc))
	router.Handle(deletePlayerModalResource.GetEndpointWithPathValues(playerID), htmxFragment.ThenFunc(server.deletePlayerModalHandlerFunc))
	router.Handle(createTournamentModalResource.GetEndpoint(), htmxFragment.ThenFunc(server.createTournamentModalHandlerFunc))
	router.Handle(createTournamentFormResource.PostEndpoint(), htmxFragment.ThenFunc(server.createTournamentFormHandlerFunc))
	router.Handle(registerPlayerModalResource.GetEndpointWithPathValues(playerID), htmxFragment.ThenFunc(server.registerPlayerModalHandlerFunc))
	router.Handle(registerPlayerButtonResource.PostEndpoint(), htmxFragment.ThenFunc(server.registerPlayerButtonHandlerFunc))
	router.Handle(deregisterPlayerButtonResource.PostEndpoint(), htmxFragment.ThenFunc(server.deregisterPlayerButtonHandlerFunc))

	router.HandleFunc("DELETE /players/"+playerID.DynamicSegment(), server.deletePlayerHandlerFunc)

	router.ServeHTTP(w, r)
}
