package web

import (
	"context"
	"net/http"
	"slices"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/google/uuid"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (server *webServer) registerPlayerModalHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	playerID, err := uuid.Parse(r.PathValue(pathKeyPlayerID.String()))
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	player, err := models.FindPlayer(ctx, server.db, playerID)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	err = modal("Register Player", server.registerPlayerForms(r.Context(), player)).Render(w)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
}

func (server *webServer) registerPlayerForms(ctx context.Context, player *models.Player) g.Node {
	var err error
	tournamentDrafts, err := models.Tournaments.Query(
		ctx,
		server.db,
		sm.Where(models.TournamentColumns.Status.EQ(psql.Arg(models.TournamentStatusDraft))),
		// models.ThenLoadTournamentParticipationParticipantPlayer(),
	).All()
	if err != nil {
		return errorAlert(err)
	}
	return h.Div(
		h.ID(fragmentRegisterPlayerForm.String()),
		g.Group(
			g.Map(tournamentDrafts, func(tournamentDraft *models.Tournament) g.Node {
				participantPlayers, err := tournamentDraft.Players(ctx, server.db).All()
				if err != nil {
					return errorAlert(err)
				}
				var participantPlayerIDs []uuid.UUID
				for _, participantPlayer := range participantPlayers {
					participantPlayerIDs = append(participantPlayerIDs, participantPlayer.ID)
				}
				isRegistered := slices.Contains(participantPlayerIDs, player.ID)
				return h.FormEl(
					h.Class("d-flex justify-content-between align-items-center mb-2"),
					h.Span(
						h.Class("h-100"),
						g.Text(tournamentDraft.Title),
					),
					h.Input(
						h.Name(formKeyPlayerID.String()),
						h.Value(player.ID.String()),
						displayNone(),
					),
					h.Input(
						h.Name(formKeyTournamentID.String()),
						h.Value(tournamentDraft.ID.String()),
						displayNone(),
					),
					g.If(isRegistered, deregisterPlayerButton()),
					g.If(!isRegistered, registerPlayerButton()),
				)
			}),
		),
	)
}

func (server *webServer) registerPlayerButtonHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	playerID, err := uuid.Parse(r.PostFormValue(pathKeyPlayerID.String()))
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	player, err := models.FindPlayer(ctx, server.db, playerID)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	tournamentID, err := uuid.Parse(r.PostFormValue(pathKeytournamentID.String()))
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	tournament, err := models.FindTournament(ctx, server.db, tournamentID)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	err = tournament.AttachPlayers(ctx, server.db, player)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	deregisterPlayerButton().Render(w)
}

func (server *webServer) deregisterPlayerButtonHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	playerID, err := uuid.Parse(r.PostFormValue(pathKeyPlayerID.String()))
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	tournamentID, err := uuid.Parse(r.PostFormValue(pathKeytournamentID.String()))
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	_, err = models.TournamentParticipations.DeleteQ(
		ctx,
		server.db,
		dm.Where(
			models.TournamentParticipationColumns.TournamentID.EQ(psql.Arg(tournamentID.String())).And(
				models.TournamentParticipationColumns.ParticipantID.EQ(psql.Arg(playerID.String())),
			),
		),
	).Exec()
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	registerPlayerButton().Render(w)
}

func registerPlayerButton() g.Node {
	return h.Button(
		hx.Post(fragmentRegisterPlayerButton.Endpoint()),
		hx.Swap("outerHTML"),
		h.Class("btn btn-sm btn-outline-success"),
		g.Text("Register"),
	)
}

func deregisterPlayerButton() g.Node {
	return h.Button(
		hx.Post(fragmentDeregisterPlayerButton.Endpoint()),
		hx.Swap("outerHTML"),
		h.Class("btn btn-sm btn-outline-danger"),
		g.Text("Deregister"),
	)
}
