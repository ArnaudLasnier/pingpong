// Code generated by the local DBGEN tool. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/google/uuid"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/clause"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
	"github.com/stephenafamo/bob/expr"
	"github.com/stephenafamo/bob/mods"
	"github.com/stephenafamo/scan"
)

// Player is an object representing the database table.
type Player struct {
	ID       uuid.UUID `db:"id,pk" `
	Username string    `db:"username" `

	R playerR `db:"-" `
}

// PlayerSlice is an alias for a slice of pointers to Player.
// This should almost always be used instead of []*Player.
type PlayerSlice []*Player

// Players contains methods to work with the player table
var Players = psql.NewTablex[*Player, PlayerSlice, *PlayerSetter]("", "player")

// PlayersQuery is a query on the player table
type PlayersQuery = *psql.ViewQuery[*Player, PlayerSlice]

// PlayersStmt is a prepared statment on player
type PlayersStmt = bob.QueryStmt[*Player, PlayerSlice]

// playerR is where relationships are stored.
type playerR struct {
	Opponent1Matches MatchSlice      // match.match_opponent_1_id_fkey
	Opponent2Matches MatchSlice      // match.match_opponent_2_id_fkey
	Tournaments      TournamentSlice // tournament_participation.tournament_participation_participant_id_fkeytournament_participation.tournament_participation_tournament_id_fkey
}

// PlayerSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type PlayerSetter struct {
	ID       omit.Val[uuid.UUID] `db:"id,pk"`
	Username omit.Val[string]    `db:"username"`
}

func (s PlayerSetter) SetColumns() []string {
	vals := make([]string, 0, 2)
	if !s.ID.IsUnset() {
		vals = append(vals, "id")
	}

	if !s.Username.IsUnset() {
		vals = append(vals, "username")
	}

	return vals
}

func (s PlayerSetter) Overwrite(t *Player) {
	if !s.ID.IsUnset() {
		t.ID, _ = s.ID.Get()
	}
	if !s.Username.IsUnset() {
		t.Username, _ = s.Username.Get()
	}
}

func (s PlayerSetter) InsertMod() bob.Mod[*dialect.InsertQuery] {
	vals := make([]bob.Expression, 2)
	if s.ID.IsUnset() {
		vals[0] = psql.Raw("DEFAULT")
	} else {
		vals[0] = psql.Arg(s.ID)
	}

	if s.Username.IsUnset() {
		vals[1] = psql.Raw("DEFAULT")
	} else {
		vals[1] = psql.Arg(s.Username)
	}

	return im.Values(vals...)
}

func (s PlayerSetter) Apply(q *dialect.UpdateQuery) {
	um.Set(s.Expressions()...).Apply(q)
}

func (s PlayerSetter) Expressions(prefix ...string) []bob.Expression {
	exprs := make([]bob.Expression, 0, 2)

	if !s.ID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "id")...),
			psql.Arg(s.ID),
		}})
	}

	if !s.Username.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "username")...),
			psql.Arg(s.Username),
		}})
	}

	return exprs
}

type playerColumnNames struct {
	ID       string
	Username string
}

type playerRelationshipJoins[Q dialect.Joinable] struct {
	Opponent1Matches bob.Mod[Q]
	Opponent2Matches bob.Mod[Q]
	Tournaments      bob.Mod[Q]
}

func buildPlayerRelationshipJoins[Q dialect.Joinable](ctx context.Context, typ string) playerRelationshipJoins[Q] {
	return playerRelationshipJoins[Q]{
		Opponent1Matches: playersJoinOpponent1Matches[Q](ctx, typ),
		Opponent2Matches: playersJoinOpponent2Matches[Q](ctx, typ),
		Tournaments:      playersJoinTournaments[Q](ctx, typ),
	}
}

func playersJoin[Q dialect.Joinable](ctx context.Context) joinSet[playerRelationshipJoins[Q]] {
	return joinSet[playerRelationshipJoins[Q]]{
		InnerJoin: buildPlayerRelationshipJoins[Q](ctx, clause.InnerJoin),
		LeftJoin:  buildPlayerRelationshipJoins[Q](ctx, clause.LeftJoin),
		RightJoin: buildPlayerRelationshipJoins[Q](ctx, clause.RightJoin),
	}
}

