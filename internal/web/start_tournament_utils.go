package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/service"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
)

func writeSuccessDataInHxTriggerHeader(w http.ResponseWriter, successData webutils.SuccessData) {
	triggerData := map[string]any{
		eventTournamentStarted.String(): nil,
		eventShowSuccess.String():       successData,
	}
	triggerDataStr, _ := json.Marshal(triggerData)
	w.Header().Set(webutils.HeaderHxTrigger, string(triggerDataStr))
}

func writeErrorDataInHxTriggerHeader(w http.ResponseWriter, errData webutils.ErrorData) {
	triggerData := map[string]webutils.ErrorData{
		eventShowError.String(): errData,
	}
	triggerDataStr, _ := json.Marshal(triggerData)
	w.Header().Set(webutils.HeaderHxTrigger, string(triggerDataStr))
}

func failBecauseTournamentIDNotValid(w http.ResponseWriter, formTournamentID string) {
	writeErrorDataInHxTriggerHeader(w, webutils.ErrorData{
		Title:  "Tournament Not Started",
		Detail: fmt.Sprintf("The tournament could not be started because \"%s\" is not a valid UUID.", formTournamentID),
		Status: http.StatusBadRequest,
		Code:   "/tournament-not-started/tournament-id-not-valid",
	})
	w.WriteHeader(http.StatusBadRequest)
}

func failBecauseTournamentNotFound(w http.ResponseWriter) {
	writeErrorDataInHxTriggerHeader(w, webutils.ErrorData{
		Title:  "Tournament Not Started",
		Detail: "The tournament could not be started because it was not found in the database.",
		Status: http.StatusBadRequest,
		Code:   "/tournament-not-started/tournament-not-found",
	})
	w.WriteHeader(http.StatusBadRequest)
}

func failBecauseTournamentAlreadyStarted(w http.ResponseWriter) {
	writeErrorDataInHxTriggerHeader(w, webutils.ErrorData{
		Title:  "Tournament Not Started",
		Detail: "The tournament has already started.",
		Status: http.StatusBadRequest,
		Code:   "/tournament-not-started/already-started",
	})
	w.WriteHeader(http.StatusBadRequest)
}

func failBecauseNotEnoughPlayers(w http.ResponseWriter) {
	writeErrorDataInHxTriggerHeader(w, webutils.ErrorData{
		Title:  "Tournament Not Started",
		Detail: "The tournament requires at least two registered players.",
		Status: http.StatusBadRequest,
		Code:   "/tournament-not-started/not-enough-players",
	})
	w.WriteHeader(http.StatusBadRequest)
}

func failBecauseOddNumberOfPlayers(w http.ResponseWriter, oddNumberErr *service.OddNumberOfParticipantsError) {
	writeErrorDataInHxTriggerHeader(w, webutils.ErrorData{
		Title:  "Tournament Not Started",
		Detail: fmt.Sprintf("The tournament has %d registered players but requires an even number of players.", oddNumberErr.Count),
		Status: http.StatusBadRequest,
		Code:   "/tournament-not-started/not-enough-players",
	})
	w.WriteHeader(http.StatusBadRequest)
}

func failBecauseUnknownError(w http.ResponseWriter, err error) {
	writeErrorDataInHxTriggerHeader(w, webutils.ErrorData{
		Title:  "Tournament Not Started",
		Detail: err.Error(),
		Status: http.StatusBadRequest,
		Code:   "/tournament-not-started/unknown-error",
	})
	w.WriteHeader(http.StatusBadRequest)
}
