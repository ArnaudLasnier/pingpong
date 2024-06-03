package main

import "net/http"

type HelloWorldHandler struct{}

func (handler *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	http.ListenAndServe(":3000", &HelloWorldHandler{})
}
