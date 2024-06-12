package web

import (
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (handler *webServer) createTournamentModalHandlerFunc(w http.ResponseWriter, r *http.Request) {
	err := modal("Create Tournament", handler.createTournamentForm(webutils.Form{})).Render(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *webServer) createTournamentFormHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	title := r.PostFormValue(formKeyTournamentTitle.String())
	form := webutils.Form{
		IsSubmitted: true,
		Fields: webutils.FormFields{
			formKeyTournamentTitle: webutils.NewValidFormValue(title),
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
		errorAlert(err).Render(w)
		return
	}
	if numberOfTournamentsWithSameTitle > 0 {
		form.Fields[formKeyTournamentTitle] = webutils.NewInvalidFormValue(form.Fields[formKeyTournamentTitle].Value, "This title already exists.")
		handler.createTournamentForm(form).Render(w)
		return
	}
	_, err = handler.service.CreateTournamentDraft(ctx, title)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
	successAlert().Render(w)
}

func (handler *webServer) createTournamentForm(form webutils.Form) g.Node {
	titleFieldName := "Title"
	return h.FormEl(
		hx.Post(fragmentCreateTournamentForm.Endpoint()),
		hx.Swap("outerHTML"),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For(formKeyTournamentTitle.String()), h.Class("form-label"), g.Text(titleFieldName)),
			h.Input(
				h.ID(formKeyTournamentTitle.String()),
				h.Name(formKeyTournamentTitle.String()),
				h.Type("text"),
				h.Required(),
				h.Value(form.Fields[formKeyTournamentTitle].Value),
				c.Classes{
					"form-control": true,
					"is-valid":     form.IsSubmitted && form.Fields[formKeyTournamentTitle].IsValid,
					"is-invalid":   form.IsSubmitted && !form.Fields[formKeyTournamentTitle].IsValid,
				},
			),
			h.Div(
				c.Classes{
					"valid-feedback":   form.IsSubmitted && form.Fields[formKeyTournamentTitle].IsValid,
					"invalid-feedback": form.IsSubmitted && !form.Fields[formKeyTournamentTitle].IsValid,
				},
				g.Text(form.Fields[formKeyTournamentTitle].ValidationMessage),
			),
		),
		h.Div(
			h.Class("d-flex justify-content-end"),
			h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
		),
	)
}
