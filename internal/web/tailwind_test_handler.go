package web

import (
	"net/http"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
)

func handleTailwindTest(w http.ResponseWriter, r *http.Request) {
	pageLayout(pageLayoutProps{
		URL:   *r.URL,
		Title: "Test",
		Body: h.Div(
			h.Button(
				hx.Get("/tailwind-test/button"),
				hx.Target("#test-button-result"),
				h.Class("text-white bg-blue-500 rounded-sm"),
				g.Text("Click Me!"),
			),
			h.Div(
				h.ID("test-button-result"),
			),
		),
	}).Render(w)
}

func tailwindTestButtonResult(w http.ResponseWriter, r *http.Request) {
	h.Div(
		g.Text("Some result!"),
	).Render(w)
}
