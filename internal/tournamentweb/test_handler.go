package tournamentweb

import (
	"net/http"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
)

func test(w http.ResponseWriter, r *http.Request) {
	pageLayout(pageLayoutProps{
		URL:   *r.URL,
		Title: "Test",
		Body: h.Div(
			h.Button(
				hx.Get("/test/button"),
				hx.Target("#test-button-result"),
				g.Text("Click Me!"),
			),
			h.Div(
				h.ID("test-button-result"),
			),
		),
	}).Render(w)
}

func testButtonResult(w http.ResponseWriter, r *http.Request) {
	h.Div(
		g.Text("Some result!"),
	).Render(w)
}
