package web

import (
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/google/uuid"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
)

func (server *webServer) deletePlayerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	playerID, err := uuid.Parse(r.PathValue(playerID.String()))
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	player, err := models.FindPlayer(ctx, server.db, playerID)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	err = player.Delete(ctx, server.db)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	SuccessAlert().Render(w)
}

func (server *webServer) deletePlayerModalHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	playerID, err := uuid.Parse(r.PathValue(playerID.String()))
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	player, err := models.FindPlayer(ctx, server.db, playerID)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	err = deletePlayerModal(player).Render(w)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
}

func deletePlayerModal(player *models.Player) g.Node {
	deletePlayerModalContent := "delete-player-modal-content"
	return Modal(
		"Delete Player",
		h.Div(
			h.ID(deletePlayerModalContent),
			h.P(
				g.Rawf("Are you sure you want to delete player <strong>%s %s</strong>?", player.FirstName, player.LastName),
			),
			h.Button(
				hx.Delete("/players/"+player.ID.String()),
				hx.Swap("outerHTML"),
				hx.Target("#"+deletePlayerModalContent),
				h.Class("btn btn-danger"),
				g.Text("Confirm"),
			),
		),
	)
}
