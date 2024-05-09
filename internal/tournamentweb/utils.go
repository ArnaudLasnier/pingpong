package tournamentweb

import "net/http"

const (
	ContentTypeHeader = "Content-Type"
	HTMLMediaType     = "text/html; charset=utf-8"
)

func todoPanic(v any) {
	panic(v)
}

type FormField struct {
	Value   string
	IsValid bool
	Messge  string
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
