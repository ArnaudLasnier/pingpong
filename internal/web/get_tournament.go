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

func (handler *handler) tournament(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	url := *r.URL
	tournamentID, err := uuid.Parse(r.PathValue(tournamentID.String()))
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	handler.tournamentPage(ctx, url, tournamentID).Render(w)
}

func (handler *handler) tournamentPage(ctx context.Context, url url.URL, tournamentID uuid.UUID) g.Node {
	var err error
	tournament, err := models.FindTournament(ctx, handler.db, tournamentID)
	if err != nil {
		return ErrorAlert(err)
	}
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: tournament.Title,
		Body:  h.Div(),
	})
}
