package web

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (handler *webServer) playersHandlerFunc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url := r.URL
	err := handler.playersPage(ctx, *url).Render(w)
	if err != nil {
		errorAlert(err).Render(w)
		return
	}
}

func (handler *webServer) playersPage(ctx context.Context, url url.URL) g.Node {
	var err error
	players, err := models.Players.Query(ctx, handler.db, sm.OrderBy(models.ColumnNames.Players.Username), sm.Limit(10)).All()
	if err != nil {
		return errorAlert(err)
	}
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: "Players",
		Body: h.Div(
			h.H1(g.Text("Players")),
			h.Div(
				h.Class("mb-3"),
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
			h.Div(
				h.Class("d-flex justify-content-center"),
				h.Table(
					h.Class("table w-75"),
					h.THead(
						h.Class("table-light"),
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
										g.Attr("data-bs-toggle", "modal"),
										g.Attr("data-bs-target", fragmentRegisterPlayerModal.IDSelector()),
										h.Class("btn btn-sm btn-primary me-3"),
										g.Text("Register"),
									),
									h.Button(
										hx.Get(fragmentDeletePlayerModal.Endpoint()+"/"+player.ID.String()),
										hx.Target(fragmentDeletePlayerModal.IDSelector()),
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
			),
			modalPlaceholder(fragmentCreatePlayerModal.String()),
			modalPlaceholder(fragmentRegisterPlayerModal.String()),
			modalPlaceholder(fragmentDeletePlayerModal.String()),
		),
	})
}
