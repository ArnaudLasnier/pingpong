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

func (handler *webServer) createPlayerModalHandlerFunc(w http.ResponseWriter, r *http.Request) {
	err := Modal("Create Player", handler.createPlayerForm(form{})).Render(w)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
}

func (handler *webServer) createPlayerFormHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	firstName := r.PostFormValue(formKeyPlayerFirstName.String())
	lastName := r.PostFormValue(formKeyPlayerLastName.String())
	email := r.PostFormValue(formKeyPlayerEmail.String())
	form := form{
		IsSubmitted: true,
		Fields: formFields{
			formKeyPlayerFirstName: newValidFormValue(firstName),
			formKeyPlayerLastName:  newValidFormValue(lastName),
			formKeyPlayerEmail:     newValidFormValue(email),
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
		emailField := form.Fields[formKeyPlayerEmail]
		emailField.IsValid = false
		emailField.ValidationMessage = "This email address already exists."
		form.Fields[formKeyPlayerEmail] = emailField
		handler.createPlayerForm(form).Render(w)
		return
	}
	_, err = handler.service.CreatePlayer(ctx, &models.PlayerSetter{
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

func (handler *webServer) createPlayerForm(form form) g.Node {
	return h.FormEl(
		hx.Post(createPlayerFormResource.Endpoint()),
		hx.Swap("outerHTML"),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For(formKeyPlayerFirstName.String()), h.Class("form-label"), g.Text("First Name")),
			h.Input(
				h.ID(formKeyPlayerFirstName.String()),
				h.Name(formKeyPlayerFirstName.String()),
				h.Type("text"),
				h.Required(),
				h.Pattern("[A-Za-z0-9]{1,50}"),
				h.Value(form.Fields[formKeyPlayerFirstName].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[formKeyPlayerFirstName].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[formKeyPlayerFirstName].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[formKeyPlayerFirstName].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[formKeyPlayerFirstName].IsValid,
				},
				g.Text(form.Fields[formKeyPlayerFirstName].ValidationMessage),
			),
		),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For("playerLastName"), h.Class("form-label"), g.Text("Last Name")),
			h.Input(
				h.ID("playerLastName"),
				h.Name(formKeyPlayerLastName.String()),
				h.Type("text"),
				h.Required(),
				h.Pattern("[A-Za-z0-9]{1,50}"),
				h.Value(form.Fields[formKeyPlayerLastName].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[formKeyPlayerLastName].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[formKeyPlayerLastName].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[formKeyPlayerLastName].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[formKeyPlayerLastName].IsValid,
				},
				g.Text(form.Fields[formKeyPlayerLastName].ValidationMessage),
			),
		),
		h.Div(
			h.Class("mb-4"),
			h.Label(h.For(formKeyPlayerEmail.String()), h.Class("form-label"), g.Text("Email")),
			h.Input(
				h.ID(formKeyPlayerEmail.String()),
				h.Name(formKeyPlayerEmail.String()),
				h.Type(formKeyPlayerEmail.String()),
				h.Required(),
				h.Value(form.Fields[formKeyPlayerEmail].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[formKeyPlayerEmail].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[formKeyPlayerEmail].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[formKeyPlayerEmail].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[formKeyPlayerEmail].IsValid,
				},
				g.Text(form.Fields[formKeyPlayerEmail].ValidationMessage),
			),
		),
		h.Div(
			h.Class("d-flex justify-content-end"),
			h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
		),
	)
}
