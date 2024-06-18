package web

import (
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
	"github.com/google/uuid"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
)

func (server *webServer) deleteTournament(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	tournamentID, err := uuid.Parse(r.PostFormValue(pathKeyTournamentID.String()))
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	tournament, err := models.FindTournament(ctx, server.db, tournamentID)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	err = tournament.Delete(ctx, server.db)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	w.Header().Set(webutils.HeaderHxTrigger, eventTournamentDeleted.String())
	successAlert().Render(w)
}

func (server *webServer) deleteTournamentModalHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	tournamentID, err := uuid.Parse(r.FormValue(pathKeyTournamentID.String()))
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	tournament, err := models.FindTournament(ctx, server.db, tournamentID)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	err = deleteTournamentModal(tournament).Render(w)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
}

func deleteTournamentModal(tournament *models.Tournament) g.Node {
	return modal(
		"Delete Tournament",
		h.Div(
			h.ID(fragmentDeleteTournamentModalContent.String()),
			h.P(
				g.Rawf("Are you sure you want to delete tournament <strong>%s</strong>?", tournament.Title),
			),
			h.FormEl(
				h.Input(
					h.Name(formKeyTournamentID.String()),
					h.Value(tournament.ID.String()),
					displayNone(),
				),
				h.Button(
					hx.Post(fragmentDeleteTournament.Endpoint()),
					hx.Swap("outerHTML"),
					hx.Target(fragmentDeleteTournamentModalContent.IDSelector()),
					h.Class("btn btn-danger"),
					g.Text("Confirm"),
				),
			),
		),
	)
}
