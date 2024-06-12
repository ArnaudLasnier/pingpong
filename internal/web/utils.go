package web

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func DisplayNone() g.Node {
	return h.StyleAttr("display: none")
}
