package web

import "net/http"

const (
	ContentTypeHeader = "Content-Type"
	HTMLMediaType     = "text/html; charset=utf-8"
)

const (
	CacheControlHeader = "Cache-Control"
	NoCache            = "no-cache"
)

func HTMLContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ContentTypeHeader, HTMLMediaType)
		next.ServeHTTP(w, r)
	})
}

func NoCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(CacheControlHeader, NoCache)
		next.ServeHTTP(w, r)
	})
}
