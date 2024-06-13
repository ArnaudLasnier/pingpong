package webutils

import (
	"errors"
	"net/http"
	stdpath "path"
)

type FormAction struct {
	method string
	path   string
}

func NewFormAction(method string, path string) FormAction {
	if method != http.MethodGet &&
		method != http.MethodPost &&
		method != http.MethodPut &&
		method != http.MethodPatch &&
		method != http.MethodDelete {
		panic(errors.New("new form action: method not supported"))
	}
	if !stdpath.IsAbs(path) {
		panic(errors.New("new form action: path is not absolute"))
	}
	return FormAction{
		method: method,
		path:   path,
	}
}

func (action FormAction) String() string {
	return action.Endpoint()
}

func (action FormAction) Endpoint() string {
	return action.method + " " + action.path
}

func (action FormAction) Path() string {
	return action.path
}
