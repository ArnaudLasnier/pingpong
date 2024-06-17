package web

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (server *webServer) playersHandlerFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url := r.URL
	err := server.playersPage(ctx, *url).Render(w)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
}

func (server *webServer) playersPage(ctx context.Context, url url.URL) g.Node {
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: "Players",
		Body: h.Div(
			h.H1(
				h.StyleAttr("font-family: 'Source Serif'"),
				h.Class("mb-5"),
				g.Text("Players"),
			),
			h.Div(
				h.Class("mb-3 d-flex justify-content-end"),
				h.Button(
					hx.Get(fragmentCreatePlayerModal.Endpoint()),
					hx.Target(fragmentCreatePlayerModal.IDSelector()),
					hx.Trigger("click"),
					g.Attr("data-bs-toggle", "modal"),
					g.Attr("data-bs-target", fragmentCreatePlayerModal.IDSelector()),
					h.Class("btn btn-primary"),
					g.Text("Create Player"),
				),
			),
			server.playersTable(ctx),
			modalPlaceholder(fragmentCreatePlayerModal.String()),
			modalPlaceholder(fragmentRegisterPlayerModal.String()),
			modalPlaceholder(fragmentDeletePlayerModal.String()),
		),
	})
}

func (server *webServer) playersTableHandlerFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	server.playersTable(ctx).Render(w)
}

func (server *webServer) playersTable(ctx context.Context) g.Node {
	var err error
	players, err := models.Players.Query(ctx, server.db, sm.OrderBy(models.ColumnNames.Players.Username), sm.Limit(10)).All()
	if err != nil {
		return errorAlert(err)
	}
	return h.Div(
		hx.Trigger(webutils.JoinEvents(eventPlayerCreated, eventPlayerDeleted)),
		hx.Get(fragmentPlayersTable.Endpoint()),
		hx.Swap("outerHTML"),
		h.Class("d-flex justify-content-center"),
		h.Table(
			h.Class("table table-striped"),
			h.THead(
				h.Class("table-primary"),
				h.Tr(
					h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Username")),
					h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Actions")),
				),
			),
			h.TBody(
				g.Group(g.Map(players, func(player *models.Player) g.Node {
					return h.Tr(
						h.Td(g.Text(player.Username)),
						h.Td(
							h.Button(
								hx.Get(fragmentRegisterPlayerModal.Endpoint()+"/"+player.ID.String()),
								hx.Target(fragmentRegisterPlayerModal.IDSelector()),
								hx.Swap("innerHTML"),
								g.Attr("data-bs-toggle", "modal"),
								g.Attr("data-bs-target", fragmentRegisterPlayerModal.IDSelector()),
								h.Class("btn btn-sm btn-primary me-3"),
								g.Text("Register"),
							),
							h.Button(
								hx.Get(fragmentDeletePlayerModal.Endpoint()+"/"+player.ID.String()),
								hx.Target(fragmentDeletePlayerModal.IDSelector()),
								hx.Swap("innerHTML"),
								g.Attr("data-bs-toggle", "modal"),
								g.Attr("data-bs-target", fragmentDeletePlayerModal.IDSelector()),
								h.Class("btn btn-sm btn-danger"),
								g.Text("Delete"),
							),
						),
					)
				})),
			),
		),
	)
}
