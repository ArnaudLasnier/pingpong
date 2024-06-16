package web

import (
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
	"github.com/aarondl/opt/omit"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (handler *webServer) createPlayerModalHandlerFunc(w http.ResponseWriter, r *http.Request) {
	err := modal("Create Player", handler.createPlayerForm(webutils.Form{})).Render(w)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
}

func (handler *webServer) createPlayerFormHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	username := r.PostFormValue(formKeyPlayerUsername.String())
	form := webutils.Form{
		IsSubmitted: true,
		Fields: webutils.FormFields{
			formKeyPlayerUsername: webutils.NewValidFormValue(username),
		},
	}
	numberOfPlayersWithSameUsername, err := models.Players.Query(
		ctx,
		handler.db,
		sm.Where(
			psql.Quote(models.ColumnNames.Players.Username).EQ(psql.Arg(username)),
		),
	).Count()
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	if numberOfPlayersWithSameUsername != 0 {
		emailField := form.Fields[formKeyPlayerUsername]
		emailField.IsValid = false
		emailField.ValidationMessage = "This username already exists."
		form.Fields[formKeyPlayerUsername] = emailField
		handler.createPlayerForm(form).Render(w)
		return
	}
	_, err = handler.service.CreatePlayer(ctx, &models.PlayerSetter{
		Username: omit.From(username),
	})
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	successAlert().Render(w)
}

func (handler *webServer) createPlayerForm(form webutils.Form) g.Node {
	return h.FormEl(
		hx.Post(fragmentCreatePlayerForm.Endpoint()),
		hx.Swap("outerHTML"),
		h.Div(
			h.Class("mb-4"),
			h.Label(h.For(formKeyPlayerUsername.String()), h.Class("form-label"), g.Text("Username")),
			h.Input(
				h.ID(formKeyPlayerUsername.String()),
				h.Name(formKeyPlayerUsername.String()),
				h.Type(formKeyPlayerUsername.String()),
				h.Required(),
				h.Value(form.Fields[formKeyPlayerUsername].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[formKeyPlayerUsername].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[formKeyPlayerUsername].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[formKeyPlayerUsername].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[formKeyPlayerUsername].IsValid,
				},
				g.Text(form.Fields[formKeyPlayerUsername].ValidationMessage),
			),
		),
		h.Div(
			h.Class("d-flex justify-content-end"),
			h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
		),
	)
}
