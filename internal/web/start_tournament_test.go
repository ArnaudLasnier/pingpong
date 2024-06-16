package web

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/ArnaudLasnier/pingpong/internal/database/models/factory"
	"github.com/ArnaudLasnier/pingpong/internal/webutils"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func TestStartTournament(t *testing.T) {
	var err error
	ctx := context.TODO()

	// ARRANGE

	tournamentTemplate := testFactory.NewTournament(
		factory.TournamentMods.Status(models.TournamentStatusDraft),
		factory.TournamentMods.WithNewPlayers(2),
	)
	tournament, err := tournamentTemplate.Create(ctx, testWebServer.db)
	if err != nil {
		t.Fatal(err)
	}

	// ACT

	responseRecorder := httptest.NewRecorder()
	form := url.Values{}
	form.Add(string(formKeyTournamentID), tournament.ID.String())
	request := webutils.NewRequestWithForm(formActionStartTournament, form)
	// This is what we actually test:
	testWebServer.startTournament(responseRecorder, request)
	response := responseRecorder.Result()

	// ASSERT

	expectedTournamentStatus := models.TournamentStatusStarted
	expectedResponseStatusCode := http.StatusNoContent
	expectedMatchCount := 1

	tournament, err = models.FindTournament(ctx, testWebServer.db, tournament.ID)
	if err != nil {
		t.Fatal(err)
	}
	if tournament.Status != expectedTournamentStatus {
		t.Errorf("wrong tournament status: expected %s, got %s", expectedTournamentStatus, tournament.Status)
	}
	if response.StatusCode != expectedResponseStatusCode {
		t.Errorf("wrong response status code: expected %d, got %d", expectedResponseStatusCode, response.StatusCode)
	}
	matchCount, err := models.Matches.Query(
		ctx,
		testWebServer.db,
		sm.Where(
			models.MatchColumns.TournamentID.EQ(psql.Arg(tournament.ID)),
		),
	).Count()
	if err != nil {
		t.Fatal(err)
	}
	if matchCount != int64(expectedMatchCount) {
		t.Errorf("wrong match count: expected %d, got %d", expectedMatchCount, matchCount)
	}
}
