package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/aarondl/opt/omit"
	"github.com/google/uuid"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (handler *handler) addParticipantModalHandler(w http.ResponseWriter, r *http.Request) {
	err := Modal("Add Participant", handler.addParticipantForm(r.Context(), Form{})).Render(w)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
}

func (handler *handler) addParticipantFormHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	firstName := r.PostFormValue(PlayerFirstName.String())
	lastName := r.PostFormValue(PlayerLastName.String())
	email := r.PostFormValue(PlayerEmail.String())
	form := Form{
		IsSubmitted: true,
		Fields: FormFields{
			PlayerFirstName: NewValidValue(firstName),
			PlayerLastName:  NewValidValue(lastName),
			PlayerEmail:     NewValidValue(email),
		},
	}
	numberOfPlayersWithSameEmail, err := models.Players.Query(
		ctx,
		handler.db,
		sm.Where(
			psql.Quote(models.ColumnNames.Players.Email).EQ(psql.Arg(email)),
		),
	).Count()
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	if numberOfPlayersWithSameEmail != 0 {
		emailField := form.Fields[PlayerEmail]
		emailField.IsValid = false
		emailField.Message = "This email address already exists."
		form.Fields[PlayerEmail] = emailField
		handler.createPlayerForm(form).Render(w)
		return
	}
	_, err = handler.tournamentService.CreatePlayer(ctx, &models.PlayerSetter{
		FirstName: omit.From(firstName),
		LastName:  omit.From(lastName),
		Email:     omit.From(email),
	})
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	SuccessAlert().Render(w)
}

func (handler *handler) addParticipantForm(ctx context.Context, _ Form) g.Node {
	var err error
	players, err := models.Players.Query(ctx, handler.db).All()
	if err != nil {
		return ErrorAlert(err)
	}
	return h.FormEl(
		hx.Post("/add-participant-modal/form"),
		hx.Swap("outerHTML"),
		h.Div(
			h.Class("mb-4"),
			h.Select(
				h.Class("form-select"),
				h.Multiple(),
				g.Group(
					g.Map(players, func(player *models.Player) g.Node {
						return h.Option(
							h.Value(player.ID.String()),
							g.Text(fmt.Sprintf("%s %s - %s", player.FirstName, player.LastName, player.Email)),
						)
					}),
				),
			),
		),
		h.Div(
			h.Class("d-flex justify-content-end"),
			h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
		),
	)
}

func (handler *handler) addParticipants(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	tournamentID, err := uuid.Parse(r.PathValue(tournamentID.String()))
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	err = r.ParseForm()
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	participantEmails := r.PostForm["participantEmails"]
	playersToAdd, err := models.Players.Query(
		ctx,
		handler.db,
		sm.Where(
			models.PlayerColumns.Email.In(psql.Arg(participantEmails)),
		),
	).All()
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	tournament, err := models.FindTournament(ctx, handler.db, tournamentID)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	err = tournament.AttachPlayers(ctx, handler.db, playersToAdd...)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
}
