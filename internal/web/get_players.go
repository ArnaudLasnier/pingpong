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

func (handler *handler) players(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url := r.URL
	err := handler.playersPage(ctx, *url).Render(w)
	if err != nil {
		ErrorAlert(err).Render(w)
		return
	}
}

func (handler *handler) playersPage(ctx context.Context, url url.URL) g.Node {
	var err error
	players, err := models.Players.Query(ctx, handler.db, sm.OrderBy(models.ColumnNames.Players.LastName), sm.Limit(10)).All()
	if err != nil {
		return ErrorAlert(err)
	}
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: "Players",
		Body: h.Div(
			h.H1(g.Text("Players")),
			h.Div(
				h.Class("mb-3"),
				h.Button(
					hx.Get("/create-player-modal"),
					hx.Target("#create-player-modal"),
					hx.Trigger("click"),
					g.Attr("data-bs-toggle", "modal"),
					g.Attr("data-bs-target", "#create-player-modal"),
					h.Class("btn btn-primary"),
					g.Text("Create Player"),
				),
			),
			ModalPlaceholder("create-player-modal"),
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
