package web

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Modal(title string, body g.Node) g.Node {
	return h.Div(
		h.Class("modal-dialog modal-dialog-centered"),
		h.Div(
			h.Class("modal-content"),
			h.Div(
				h.Class("modal-header d-flex justify-content-between"),
				h.H5(h.Class("modal-title"), g.Text(title)),
				h.Button(h.Type("button"), h.Class("btn-close"), g.Attr("data-bs-dismiss", "modal")),
			),
			h.Div(
				h.Class("modal-body"),
				body,
			),
		),
	)
}

func ModalPlaceholder(id string) g.Node {
	return h.Div(
		h.ID(id),
		h.Class("modal modal-blur fade"),
		h.StyleAttr("display: none"),
		g.Attr("aria-hidden", "false"),
		g.Attr("tab-index", "-1"),
		h.Div(
			h.Class("modal-dialog modal-dialog-centered"),
			h.Role("document"),
			h.Div(h.Class("modal-content")),
		),
	)
}