var PlayerColumns = struct {
	ID       psql.Expression
	Username psql.Expression
}{
	ID:       psql.Quote("player", "id"),
	Username: psql.Quote("player", "username"),
}

type playerWhere[Q psql.Filterable] struct {
	ID       psql.WhereMod[Q, uuid.UUID]
	Username psql.WhereMod[Q, string]
}

func PlayerWhere[Q psql.Filterable]() playerWhere[Q] {
	return playerWhere[Q]{
		ID:       psql.Where[Q, uuid.UUID](PlayerColumns.ID),
		Username: psql.Where[Q, string](PlayerColumns.Username),
	}
}

// FindPlayer retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindPlayer(ctx context.Context, exec bob.Executor, IDPK uuid.UUID, cols ...string) (*Player, error) {
	if len(cols) == 0 {
		return Players.Query(
			ctx, exec,
			SelectWhere.Players.ID.EQ(IDPK),
		).One()
	}

	return Players.Query(
		ctx, exec,
		SelectWhere.Players.ID.EQ(IDPK),
		sm.Columns(Players.Columns().Only(cols...)),
	).One()
}

// PlayerExists checks the presence of a single record by primary key
func PlayerExists(ctx context.Context, exec bob.Executor, IDPK uuid.UUID) (bool, error) {
	return Players.Query(
		ctx, exec,
		SelectWhere.Players.ID.EQ(IDPK),
	).Exists()
}

// PrimaryKeyVals returns the primary key values of the Player
func (o *Player) PrimaryKeyVals() bob.Expression {
	return psql.Arg(o.ID)
}

// Update uses an executor to update the Player
func (o *Player) Update(ctx context.Context, exec bob.Executor, s *PlayerSetter) error {
	return Players.Update(ctx, exec, s, o)
}

// Delete deletes a single Player record with an executor
func (o *Player) Delete(ctx context.Context, exec bob.Executor) error {
	return Players.Delete(ctx, exec, o)
}

// Reload refreshes the Player using the executor
func (o *Player) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := Players.Query(
		ctx, exec,
		SelectWhere.Players.ID.EQ(o.ID),
	).One()
	if err != nil {
		return err
	}
	o2.R = o.R
	*o = *o2

	return nil
}

func (o PlayerSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals PlayerSetter) error {
	return Players.Update(ctx, exec, &vals, o...)
}

func (o PlayerSlice) DeleteAll(ctx context.Context, exec bob.Executor) error {
	return Players.Delete(ctx, exec, o...)
}

func (o PlayerSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	var mods []bob.Mod[*dialect.SelectQuery]

	IDPK := make([]uuid.UUID, len(o))

	for i, o := range o {
		IDPK[i] = o.ID
	}

	mods = append(mods,
		SelectWhere.Players.ID.In(IDPK...),
	)

	o2, err := Players.Query(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, old := range o {
		for _, new := range o2 {
			if new.ID != old.ID {
				continue
			}
			new.R = old.R
			*old = *new
			break
		}
	}

	return nil
}

func playersJoinOpponent1Matches[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, Matches.NameAs(ctx)).On(
			MatchColumns.Opponent1ID.EQ(PlayerColumns.ID),
		),
	}
}

func playersJoinOpponent2Matches[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, Matches.NameAs(ctx)).On(
			MatchColumns.Opponent2ID.EQ(PlayerColumns.ID),
		),
	}
}

func playersJoinTournaments[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, TournamentParticipations.NameAs(ctx)).On(
			TournamentParticipationColumns.ParticipantID.EQ(PlayerColumns.ID),
		),
		dialect.Join[Q](typ, Tournaments.NameAs(ctx)).On(
			TournamentColumns.ID.EQ(TournamentParticipationColumns.TournamentID),
		),
	}
}

// Opponent1Matches starts a query for related objects on match
func (o *Player) Opponent1Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) MatchesQuery {
	return Matches.Query(ctx, exec, append(mods,
		sm.Where(MatchColumns.Opponent1ID.EQ(psql.Arg(o.ID))),
	)...)
}

func (os PlayerSlice) Opponent1Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) MatchesQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = psql.ArgGroup(o.ID)
	}

	return Matches.Query(ctx, exec, append(mods,
		sm.Where(psql.Group(MatchColumns.Opponent1ID).In(PKArgs...)),
	)...)
}

