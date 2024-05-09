package tournamentweb

import (
	"net/http"
	"net/url"

	g "github.com/maragudk/gomponents"
)

func (handler *handler) tournamentsPage(url url.URL) g.Node {
	return pageLayout(pageLayoutProps{
		URL:   url,
		Title: "Tournaments",
		Body:  g.Text(""),
	})
}

func (handler *handler) handleGetTournamentsPage(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	err := handler.tournamentsPage(*url).Render(w)
	if err != nil {
		todoPanic(err)
	}
}
