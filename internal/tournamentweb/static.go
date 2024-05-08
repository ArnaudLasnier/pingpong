package tournamentweb

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static
var staticFS embed.FS

type staticHandler struct{}

func newStaticHandler() *staticHandler {
	return &staticHandler{}
}

func (sh *staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	staticFS, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
	http.FileServerFS(staticFS).ServeHTTP(w, r)
}
