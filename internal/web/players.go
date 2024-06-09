package web

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/aarondl/opt/omit"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (handler *handler) playersPage(ctx context.Context, url url.URL) g.Node {
	var err error
	players, err := models.Players.Query(ctx, handler.db, sm.OrderBy(models.ColumnNames.Players.LastName), sm.Limit(10)).All()
	if err != nil {
		todoPanic(err)
	}
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: "Players",
		Body: h.Div(
			h.H1(g.Text("Players")),
			h.Div(
				h.Class("mb-3"),
				h.Button(
					hx.Get("/add-player-modal"),
					hx.Target("#add-player-modal"),
					hx.Trigger("click"),
					g.Attr("data-bs-toggle", "modal"),
					g.Attr("data-bs-target", "#add-player-modal"),
					h.Class("btn btn-primary"),
					g.Text("Create Player"),
				),
			),
			h.Div(
				h.ID("add-player-modal"),
				h.Class("modal modal-blur fade"),
				h.StyleAttr("display: none"),
				g.Attr("aria-hidden", "false"),
				g.Attr("tab-index", "-1"),
				h.Div(
					h.Class("modal-dialog modal-lg modal-dialog-centered"),
					h.Role("document"),
					h.Div(h.Class("modal-content")),
				),
			),
			h.Div(
				h.Class("d-flex justify-content-center"),
				h.Table(
					h.Class("table w-75"),
					h.THead(
						h.Class("table-light"),
						h.Tr(
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("First Name")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Last Name")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Email")),
						),
					),
					h.TBody(
						g.Group(g.Map(players, func(player *models.Player) g.Node {
							return h.Tr(
								h.Td(g.Text(player.FirstName)),
								h.Td(g.Text(player.LastName)),
								h.Td(g.Text(player.Email)),
							)
						})),
					),
				),
			),
		),
	})
}

func (handler *handler) handleGetPlayersPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url := r.URL
	err := handler.playersPage(ctx, *url).Render(w)
	if err != nil {
		todoPanic(err)
	}
}

func (handler *handler) playerCreationModal(form Form) g.Node {
	return h.Div(
		h.Class("modal-dialog modal-dialog-centered"),
		h.Div(
			h.Class("modal-content"),
			h.Div(
				h.Class("modal-header d-flex justify-content-between"),
				h.H5(h.Class("modal-title"), g.Text("Create Player")),
				h.Button(h.Type("button"), h.Class("btn-close"), g.Attr("data-bs-dismiss", "modal")),
			),
			h.Div(
				h.Class("modal-body"),
				handler.playerCreationForm(form),
			),
		),
	)
}

func (handler *handler) playerCreationForm(form Form) g.Node {
	return h.FormEl(
		hx.Post("/add-player-modal/form"),
		hx.Swap("outerHTML"),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For("playerFirstName"), h.Class("form-label"), g.Text("First Name")),
			h.Input(
				h.ID("playerFirstName"),
				h.Name("firstName"),
				h.Type("text"),
				h.Required(),
				h.Pattern("[A-Za-z0-9]{1,50}"),
				h.Value(form.Fields["firstName"].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields["firstName"].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields["firstName"].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields["firstName"].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields["firstName"].IsValid,
				},
				g.Text(form.Fields["firstName"].Message),
			),
		),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For("playerLastName"), h.Class("form-label"), g.Text("Last Name")),
			h.Input(
				h.ID("playerLastName"),
				h.Name("lastName"),
				h.Type("text"),
				h.Required(),
				h.Pattern("[A-Za-z0-9]{1,50}"),
				h.Value(form.Fields["lastName"].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields["lastName"].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields["lastName"].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields["lastName"].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields["lastName"].IsValid,
				},
				g.Text(form.Fields["lastName"].Message),
			),
		),
		h.Div(
			h.Class("mb-4"),
			h.Label(h.For("playerEmail"), h.Class("form-label"), g.Text("Email")),
			h.Input(
				h.ID("playerEmail"),
				h.Name("email"),
				h.Type("email"),
				h.Required(),
				h.Value(form.Fields["email"].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields["email"].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields["email"].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields["email"].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields["email"].IsValid,
				},
				g.Text(form.Fields["email"].Message),
			),
		),
		h.Div(
			h.Class("d-flex justify-content-end"),
			h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
		),
	)
}

func (handler *handler) handleGetCreatePlayerModal(w http.ResponseWriter, r *http.Request) {
	err := handler.playerCreationModal(Form{}).Render(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *handler) handlePostPlayerCreationForm(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	firstName := r.PostFormValue("firstName")
	lastName := r.PostFormValue("lastName")
	email := r.PostFormValue("email")
	form := Form{
		IsSubmitted: true,
		Fields: map[string]FormField{
			"firstName": {
				Value:   firstName,
				IsValid: true,
				Message: "Looks good!",
			},
			"lastName": {
				Value:   lastName,
				IsValid: true,
				Message: "Looks good!",
			},
			"email": {
				Value:   email,
				IsValid: true,
				Message: "Looks good!",
			},
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
		h.Div(
			h.Class("alert alert-danger"),
			h.Role("alert"),
			h.H5(
				h.Class("alert-heading"),
				g.Text("Error"),
			),
			h.P(g.Text(err.Error())),
		).Render(w)
		return
	}
	if numberOfPlayersWithSameEmail != 0 {
		emailField := form.Fields["email"]
		emailField.IsValid = false
		emailField.Message = "This email address already exists."
		form.Fields["email"] = emailField
		handler.playerCreationForm(form).Render(w)
		return
	}
	_, err = handler.tournamentService.CreatePlayer(ctx, &models.PlayerSetter{
		FirstName: omit.From(firstName),
		LastName:  omit.From(lastName),
		Email:     omit.From(email),
	})
	if err != nil {
		h.Div(
			h.Class("alert alert-danger"),
			h.Role("alert"),
			h.H5(
				h.Class("alert-heading"),
				g.Text("Error"),
			),
			h.P(g.Text(err.Error())),
		).Render(w)
		return
	}
	h.Div(
		h.Class("alert alert-success"),
		h.Role("alert"),
		g.Text("Success!"),
	).Render(w)
}
