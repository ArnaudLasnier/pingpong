package web

import (
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/aarondl/opt/omit"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (handler *handler) createPlayerModalHandler(w http.ResponseWriter, r *http.Request) {
	err := Modal("Create Player", handler.createPlayerForm(Form{})).Render(w)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
}

func (handler *handler) createPlayerFormHandler(w http.ResponseWriter, r *http.Request) {
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

func (handler *handler) createPlayerForm(form Form) g.Node {
	return h.FormEl(
		hx.Post("/create-player-modal/form"),
		hx.Swap("outerHTML"),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For("playerFirstName"), h.Class("form-label"), g.Text("First Name")),
			h.Input(
				h.ID(PlayerFirstName.String()),
				h.Name(PlayerFirstName.String()),
				h.Type("text"),
				h.Required(),
				h.Pattern("[A-Za-z0-9]{1,50}"),
				h.Value(form.Fields[PlayerFirstName].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[PlayerFirstName].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[PlayerFirstName].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[PlayerFirstName].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[PlayerFirstName].IsValid,
				},
				g.Text(form.Fields[PlayerFirstName].Message),
			),
		),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For("playerLastName"), h.Class("form-label"), g.Text("Last Name")),
			h.Input(
				h.ID("playerLastName"),
				h.Name(PlayerLastName.String()),
				h.Type("text"),
				h.Required(),
				h.Pattern("[A-Za-z0-9]{1,50}"),
				h.Value(form.Fields[PlayerLastName].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[PlayerLastName].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[PlayerLastName].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[PlayerLastName].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[PlayerLastName].IsValid,
				},
				g.Text(form.Fields[PlayerLastName].Message),
			),
		),
		h.Div(
			h.Class("mb-4"),
			h.Label(h.For(PlayerEmail.String()), h.Class("form-label"), g.Text("Email")),
			h.Input(
				h.ID(PlayerEmail.String()),
				h.Name(PlayerEmail.String()),
				h.Type(PlayerEmail.String()),
				h.Required(),
				h.Value(form.Fields[PlayerEmail].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[PlayerEmail].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[PlayerEmail].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[PlayerEmail].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[PlayerEmail].IsValid,
				},
				g.Text(form.Fields[PlayerEmail].Message),
			),
		),
		h.Div(
			h.Class("d-flex justify-content-end"),
			h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
		),
	)
}
