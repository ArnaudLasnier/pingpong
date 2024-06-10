package web

import (
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (handler *handler) createTournamentModalHandler(w http.ResponseWriter, r *http.Request) {
	err := Modal("Create Tournament", handler.createTournamentForm(Form{})).Render(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *handler) createTournamentFormHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	title := r.PostFormValue(TournamentTitle.String())
	form := Form{
		IsSubmitted: true,
		Fields: FormFields{
			TournamentTitle: NewValidValue(title),
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
		ErrorAlert(err).Render(w)
		return
	}
	if numberOfTournamentsWithSameTitle > 0 {
		form.Fields[TournamentTitle] = NewInvalidValue(form.Fields[TournamentTitle].Value, "This title already exists.")
		handler.createTournamentForm(form).Render(w)
		return
	}
	_, err = handler.tournamentService.CreateTournamentDraft(ctx, title)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
	SuccessAlert().Render(w)
}

func (handler *handler) createTournamentForm(form Form) g.Node {
	titleFieldName := "Title"
	return h.FormEl(
		hx.Post("/create-tournament-modal/form"),
		hx.Swap("outerHTML"),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For(TournamentTitle.String()), h.Class("form-label"), g.Text(titleFieldName)),
			h.Input(
				h.ID(TournamentTitle.String()),
				h.Name(TournamentTitle.String()),
				h.Type("text"),
				h.Required(),
				h.Value(form.Fields[TournamentTitle].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[TournamentTitle].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[TournamentTitle].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[TournamentTitle].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[TournamentTitle].IsValid,
				},
				g.Text(form.Fields[TournamentTitle].Message),
			),
		),
		h.Div(
			h.Class("d-flex justify-content-end"),
			h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
		),
	)
}
