package web

import "net/http"

const (
	contentTypeHeader = "Content-Type"
	htmlMediaType     = "text/html; charset=utf-8"
)

func htmlContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeHeader, htmlMediaType)
		next.ServeHTTP(w, r)
	})
}
