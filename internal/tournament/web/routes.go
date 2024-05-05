package web

import (
	"net/http"

	tournamentService "github.com/ArnaudLasnier/pingpong/internal/tournament/service"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob"
)

type handler struct {
	db                bob.Executor
	tournamentService *tournamentService.Service
}

func NewPingPongHandler(db bob.Executor, tournamentService *tournamentService.Service) http.Handler {
	return &handler{
		db:                db,
		tournamentService: tournamentService,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		err := HelloWorld().Render(w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	router.ServeHTTP(w, r)
}

func HelloWorld() g.Node {
	return h.P(g.Text("Hello world!"))
}
