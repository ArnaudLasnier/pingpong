package web

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/google/uuid"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func (handler *webServer) tournamentHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	url := *r.URL
	tournamentID, err := uuid.Parse(r.PathValue(pathKeytournamentID.String()))
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	handler.tournamentPage(ctx, url, tournamentID).Render(w)
}

func (handler *webServer) tournamentPage(ctx context.Context, url url.URL, tournamentID uuid.UUID) g.Node {
	var err error
	tournament, err := models.FindTournament(ctx, handler.db, tournamentID)
	if err != nil {
		return errorAlert(err)
	}
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: tournament.Title,
		Body:  h.Div(),
	})
}
