package tournamentweb

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

type PageLayoutProps struct {
	Title       string
	Description string
	Language    string
	Body        []g.Node
}

// A layout that makes use of HTMX, Bootstrap and Source fonts.
func PageLayout(props PageLayoutProps) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:       props.Title,
		Description: props.Description,
		Language:    props.Language,
		Head: []g.Node{
			bootstrapCSSLink(),
			sourceFontsStyleEl(),
			customBootstrapVariablesStyleEl(),
			bootstrapScript(),
			htmxScript(),
		},
		Body: props.Body,
	})
}

func bootstrapCSSLink() g.Node {
	return h.Link(
		h.Href("/static/bootstrap/5.3.3/css/bootstrap.min.css"),
		h.Rel("stylesheet"),
		g.Attr("integrity", "sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH"),
		g.Attr("crossorigin", "anonymous"),
	)
}

func customBootstrapVariablesStyleEl() g.Node {
	return h.StyleEl(
		g.Raw(
			`
			body {
				--bs-body-font-family: 'Source Sans';
			}
			`,
		),
	)
}

func bootstrapScript() g.Node {
	return h.Script(
		h.Src("/static/bootstrap/5.3.3/js/bootstrap.bundle.min.js"),
		g.Attr("integrity", "sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz"),
		g.Attr("crossorigin", "anonymous"),
	)
}

func sourceFontsStyleEl() g.Node {
	return h.StyleEl(
		g.Raw(
			`
			@font-face {
				font-family: 'Source Serif';
				font-style: normal;
				font-weight: normal;
				src: url('/static/source-serif/4.005/SourceSerif4-Regular.otf.woff2') format('woff2');
			}

			@font-face {
				font-family: 'Source Serif';
				font-style: normal;
				font-weight: bold;
				src: url('/static/source-serif/4.005/SourceSerif4-Bold.otf.woff2') format('woff2');
			}

			@font-face {
				font-family: 'Source Serif';
				font-style: italic;
				font-weight: normal;
				src: url('/static/source-serif/4.005/SourceSerif4-It.otf.woff2') format('woff2');
			}

			@font-face {
				font-family: 'Source Serif';
				font-style: italic;
				font-weight: bold;
				src: url('/static/source-serif/4.005/SourceSerif4-BoltIt.otf.woff2') format('woff2');
			}

			@font-face {
				font-family: 'Source Sans';
				font-style: normal;
				font-weight: normal;
				src: url('/static/source-sans/3.052/SourceSans3-Regular.otf.woff2') format('woff2');
			}

			@font-face {
				font-family: 'Source Sans';
				font-style: normal;
				font-weight: bold;
				src: url('/static/source-sans/3.052/SourceSans3-Bold.otf.woff2') format('woff2');
			}

			@font-face {
				font-family: 'Source Sans';
				font-style: italic;
				font-weight: normal;
				src: url('/static/source-sans/3.052/SourceSans3-It.otf.woff2') format('woff2');
			}

			@font-face {
				font-family: 'Source Sans';
				font-style: italic;
				font-weight: bold;
				src: url('/static/source-sans/3.052/SourceSans3-BoltIt.otf.woff2') format('woff2');
			}

			@font-face {
				font-family: 'Source Code Pro';
				font-style: normal;
				font-weight: normal;
				src: url('/static/source-code-pro/2.040/SourceCodePro-Regular.otf.woff2') format('woff2');
				size-adjust: 90%;
			}

			@font-face {
				font-family: 'Source Code Pro';
				font-style: normal;
				font-weight: bold;
				src: url('/static/source-code-pro/2.040/SourceCodePro-Bold.otf.woff2') format('woff2');
				size-adjust: 90%;
			}

			@font-face {
				font-family: 'Source Code Pro';
				font-style: italic;
				font-weight: normal;
				src: url('/static/source-code-pro/2.040/SourceCodePro-It.otf.woff2') format('woff2');
				size-adjust: 90%;
			}

			@font-face {
				font-family: 'Source Code Pro';
				font-style: italic;
				font-weight: bold;
				src: url('/static/source-code-pro/2.040/SourceCodePro-BoltIt.otf.woff2') format('woff2');
				size-adjust: 90%;
			}
			`,
		),
	)
}

func htmxScript() g.Node {
	return h.Script(
		h.Src("/static/htmx/1.9.9/htmx.min.js"),
	)
}
