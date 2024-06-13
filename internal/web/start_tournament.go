package web

import (
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/webutils"
	"github.com/mavolin/go-htmx"
)

func (server *webServer) startTournament(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusInternalServerError
	errData := webutils.ErrorData{
		Title:  "Tournament Not Started",
		Detail: "The tournament could not be started",
		Status: statusCode,
		Code:   "tournament-not-started",
	}
	htmx.Trigger(r, eventShowError.String(), errData)
	w.WriteHeader(statusCode)
}
