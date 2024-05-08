package tournamentweb

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/tournamentdatabase/models"
	"github.com/aarondl/opt/omit"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (handler *handler) homePage(ctx context.Context) g.Node {
	var err error
	players, err := models.Players.Query(ctx, handler.db, sm.OrderBy(models.ColumnNames.Players.LastName), sm.Limit(10)).All()
	if err != nil {
		todoPanic(err)
	}
	return PageLayout(PageLayoutProps{
		Title: "Ping Pong",
		Body: []g.Node{
			h.Div(
				h.Div(
					h.Class("d-flex justify-content-center mb-3 p-3 text-bg-primary fs-5"),
					g.Text("Hello world! Welcome to BackCon!"),
				),
				h.Div(
					h.Class("p-3"),
					h.H1(g.Text("Ping Pong - A tournament generator")),
				),
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
								h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("ID")),
								h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("Last Name")),
								h.Th(g.Attr("scope", "col"), h.Class("col-1"), g.Text("First Name")),
							),
						),
						h.TBody(
							g.Group(g.Map(players, func(player *models.Player) g.Node {
								return h.Tr(
									h.Td(g.Text(player.ID.String())),
									h.Td(g.Text(player.LastName)),
									h.Td(g.Text(player.FirstName)),
								)
							})),
						),
					),
				),
			),
		},
	})
}

func (handler *handler) handleGetHomePage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := handler.homePage(ctx).Render(w)
	if err != nil {
		todoPanic(err)
	}
}

func (handler *handler) createPlayerModal() g.Node {
	return h.Div(
		h.Class("modal-dialog modal-dialog-centered"),
		h.Div(
			h.Class("modal-content"),
			h.Div(
				h.Class("modal-header d-flex justify-content-between"),
				h.H5(h.Class("modal-title"), g.Text("Create Player")),
				h.Button(h.Type("button"), h.Class("btn btn-secondary"), g.Attr("data-bs-dismiss", "modal"), g.Text("Close")),
			),
			h.Div(
				h.Class("modal-body"),
				handler.createPlayerForm(),
			),
		),
	)
}

func (handler *handler) createPlayerForm() g.Node {
	return h.FormEl(
		hx.Post("/add-player-modal"),
		h.Div(
			h.Class("mb-3"),
			h.Label(h.For("playerFirstName"), h.Class("form-label"), g.Text("First Name")),
			h.Input(
				h.ID("playerFirstName"),
				h.Name("firstName"),
				h.Type("text"),
				h.Required(),
				h.Pattern("[A-Za-z0-9]{1,50}"),
				h.Class("form-control"),
			),
		),
		h.Div(
			h.Class("mb-4"),
			h.Label(h.For("playerLastName"), h.Class("form-label"), g.Text("Last Name")),
			h.Input(
				h.ID("playerLastName"),
				h.Name("lastName"),
				h.Type("text"),
				h.Required(),
				h.Pattern("[A-Za-z0-9]{1,50}"),
				h.Class("form-control"),
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
				h.Class("form-control"),
			),
		),
		h.Button(h.Type("submit"), h.Class("btn btn-primary"), g.Text("Submit")),
	)
}

func (handler *handler) handleGetCreatePlayerModal(w http.ResponseWriter, r *http.Request) {
	err := handler.createPlayerModal().Render(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *handler) handlePostCreatePlayerModal(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	_, err = handler.tournamentService.CreatePlayer(ctx, &models.PlayerSetter{
		FirstName: omit.From(r.PostFormValue("firstName")),
		LastName:  omit.From(r.PostFormValue("lastName")),
		Email:     omit.From(r.PostFormValue("email")),
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		fmt.Println("pgErr =", pgErr)
		if pgErr.Code == pgerrcode.UniqueViolation {
		}
	}
	if err != nil {
		h.Div(g.Text(err.Error())).Render(w)
		return
	}
	h.Div(g.Text("Success!")).Render(w)
}
