package web

import (
	"errors"
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/ArnaudLasnier/pingpong/internal/service"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
	"github.com/google/uuid"
)

func (server *webServer) startTournament(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	formTournamentID := r.PostFormValue(formKeyTournamentID.String())
	tournamentID, err := uuid.Parse(formTournamentID)

	if err != nil {
		failBecauseTournamentIDNotValid(w, formTournamentID)
		return
	}

	tournament, err := models.FindTournament(ctx, server.db, tournamentID)

	if err != nil {
		failBecauseTournamentNotFound(w)
		return
	}

	err = server.service.StartTournament(ctx, tournament)

	var notEnoughErr *service.NotEnoughParticipantsError
	var oddNumberErr *service.OddNumberOfParticipantsError

	switch {
	// 1. Is the tournament already started?
	case errors.Is(err, service.ErrTournamentAlreadyStarted):
		failBecauseTournamentAlreadyStarted(w)
		return

	// 2. Are there enough players?
	case errors.As(err, &notEnoughErr):
		failBecauseNotEnoughPlayers(w)
		return

	// 3. Do we have an even number of players?
	case errors.As(err, &oddNumberErr):
		failBecauseOddNumberOfPlayers(w, oddNumberErr)
		return

	// 4. Is there still an error that we don't handle specially?
	case err != nil:
		failBecauseUnknownError(w, err)
		return
	}

	writeSuccessDataInHxTriggerHeader(w, webutils.SuccessData{
		Title:  "Tournament Started",
		Detail: "The tournament has been successfully started.",
	})
	w.WriteHeader(http.StatusNoContent)
}
