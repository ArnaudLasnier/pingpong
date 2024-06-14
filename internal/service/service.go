package service

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"slices"
	"time"

	"github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/google/uuid"
	"github.com/jonboulle/clockwork"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/dm"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

var week = time.Hour * 24 * 7

type Service struct {
	db    bob.DB
	clock clockwork.Clock
}

func NewService(db bob.DB, clock clockwork.Clock) *Service {
	return &Service{
		db:    db,
		clock: clock,
	}
}

func (s *Service) CreatePlayer(ctx context.Context, playerToCreate *models.PlayerSetter) (*models.Player, error) {
	return models.Players.Insert(ctx, s.db, playerToCreate)
}

func (s *Service) DeletePlayer(ctx context.Context, player *models.Player) error {
	return models.Players.Delete(ctx, s.db, player)
}

func (s *Service) CreateTournamentDraft(ctx context.Context, title string) (*models.Tournament, error) {
	var err error
	tournament, err := models.Tournaments.Insert(ctx, s.db, &models.TournamentSetter{
		Title:  omit.From(title),
		Status: omit.From(models.TournamentStatusDraft),
	})
	return tournament, err
}

func (s *Service) AddParticipants(ctx context.Context, tournament *models.Tournament, participants ...*models.Player) error {
	if len(participants) == 0 {
		return nil
	}
	var err error
	var values []bob.Expression
	for _, participant := range participants {
		values = append(values, psql.Arg(tournament.ID, participant.ID))
	}
	stmt := psql.Insert(
		im.Into(
			models.TableNames.TournamentParticipations,
			models.ColumnNames.TournamentParticipations.TournamentID,
			models.ColumnNames.TournamentParticipations.ParticipantID,
		),
		im.Values(values...),
	)
	_, err = stmt.Exec(ctx, s.db)
	// models.Tournaments.InsertMany()
	// models.TournamentParticipations.InsertMany()
	return err
}

func (s *Service) AddParticipants2(ctx context.Context, tournament *models.Tournament, participants ...*models.Player) error {
	return tournament.AttachPlayers(ctx, s.db, participants...)

}

func (s *Service) RemoveParticipants(ctx context.Context, tournament *models.Tournament, participants ...*models.Player) error {
	if len(participants) == 0 {
		return nil
	}
	var err error
	participantsTable := models.TableNames.TournamentParticipations
	tournamentID := models.ColumnNames.TournamentParticipations.TournamentID
	participantID := models.ColumnNames.TournamentParticipations.ParticipantID
	var whereClause dialect.Expression
	for _, participant := range participants {
		whereClause = whereClause.Or(
			psql.Quote(tournamentID).EQ(psql.Arg(tournament.ID)).And(
				psql.Quote(participantID).EQ(psql.Arg(participant.ID)),
			),
		)
	}
	stmt := psql.Delete(
		dm.From(participantsTable),
		dm.Where(whereClause),
	)
	_, err = stmt.Exec(ctx, s.db)
	return err
}

func (s *Service) RemoveParticipants2(ctx context.Context, tournament *models.Tournament, participants ...*models.Player) error {
	var err error
	var participantIDs []bob.Expression
	for _, participant := range participants {
		participantIDs = append(participantIDs, psql.Arg(participant.ID))
	}
	_, err = models.TournamentParticipations.DeleteQ(
		ctx,
		s.db,
		dm.Where(
			psql.Quote(models.ColumnNames.TournamentParticipations.TournamentID).EQ(psql.Arg(tournament.ID)).And(
				psql.Quote(models.ColumnNames.TournamentParticipations.ParticipantID).In(participantIDs...),
			),
		),
	).Exec()
	return err
}

func (s *Service) CheckTournamentDraft(ctx context.Context, tournament *models.Tournament) error {
	numberOfParticipants, err := tournament.Players(ctx, s.db).Count()
	if err != nil {
		return err
	}
	if numberOfParticipants < 2 {
		return &NotEnoughParticipantsError{Count: numberOfParticipants}
	}
	if numberOfParticipants%2 != 0 {
		return &OddNumberOfParticipantsError{Count: numberOfParticipants}
	}
	return nil
}