// Opponent2Matches starts a query for related objects on match
func (o *Player) Opponent2Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) MatchesQuery {
	return Matches.Query(ctx, exec, append(mods,
		sm.Where(MatchColumns.Opponent2ID.EQ(psql.Arg(o.ID))),
	)...)
}

func (os PlayerSlice) Opponent2Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) MatchesQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = psql.ArgGroup(o.ID)
	}

	return Matches.Query(ctx, exec, append(mods,
		sm.Where(psql.Group(MatchColumns.Opponent2ID).In(PKArgs...)),
	)...)
}

// Tournaments starts a query for related objects on tournament
func (o *Player) Tournaments(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) TournamentsQuery {
	return Tournaments.Query(ctx, exec, append(mods,
		sm.InnerJoin(TournamentParticipations.NameAs(ctx)).On(
			TournamentColumns.ID.EQ(TournamentParticipationColumns.TournamentID)),
		sm.Where(TournamentParticipationColumns.ParticipantID.EQ(psql.Arg(o.ID))),
	)...)
}

func (os PlayerSlice) Tournaments(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) TournamentsQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = psql.ArgGroup(o.ID)
	}

	return Tournaments.Query(ctx, exec, append(mods,
		sm.InnerJoin(TournamentParticipations.NameAs(ctx)).On(
			TournamentColumns.ID.EQ(TournamentParticipationColumns.TournamentID),
		),
		sm.Where(psql.Group(TournamentParticipationColumns.ParticipantID).In(PKArgs...)),
	)...)
}

func (o *Player) Preload(name string, retrieved any) error {
	if o == nil {
		return nil
	}

	switch name {
	case "Opponent1Matches":
		rels, ok := retrieved.(MatchSlice)
		if !ok {
			return fmt.Errorf("player cannot load %T as %q", retrieved, name)
		}

		o.R.Opponent1Matches = rels

		for _, rel := range rels {
			if rel != nil {
				rel.R.Opponent1Player = o
			}
		}
		return nil
	case "Opponent2Matches":
		rels, ok := retrieved.(MatchSlice)
		if !ok {
			return fmt.Errorf("player cannot load %T as %q", retrieved, name)
		}

		o.R.Opponent2Matches = rels

		for _, rel := range rels {
			if rel != nil {
				rel.R.Opponent2Player = o
			}
		}
		return nil
	case "Tournaments":
		rels, ok := retrieved.(TournamentSlice)
		if !ok {
			return fmt.Errorf("player cannot load %T as %q", retrieved, name)
		}

		o.R.Tournaments = rels

		for _, rel := range rels {
			if rel != nil {
				rel.R.Players = PlayerSlice{o}
			}
		}
		return nil
	default:
		return fmt.Errorf("player has no relationship %q", name)
	}
}

