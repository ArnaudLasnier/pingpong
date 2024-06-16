package webutils

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func NewRequestWithForm(formAction FormAction, form url.Values) *http.Request {
	request := httptest.NewRequest(
		formAction.Method(),
		formAction.Path(),
		strings.NewReader(form.Encode()),
	)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return request
}
