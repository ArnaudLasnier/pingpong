package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/ArnaudLasnier/pingpong/internal/service"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
	"github.com/google/uuid"
	"github.com/mavolin/go-htmx"
)

func (server *webServer) startTournament(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()
	formTournamentID := r.PostFormValue(formKeyTournamentID.String())
	tournamentID, err := uuid.Parse(formTournamentID)

	if err != nil {
		errData := webutils.ErrorData{
			Title:  "Tournament Not Started",
			Detail: fmt.Sprintf("The tournament could not be started because \"%s\" is not a valid UUID.", formTournamentID),
			Status: http.StatusBadRequest,
			Code:   "/tournament-not-started/tournament-id-not-valid",
		}
		htmx.Trigger(r, eventShowError.String(), errData)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tournament, err := models.FindTournament(ctx, server.db, tournamentID)
	if err != nil {
		errData := webutils.ErrorData{
			Title:  "Tournament Not Started",
			Detail: "The tournament could not be started because it was not found in the database.",
			Status: http.StatusBadRequest,
			Code:   "/tournament-not-started/tournament-not-found",
		}
		htmx.Trigger(r, eventShowError.String(), errData)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = server.service.StartTournament(ctx, tournament)
	var notEnoughErr *service.NotEnoughParticipantsError
	var oddNumberErr *service.OddNumberOfParticipantsError
	switch {
	case errors.Is(err, service.ErrTournamentAlreadyStarted):
		errData := webutils.ErrorData{
			Title:  "Tournament Not Started",
			Detail: "The tournament has already started.",
			Status: http.StatusBadRequest,
			Code:   "/tournament-not-started/already-started",
		}
		htmx.Trigger(r, eventShowError.String(), errData)
		w.WriteHeader(http.StatusBadRequest)
		return
	case errors.As(err, &notEnoughErr):
		errData := webutils.ErrorData{
			Title:  "Tournament Not Started",
			Detail: fmt.Sprintf("The tournament has only %d registered players but requires at least two.", notEnoughErr.Count),
			Status: http.StatusBadRequest,
			Code:   "/tournament-not-started/not-enough-players",
		}
		htmx.Trigger(r, eventShowError.String(), errData)
		w.WriteHeader(http.StatusBadRequest)
		return
	case errors.As(err, &oddNumberErr):
		errData := webutils.ErrorData{
			Title:  "Tournament Not Started",
			Detail: fmt.Sprintf("The tournament has only %d registered players but requires an even number of players.", oddNumberErr.Count),
			Status: http.StatusBadRequest,
			Code:   "/tournament-not-started/not-enough-players",
		}
		htmx.Trigger(r, eventShowError.String(), errData)
		w.WriteHeader(http.StatusBadRequest)
		return
	case err != nil:
		errData := webutils.ErrorData{
			Title:  "Tournament Not Started",
			Detail: err.Error(),
			Status: http.StatusBadRequest,
			Code:   "/tournament-not-started/unknown-error",
		}
		htmx.Trigger(r, eventShowError.String(), errData)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