func ThenLoadPlayerOpponent1Matches(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadPlayerOpponent1Matches(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load PlayerOpponent1Matches", retrieved)
		}

		err := loader.LoadPlayerOpponent1Matches(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadPlayerOpponent1Matches loads the player's Opponent1Matches into the .R struct
func (o *Player) LoadPlayerOpponent1Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Opponent1Matches = nil

	related, err := o.Opponent1Matches(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, rel := range related {
		rel.R.Opponent1Player = o
	}

	o.R.Opponent1Matches = related
	return nil
}

// LoadPlayerOpponent1Matches loads the player's Opponent1Matches into the .R struct
func (os PlayerSlice) LoadPlayerOpponent1Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	matches, err := os.Opponent1Matches(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, o := range os {
		o.R.Opponent1Matches = nil
	}

	for _, o := range os {
		for _, rel := range matches {
			if o.ID != rel.Opponent1ID.GetOrZero() {
				continue
			}

			rel.R.Opponent1Player = o

			o.R.Opponent1Matches = append(o.R.Opponent1Matches, rel)
		}
	}

	return nil
}

func ThenLoadPlayerOpponent2Matches(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadPlayerOpponent2Matches(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load PlayerOpponent2Matches", retrieved)
		}

		err := loader.LoadPlayerOpponent2Matches(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadPlayerOpponent2Matches loads the player's Opponent2Matches into the .R struct
func (o *Player) LoadPlayerOpponent2Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Opponent2Matches = nil

	related, err := o.Opponent2Matches(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, rel := range related {
		rel.R.Opponent2Player = o
	}

	o.R.Opponent2Matches = related
	return nil
}

// LoadPlayerOpponent2Matches loads the player's Opponent2Matches into the .R struct
func (os PlayerSlice) LoadPlayerOpponent2Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	matches, err := os.Opponent2Matches(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, o := range os {
		o.R.Opponent2Matches = nil
	}

	for _, o := range os {
		for _, rel := range matches {
			if o.ID != rel.Opponent2ID.GetOrZero() {
				continue
			}

			rel.R.Opponent2Player = o

			o.R.Opponent2Matches = append(o.R.Opponent2Matches, rel)
		}
	}

	return nil
}

func ThenLoadPlayerTournaments(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadPlayerTournaments(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load PlayerTournaments", retrieved)
		}

		err := loader.LoadPlayerTournaments(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadPlayerTournaments loads the player's Tournaments into the .R struct
func (o *Player) LoadPlayerTournaments(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Tournaments = nil

	related, err := o.Tournaments(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, rel := range related {
		rel.R.Players = PlayerSlice{o}
	}

	o.R.Tournaments = related
	return nil
}

// LoadPlayerTournaments loads the player's Tournaments into the .R struct
func (os PlayerSlice) LoadPlayerTournaments(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	// since we are changing the columns, we need to check if the original columns were set or add the defaults
	sq := dialect.SelectQuery{}
	for _, mod := range mods {
		mod.Apply(&sq)
	}

	if len(sq.SelectList.Columns) == 0 {
		mods = append(mods, sm.Columns(Tournaments.Columns()))
	}

	q := os.Tournaments(ctx, exec, append(
		mods,
		sm.Columns(TournamentParticipationColumns.ParticipantID.As("related_player.ID")),
	)...)

	IDSlice := []uuid.UUID{}

	mapper := scan.Mod(scan.StructMapper[*Tournament](), func(ctx context.Context, cols []string) (scan.BeforeFunc, func(any, any) error) {
		return func(row *scan.Row) (any, error) {
				IDSlice = append(IDSlice, *new(uuid.UUID))
				row.ScheduleScan("related_player.ID", &IDSlice[len(IDSlice)-1])

				return nil, nil
			},
			func(any, any) error {
				return nil
			}
	})

	tournaments, err := bob.Allx[*Tournament, TournamentSlice](ctx, exec, q, mapper)
	if err != nil {
		return err
	}

	for _, o := range os {
		o.R.Tournaments = nil
	}

	for _, o := range os {
		for i, rel := range tournaments {
			if o.ID != IDSlice[i] {
				continue
			}

			rel.R.Players = append(rel.R.Players, o)

			o.R.Tournaments = append(o.R.Tournaments, rel)
		}
	}

	return nil
}

func insertPlayerOpponent1Matches0(ctx context.Context, exec bob.Executor, matches1 []*MatchSetter, player0 *Player) (MatchSlice, error) {
	for i := range matches1 {
		matches1[i].Opponent1ID = omitnull.From(player0.ID)
	}

	ret, err := Matches.InsertMany(ctx, exec, matches1...)
	if err != nil {
		return ret, fmt.Errorf("insertPlayerOpponent1Matches0: %w", err)
	}

	return ret, nil
}

func attachPlayerOpponent1Matches0(ctx context.Context, exec bob.Executor, count int, matches1 MatchSlice, player0 *Player) (MatchSlice, error) {
	setter := &MatchSetter{
		Opponent1ID: omitnull.From(player0.ID),
	}

	err := Matches.Update(ctx, exec, setter, matches1...)
	if err != nil {
		return nil, fmt.Errorf("attachPlayerOpponent1Matches0: %w", err)
	}

	return matches1, nil
}

func (player0 *Player) InsertOpponent1Matches(ctx context.Context, exec bob.Executor, related ...*MatchSetter) error {
	if len(related) == 0 {
		return nil
	}

	matches1, err := insertPlayerOpponent1Matches0(ctx, exec, related, player0)
	if err != nil {
		return err
	}

	player0.R.Opponent1Matches = append(player0.R.Opponent1Matches, matches1...)

	for _, rel := range matches1 {
		rel.R.Opponent1Player = player0
	}
	return nil
}

func (player0 *Player) AttachOpponent1Matches(ctx context.Context, exec bob.Executor, related ...*Match) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	matches1 := MatchSlice(related)

	_, err = attachPlayerOpponent1Matches0(ctx, exec, len(related), matches1, player0)
	if err != nil {
		return err
	}

	player0.R.Opponent1Matches = append(player0.R.Opponent1Matches, matches1...)

	for _, rel := range related {
		rel.R.Opponent1Player = player0
	}

	return nil
}

func insertPlayerOpponent2Matches0(ctx context.Context, exec bob.Executor, matches1 []*MatchSetter, player0 *Player) (MatchSlice, error) {
	for i := range matches1 {
		matches1[i].Opponent2ID = omitnull.From(player0.ID)
	}

	ret, err := Matches.InsertMany(ctx, exec, matches1...)
	if err != nil {
		return ret, fmt.Errorf("insertPlayerOpponent2Matches0: %w", err)
	}

	return ret, nil
}

func attachPlayerOpponent2Matches0(ctx context.Context, exec bob.Executor, count int, matches1 MatchSlice, player0 *Player) (MatchSlice, error) {
	setter := &MatchSetter{
		Opponent2ID: omitnull.From(player0.ID),
	}

	err := Matches.Update(ctx, exec, setter, matches1...)
	if err != nil {
		return nil, fmt.Errorf("attachPlayerOpponent2Matches0: %w", err)
	}

	return matches1, nil
}

func (player0 *Player) InsertOpponent2Matches(ctx context.Context, exec bob.Executor, related ...*MatchSetter) error {
	if len(related) == 0 {
		return nil
	}

	matches1, err := insertPlayerOpponent2Matches0(ctx, exec, related, player0)
	if err != nil {
		return err
	}

	player0.R.Opponent2Matches = append(player0.R.Opponent2Matches, matches1...)

	for _, rel := range matches1 {
		rel.R.Opponent2Player = player0
	}
	return nil
}

func (player0 *Player) AttachOpponent2Matches(ctx context.Context, exec bob.Executor, related ...*Match) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	matches1 := MatchSlice(related)

	_, err = attachPlayerOpponent2Matches0(ctx, exec, len(related), matches1, player0)
	if err != nil {
		return err
	}

	player0.R.Opponent2Matches = append(player0.R.Opponent2Matches, matches1...)

	for _, rel := range related {
		rel.R.Opponent2Player = player0
	}

	return nil
}

func attachPlayerTournaments0(ctx context.Context, exec bob.Executor, count int, player0 *Player, tournaments2 TournamentSlice) (TournamentParticipationSlice, error) {
	setters := make([]*TournamentParticipationSetter, count)
	for i := 0; i < count; i++ {
		setters[i] = &TournamentParticipationSetter{
			ParticipantID: omit.From(player0.ID),
			TournamentID:  omit.From(tournaments2[i].ID),
		}
	}

	tournamentParticipations1, err := TournamentParticipations.InsertMany(ctx, exec, setters...)
	if err != nil {
		return nil, fmt.Errorf("attachPlayerTournaments0: %w", err)
	}

	return tournamentParticipations1, nil
}

func (player0 *Player) InsertTournaments(ctx context.Context, exec bob.Executor, related ...*TournamentSetter) error {
	if len(related) == 0 {
		return nil
	}

	inserted, err := Tournaments.InsertMany(ctx, exec, related...)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}
	tournaments2 := TournamentSlice(inserted)

	_, err = attachPlayerTournaments0(ctx, exec, len(related), player0, tournaments2)
	if err != nil {
		return err
	}

	player0.R.Tournaments = append(player0.R.Tournaments, tournaments2...)

	for _, rel := range tournaments2 {
		rel.R.Players = append(rel.R.Players, player0)
	}
	return nil
}

func (player0 *Player) AttachTournaments(ctx context.Context, exec bob.Executor, related ...*Tournament) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	tournaments2 := TournamentSlice(related)

	_, err = attachPlayerTournaments0(ctx, exec, len(related), player0, tournaments2)
	if err != nil {
		return err
	}

	player0.R.Tournaments = append(player0.R.Tournaments, tournaments2...)

	for _, rel := range related {
		rel.R.Players = append(rel.R.Players, player0)
	}

	return nil
}
