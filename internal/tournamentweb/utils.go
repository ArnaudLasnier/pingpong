package tournamentweb

import "net/http"

const (
	ContentTypeHeader = "Content-Type"
	HTMLMediaType     = "text/html; charset=utf-8"
)

const (
	CacheControlHeader = "Cache-Control"
	NoCache            = "no-cache"
)

func todoPanic(v any) {
	panic(v)
}

type FormField struct {
	Value   string
	IsValid bool
	Message string
}

type Form struct {
	IsSubmitted bool
	Fields      map[string]FormField
}

func MiddlewareHTML(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ContentTypeHeader, HTMLMediaType)
		next.ServeHTTP(w, r)
	})
}

func MiddlewareNoCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(CacheControlHeader, NoCache)
		next.ServeHTTP(w, r)
	})
}
