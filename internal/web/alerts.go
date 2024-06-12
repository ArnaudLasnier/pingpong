package web

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func successAlert() g.Node {
	return h.Div(
		h.Class("alert alert-success"),
		h.Role("alert"),
		g.Text("Success!"),
	)
}

func errorAlert(err error) g.Node {
	return h.Div(
		h.Class("alert alert-danger"),
		h.Role("alert"),
		h.H5(
			h.Class("alert-heading"),
			g.Text("Error"),
		),
		h.P(g.Text(err.Error())),
	)
}