func (s *Service) StartTournament(ctx context.Context, tournament *models.Tournament) error {
	var err error
	err = s.CheckTournamentDraft(ctx, tournament)
	if err != nil {
		return err
	}
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	startTime := s.clock.Now()
	tournament.Update(ctx, tx, &models.TournamentSetter{
		Status:    omit.From(models.TournamentStatusStarted),
		StartedAt: omitnull.From(startTime),
	})
	participants, err := tournament.Players(ctx, tx).All()
	if err != nil {
		return err
	}
	shuffle(participants)
	matchesToInsert := s.generateMatchesToInsert(tournament, participants, startTime)
	err = tournament.InsertMatches(ctx, tx, matchesToInsert...)
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

func (s *Service) generateMatchesToInsert(tournament *models.Tournament, participants []*models.Player, startTime time.Time) []*models.MatchSetter {
	numberOfRounds := int(math.Log2(float64(len(participants))))
	matchesPerRound := make([][]*models.MatchSetter, numberOfRounds)
	matchesPerRound[0] = make([]*models.MatchSetter, len(participants)/2)
	for k := 0; k < len(participants)/2; k += 2 {
		matchesPerRound[0] = append(matchesPerRound[0], &models.MatchSetter{
			ID:             omit.From(uuid.New()),
			TournamentID:   omit.From(tournament.ID),
			ParentMatch1ID: omitnull.FromPtr(new(uuid.UUID)),
			ParentMatch2ID: omitnull.FromPtr(new(uuid.UUID)),
			DueAt:          omit.From(startTime.Add(1 * week)),
			Opponent1ID:    omitnull.From(participants[k].ID),
			Opponent2ID:    omitnull.From(participants[k+1].ID),
		})
	}
	for i := 0; i < numberOfRounds-2; i++ {
		matchesPerRound[i+1] = make([]*models.MatchSetter, 0)
		for j := 0; j < len(matchesPerRound[i])/2; j += 2 {
			matchesPerRound[i+1] = append(matchesPerRound[i+1], &models.MatchSetter{
				ID:             omit.From(uuid.New()),
				TournamentID:   omit.From(tournament.ID),
				ParentMatch1ID: omitnull.From(matchesPerRound[i][j].ID.MustGet()),
				ParentMatch2ID: omitnull.From(matchesPerRound[i][j+1].ID.MustGet()),
				DueAt:          omit.From(matchesPerRound[i][j].DueAt.MustGet().Add(1 * week)),
			})
		}
	}
	return slices.Concat(matchesPerRound...)
}

func (s *Service) EnterMatchResult(ctx context.Context, match *models.Match, score1 int, score2 int) error {
	var err error
	match.Update(ctx, s.db, &models.MatchSetter{
		Opponent1Score: omitnull.From(int32(score1)),
		Opponent2Score: omitnull.From(int32(score2)),
	})
	children, err := models.Matches.Query(
		ctx,
		s.db,
		sm.Where(
			psql.Quote(models.ColumnNames.Matches.ParentMatch1ID).EQ(psql.Arg(match.ID)).Or(
				psql.Quote(models.ColumnNames.Matches.ParentMatch2ID).EQ(psql.Arg(match.ID)),
			),
		),
	).All()
	if err != nil {
		return err
	}
	if len(children) == 0 {
		return nil
	}
	nextMatch := children[0]
	var winner *models.Player
	if match.Opponent1Score.MustGet() > match.Opponent2Score.MustGet() {
		winner, err = models.FindPlayer(ctx, s.db, match.Opponent1ID.MustGet())
		if err != nil {
			return err
		}
	} else {
		winner, err = models.FindPlayer(ctx, s.db, match.Opponent2ID.MustGet())
		if err != nil {
			return err
		}
	}
	if nextMatch.Opponent1ID.IsNull() {
		err = nextMatch.AttachOpponent1Player(ctx, s.db, winner)
		if err != nil {
			return err
		}
	} else if nextMatch.Opponent2ID.IsNull() {
		nextMatch.AttachOpponent2Player(ctx, s.db, winner)
		if err != nil {
			return err
		}
	}
	return nil
}

type OddNumberOfParticipantsError struct {
	Count int64
}

func (e *OddNumberOfParticipantsError) Error() string {
	return fmt.Sprintf("the tournament has an odd number of participants: %d participants", e.Count)
}

type NotEnoughParticipantsError struct {
	Count int64
}

func (e *NotEnoughParticipantsError) Error() string {
	return fmt.Sprintf("the tournament does not have enough participants: %d participants", e.Count)
}

// Fisher–Yates shuffle.
func shuffle[T any](s []T) {
	n := len(s)
	for i := range n - 2 {
		j := rand.IntN(n-i) + i
		s[i], s[j] = s[j], s[i]
	}
}
