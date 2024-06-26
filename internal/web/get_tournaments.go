package web

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
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

func (server *webServer) tournamentsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url := r.URL
	err := server.tournamentsPage(ctx, *url).Render(w)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
}

func (server *webServer) tournamentsPage(ctx context.Context, url url.URL) g.Node {
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: "Tournaments",
		Body: h.Div(
			h.H1(
				h.StyleAttr("font-family: 'Source Serif'"),
				h.Class("mb-5"),
				g.Text("Tournaments"),
			),
			h.Div(
				h.Class("mb-3 d-flex justify-content-end"),
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
			server.tournamentsTable(ctx),
			modalPlaceholder(createTournamentModalID),
			modalPlaceholder("add-participant-modal"),
			modalPlaceholder(fragmentDeleteTournamentModal.String()),
		),
	})
}

func (server *webServer) tournamentsTableHandlerFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	server.tournamentsTable(ctx).Render(w)
}

func (server *webServer) tournamentsTable(ctx context.Context) g.Node {
	var err error
	tournaments, err := models.Tournaments.Query(ctx, server.db, sm.OrderBy(models.ColumnNames.Tournaments.StartedAt), sm.Limit(10)).All()
	if err != nil {
		return errorAlert(err)
	}
	return h.Div(
		hx.Trigger(webutils.JoinEvents(eventTournamentCreated, eventTournamentDeleted, eventTournamentStarted)),
		hx.Get(fragmentTournamentsTable.Endpoint()),
		hx.Swap("outerHTML"),
		h.Class("d-flex justify-content-center"),
		h.Table(
			h.Class("table table-striped"),
			h.THead(
				h.Class("table-primary"),
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
					return server.tournamentRow(ctx, tournament)
				})),
			),
		),
	)
}

func (handler *webServer) tournamentRow(ctx context.Context, tournament *models.Tournament) g.Node {
	return h.Tr(
		h.Td(
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
					h.Class("btn btn-sm btn-primary me-3"),
					g.Text("Start"),
				),
				h.Button(
					hx.Get(fragmentDeleteTournamentModal.Endpoint()),
					hx.Include("closest form"),
					hx.Target(fragmentDeleteTournamentModal.IDSelector()),
					hx.Swap("innerHTML"),
					g.Attr("data-bs-toggle", "modal"),
					g.Attr("data-bs-target", fragmentDeleteTournamentModal.IDSelector()),
					h.Class("btn btn-sm btn-danger"),
					g.Text("Delete"),
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
