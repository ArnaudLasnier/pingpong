package web

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/aarondl/opt/null"
	"github.com/google/uuid"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

const (
	createTournamentModalID       = "create-tournament-modal"
	createTournamentModalSelector = "#create-tournament-modal"
)

// type TournamentCreationFormFields struct {
// 	FirstName string
// }

var tournamentCreationFormFields struct {
	FirstName string
	// Last
}

func (handler *handler) tournamentsPage(ctx context.Context, url url.URL) g.Node {
	var err error
	tournaments, err := models.Tournaments.Query(ctx, handler.db, sm.OrderBy(models.ColumnNames.Tournaments.StartedAt), sm.Limit(10)).All()
	if err != nil {
		todoPanic(err)
	}
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: "Tournaments",
		Body: h.Div(
			h.H1(g.Text("Tournaments")),
			h.Div(
				h.Class("mb-3"),
				h.Button(
					hx.Get("/create-tournament-modal"),
					hx.Target(createTournamentModalSelector),
					hx.Trigger("click"),
					g.Attr("data-bs-toggle", "modal"),
					g.Attr("data-bs-target", createTournamentModalSelector),
					h.Class("btn btn-primary"),
					g.Text("Create Tournament"),
				),
			),
			h.Div(
				h.ID(createTournamentModalID),
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
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Title")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Status")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Start Date")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("End Date")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Actions")),
						),
					),
					h.TBody(
						g.Group(g.Map(tournaments, func(tournament *models.Tournament) g.Node {
							return h.Tr(
								h.Td(g.Text(tournament.Title)),
								h.Td(tournamentStatusBadge(tournament.Status)),
								h.Td(g.Text(formatNullTime(tournament.StartedAt))),
								h.Td(g.Text(formatNullTime(tournament.EndedAt))),
								h.Td(
									// h.Button(
									// 	h.Class("btn btn-sm btn-primary"),
									// 	hx.Post(fmt.Sprintf("/add-participant-to-tournament/%s", tournament.ID.String())),
									// 	g.Text("Add Participant"),
									// ),
									h.Button(
										hx.Get("/add-participant-modal"),
										hx.Target("#add-participant-modal"),
										hx.Trigger("click"),
										g.Attr("data-bs-toggle", "modal"),
										g.Attr("data-bs-target", "#add-participant-modal"),
										h.Class("btn btn-sm btn-primary"),
										g.Text("Add Participant"),
									),
								),
							)
						})),
					),
				),
			),
			h.Div(
				h.ID("add-participant-modal"),
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
		),
	})
}

func (handler *handler) addParticipantModal(w http.ResponseWriter, r *http.Request) {

}

func (handler *handler) addParticipantModalForm(w http.ResponseWriter, r *http.Request) {

}

func tournamentStatusBadge(status models.TournamentStatus) g.Node {
	var badgeClass string
	if status == models.TournamentStatusDraft {
		badgeClass = "text-bg-secondary"
	} else if status == models.TournamentStatusStarted {
		badgeClass = "text-bg-warning"
	} else if status == models.TournamentStatusEnded {
		badgeClass = "text-bg-success"
	} else {
		badgeClass = "text-bg-light"
	}
	return h.Span(
		c.Classes{"badge": true, badgeClass: true},
		g.Text(strings.ToUpper(string(status))),
	)
}

func formatNullTime(t null.Val[time.Time]) string {
	if t.IsNull() {
		return "-"
	} else {
		return t.MustGet().Format(time.DateOnly)
	}
}

func (handler *handler) tournaments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url := r.URL
	err := handler.tournamentsPage(ctx, *url).Render(w)
	if err != nil {
		todoPanic(err)
	}
}

func (handler *handler) tournamentCreationModal() g.Node {
	return h.Div(
		h.Class("modal-dialog modal-dialog-centered"),
		h.Div(
			h.Class("modal-content"),
			h.Div(
				h.Class("modal-header d-flex justify-content-between"),
				h.H5(h.Class("modal-title"), g.Text("Create Tournament")),
				h.Button(h.Type("button"), h.Class("btn-close"), g.Attr("data-bs-dismiss", "modal")),
			),
			h.Div(
				h.Class("modal-body"),
				handler.tournamentCreationForm(Form{}),
			),
		),
	)
}

func (handler *handler) tournamentCreationForm(form Form) g.Node {
	titleFieldID := FormFieldIDTournamentTitle
	titleFieldName := "Title"
	return h.FormEl(
		hx.Post("/create-tournament-modal/form"),
		hx.Swap("outerHTML"),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For(titleFieldID), h.Class("form-label"), g.Text(titleFieldName)),
			h.Input(
				h.ID(titleFieldID),
				h.Name(titleFieldID),
				h.Type("text"),
				h.Required(),
				h.Value(form.Fields[titleFieldID].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[titleFieldID].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[titleFieldID].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[titleFieldID].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[titleFieldID].IsValid,
				},
				g.Text(form.Fields[titleFieldID].Message),
			),
		),
		h.Div(
			h.Class("d-flex justify-content-end"),
			h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
		),
	)
}

func (handler *handler) createTournamentModal(w http.ResponseWriter, r *http.Request) {
	err := handler.tournamentCreationModal().Render(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *handler) createTournamentModalForm(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	title := r.PostFormValue(FormFieldIDTournamentTitle)
	form := Form{
		IsSubmitted: true,
		Fields: map[string]FormField{
			FormFieldIDTournamentTitle: {
				Value:   title,
				IsValid: true,
				Message: FormFieldOKMessage,
			},
		},
	}
	numberOfTournamentsWithSameTitle, err := models.Tournaments.Query(
		ctx,
		handler.db,
		sm.Where(
			models.TournamentColumns.Title.EQ(psql.Arg(title)),
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
	if numberOfTournamentsWithSameTitle != 0 {
		emailField := form.Fields["email"]
		emailField.IsValid = false
		emailField.Message = "This email address already exists."
		form.Fields["email"] = emailField
		handler.tournamentCreationForm(form).Render(w)
		return
	}
	_, err = handler.tournamentService.CreateTournamentDraft(ctx, title)
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

func (handler *handler) addParticipants(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	tournamentID, err := uuid.FromBytes([]byte(r.PathValue("tournamentID")))
	if err != nil {
		ErrorAlert(err).Render(w)
	}
	err = r.ParseForm()
	if err != nil {
		ErrorAlert(err).Render(w)
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
	}
	tournament, err := models.FindTournament(ctx, handler.db, tournamentID)
	if err != nil {
		ErrorAlert(err).Render(w)
	}
	err = tournament.AttachPlayers(ctx, handler.db, playersToAdd...)
	if err != nil {
		ErrorAlert(err).Render(w)
	}
}

func SuccessAlert() g.Node {
	return h.Div(
		h.Class("alert alert-success"),
		h.Role("alert"),
		g.Text("Success!"),
	)
}

func ErrorAlert(err error) g.Node {
	return h.Div(
		h.Class("alert alert-danger"),
		h.Role("alert"),
		h.H5(
			h.Class("alert-heading"),
			g.Text("Error"),
		),
		h.P(g.Text(err.Error())),
	)
}
