package web

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

// Technique taken from: https://blog.benoitblanchon.fr/django-htmx-toasts/
func toastPlaceholder() g.Node {
	return h.Div(
		h.ID(fragmentToastContainer.String()),
		h.Class("toast-container position-fixed top-0 end-0 p-3"),
		h.Div(
			h.ID(fragmentToast.String()),
			// Show with `(new bootstrap.Toast(el)).show()`.
			h.Class("toast"),
			h.Role("alert"),
			h.Aria("live", "assertive"),
			h.Aria("atomic", "true"),
			h.Div(
				h.Class("toast-header"),
				h.ID(fragmentToastHeader.String()),
				h.Strong(
					h.Class("me-auto"),
					h.ID(fragmentToastHeaderTitle.String()),
					// Set toast header with the DOM API `el.innerText = ...`.
				),
				h.Button(
					h.Type("button"),
					h.Class("btn-close"),
					h.DataAttr("bs-dismiss", "toast"),
					h.Aria("label", "Close"),
				),
			),
			h.Div(
				h.Class("toast-body"),
				h.ID(fragmentToastBody.String()),
				// Set toast body with the DOM API `el.innerText = ...`.
			),
		),
	)
}
