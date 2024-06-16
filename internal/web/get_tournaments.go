package web

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/aarondl/opt/null"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

const (
	createTournamentModalID       = "create-tournament-modal"
	createTournamentModalSelector = "#create-tournament-modal"
)

func (handler *webServer) tournamentsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url := r.URL
	err := handler.tournamentsPage(ctx, *url).Render(w)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
}

func (handler *webServer) tournamentsPage(ctx context.Context, url url.URL) g.Node {
	var err error
	tournaments, err := models.Tournaments.Query(ctx, handler.db, sm.OrderBy(models.ColumnNames.Tournaments.StartedAt), sm.Limit(10)).All()
	if err != nil {
		return errorAlert(err)
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
			modalPlaceholder(createTournamentModalID),
			h.Div(
				h.Class("d-flex justify-content-center"),
				h.Table(
					h.Class("table w-75"),
					h.THead(
						h.Class("table-light"),
						h.Tr(
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Title")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Status")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("# Players")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Start Date")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("End Date")),
							h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Actions")),
						),
					),
					h.TBody(
						g.Group(g.Map(tournaments, func(tournament *models.Tournament) g.Node {
							return handler.tournamentRow(ctx, tournament)
						})),
					),
				),
			),
			modalPlaceholder("add-participant-modal"),
		),
	})
}

func (handler *webServer) tournamentRow(ctx context.Context, tournament *models.Tournament) g.Node {
	return h.Tr(
		h.Td(
			// h.A(
			// 	h.Href("/tournaments/"+tournament.ID.String()),
			// 	g.Text(tournament.Title),
			// ),
			g.Text(tournament.Title),
		),
		h.Td(tournamentStatusBadge(tournament.Status)),
		h.Td(g.Text(handler.playerCount(ctx, tournament))),
		h.Td(g.Text(formatNullTime(tournament.StartedAt))),
		h.Td(g.Text(formatNullTime(tournament.EndedAt))),
		h.Td(
			h.FormEl(
				h.Input(
					h.Name(formKeyTournamentID.String()),
					h.Value(tournament.ID.String()),
					displayNone(),
				),
				h.Button(
					hx.Post(formActionStartTournament.Path()),
					h.Class("btn btn-sm btn-primary"),
					g.Text("Start"),
				),
			),
		),
	)
}

func (handler *webServer) playerCount(ctx context.Context, tournament *models.Tournament) string {
	playerCount, err := tournament.Players(ctx, handler.db).Count()
	if err != nil {
		return "-"
	}
	return strconv.Itoa(int(playerCount))
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
	}
	return t.MustGet().Format(time.DateOnly)
}
