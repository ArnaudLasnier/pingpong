// Code generated by the local DBGEN tool. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/aarondl/opt/null"
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

// Tournament is an object representing the database table.
type Tournament struct {
	ID        uuid.UUID           `db:"id,pk" `
	Title     string              `db:"title" `
	Status    TournamentStatus    `db:"status" `
	StartedAt null.Val[time.Time] `db:"started_at" `
	EndedAt   null.Val[time.Time] `db:"ended_at" `

	R tournamentR `db:"-" `
}

// TournamentSlice is an alias for a slice of pointers to Tournament.
// This should almost always be used instead of []*Tournament.
type TournamentSlice []*Tournament

// Tournaments contains methods to work with the tournament table
var Tournaments = psql.NewTablex[*Tournament, TournamentSlice, *TournamentSetter]("", "tournament")

// TournamentsQuery is a query on the tournament table
type TournamentsQuery = *psql.ViewQuery[*Tournament, TournamentSlice]

// TournamentsStmt is a prepared statment on tournament
type TournamentsStmt = bob.QueryStmt[*Tournament, TournamentSlice]

// tournamentR is where relationships are stored.
type tournamentR struct {
	Matches MatchSlice  // match.match_tournament_id_fkey
	Players PlayerSlice // tournament_participation.tournament_participation_participant_id_fkeytournament_participation.tournament_participation_tournament_id_fkey
}

// TournamentSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type TournamentSetter struct {
	ID        omit.Val[uuid.UUID]        `db:"id,pk"`
	Title     omit.Val[string]           `db:"title"`
	Status    omit.Val[TournamentStatus] `db:"status"`
	StartedAt omitnull.Val[time.Time]    `db:"started_at"`
	EndedAt   omitnull.Val[time.Time]    `db:"ended_at"`
}

func (s TournamentSetter) SetColumns() []string {
	vals := make([]string, 0, 5)
	if !s.ID.IsUnset() {
		vals = append(vals, "id")
	}

	if !s.Title.IsUnset() {
		vals = append(vals, "title")
	}

	if !s.Status.IsUnset() {
		vals = append(vals, "status")
	}

	if !s.StartedAt.IsUnset() {
		vals = append(vals, "started_at")
	}

	if !s.EndedAt.IsUnset() {
		vals = append(vals, "ended_at")
	}

	return vals
}

func (s TournamentSetter) Overwrite(t *Tournament) {
	if !s.ID.IsUnset() {
		t.ID, _ = s.ID.Get()
	}
	if !s.Title.IsUnset() {
		t.Title, _ = s.Title.Get()
	}
	if !s.Status.IsUnset() {
		t.Status, _ = s.Status.Get()
	}
	if !s.StartedAt.IsUnset() {
		t.StartedAt, _ = s.StartedAt.GetNull()
	}
	if !s.EndedAt.IsUnset() {
		t.EndedAt, _ = s.EndedAt.GetNull()
	}
}

func (s TournamentSetter) InsertMod() bob.Mod[*dialect.InsertQuery] {
	vals := make([]bob.Expression, 5)
	if s.ID.IsUnset() {
		vals[0] = psql.Raw("DEFAULT")
	} else {
		vals[0] = psql.Arg(s.ID)
	}

	if s.Title.IsUnset() {
		vals[1] = psql.Raw("DEFAULT")
	} else {
		vals[1] = psql.Arg(s.Title)
	}

	if s.Status.IsUnset() {
		vals[2] = psql.Raw("DEFAULT")
	} else {
		vals[2] = psql.Arg(s.Status)
	}

	if s.StartedAt.IsUnset() {
		vals[3] = psql.Raw("DEFAULT")
	} else {
		vals[3] = psql.Arg(s.StartedAt)
	}

	if s.EndedAt.IsUnset() {
		vals[4] = psql.Raw("DEFAULT")
	} else {
		vals[4] = psql.Arg(s.EndedAt)
	}

	return im.Values(vals...)
}

func (s TournamentSetter) Apply(q *dialect.UpdateQuery) {
	um.Set(s.Expressions()...).Apply(q)
}

func (s TournamentSetter) Expressions(prefix ...string) []bob.Expression {
	exprs := make([]bob.Expression, 0, 5)

	if !s.ID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "id")...),
			psql.Arg(s.ID),
		}})
	}

	if !s.Title.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "title")...),
			psql.Arg(s.Title),
		}})
	}

	if !s.Status.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "status")...),
			psql.Arg(s.Status),
		}})
	}

	if !s.StartedAt.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "started_at")...),
			psql.Arg(s.StartedAt),
		}})
	}

	if !s.EndedAt.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "ended_at")...),
			psql.Arg(s.EndedAt),
		}})
	}

	return exprs
}

type tournamentColumnNames struct {
	ID        string
	Title     string
	Status    string
	StartedAt string
	EndedAt   string
}

type tournamentRelationshipJoins[Q dialect.Joinable] struct {
	Matches bob.Mod[Q]
	Players bob.Mod[Q]
}

func buildTournamentRelationshipJoins[Q dialect.Joinable](ctx context.Context, typ string) tournamentRelationshipJoins[Q] {
	return tournamentRelationshipJoins[Q]{
		Matches: tournamentsJoinMatches[Q](ctx, typ),
		Players: tournamentsJoinPlayers[Q](ctx, typ),
	}
}

func tournamentsJoin[Q dialect.Joinable](ctx context.Context) joinSet[tournamentRelationshipJoins[Q]] {
	return joinSet[tournamentRelationshipJoins[Q]]{
		InnerJoin: buildTournamentRelationshipJoins[Q](ctx, clause.InnerJoin),
		LeftJoin:  buildTournamentRelationshipJoins[Q](ctx, clause.LeftJoin),
		RightJoin: buildTournamentRelationshipJoins[Q](ctx, clause.RightJoin),
	}
}

var TournamentColumns = struct {
	ID        psql.Expression
	Title     psql.Expression
	Status    psql.Expression
	StartedAt psql.Expression
	EndedAt   psql.Expression
}{
	ID:        psql.Quote("tournament", "id"),
	Title:     psql.Quote("tournament", "title"),
	Status:    psql.Quote("tournament", "status"),
	StartedAt: psql.Quote("tournament", "started_at"),
	EndedAt:   psql.Quote("tournament", "ended_at"),
}

type tournamentWhere[Q psql.Filterable] struct {
	ID        psql.WhereMod[Q, uuid.UUID]
	Title     psql.WhereMod[Q, string]
	Status    psql.WhereMod[Q, TournamentStatus]
	StartedAt psql.WhereNullMod[Q, time.Time]
	EndedAt   psql.WhereNullMod[Q, time.Time]
}

func TournamentWhere[Q psql.Filterable]() tournamentWhere[Q] {
	return tournamentWhere[Q]{
		ID:        psql.Where[Q, uuid.UUID](TournamentColumns.ID),
		Title:     psql.Where[Q, string](TournamentColumns.Title),
		Status:    psql.Where[Q, TournamentStatus](TournamentColumns.Status),
		StartedAt: psql.WhereNull[Q, time.Time](TournamentColumns.StartedAt),
		EndedAt:   psql.WhereNull[Q, time.Time](TournamentColumns.EndedAt),
	}
}

// FindTournament retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindTournament(ctx context.Context, exec bob.Executor, IDPK uuid.UUID, cols ...string) (*Tournament, error) {
	if len(cols) == 0 {
		return Tournaments.Query(
			ctx, exec,
			SelectWhere.Tournaments.ID.EQ(IDPK),
		).One()
	}

	return Tournaments.Query(
		ctx, exec,
		SelectWhere.Tournaments.ID.EQ(IDPK),
		sm.Columns(Tournaments.Columns().Only(cols...)),
	).One()
}

// TournamentExists checks the presence of a single record by primary key
func TournamentExists(ctx context.Context, exec bob.Executor, IDPK uuid.UUID) (bool, error) {
	return Tournaments.Query(
		ctx, exec,
		SelectWhere.Tournaments.ID.EQ(IDPK),
	).Exists()
}

// PrimaryKeyVals returns the primary key values of the Tournament
func (o *Tournament) PrimaryKeyVals() bob.Expression {
	return psql.Arg(o.ID)
}

// Update uses an executor to update the Tournament
func (o *Tournament) Update(ctx context.Context, exec bob.Executor, s *TournamentSetter) error {
	return Tournaments.Update(ctx, exec, s, o)
}

// Delete deletes a single Tournament record with an executor
func (o *Tournament) Delete(ctx context.Context, exec bob.Executor) error {
	return Tournaments.Delete(ctx, exec, o)
}

// Reload refreshes the Tournament using the executor
func (o *Tournament) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := Tournaments.Query(
		ctx, exec,
		SelectWhere.Tournaments.ID.EQ(o.ID),
	).One()
	if err != nil {
		return err
	}
	o2.R = o.R
	*o = *o2

	return nil
}

func (o TournamentSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals TournamentSetter) error {
	return Tournaments.Update(ctx, exec, &vals, o...)
}

func (o TournamentSlice) DeleteAll(ctx context.Context, exec bob.Executor) error {
	return Tournaments.Delete(ctx, exec, o...)
}

func (o TournamentSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	var mods []bob.Mod[*dialect.SelectQuery]

	IDPK := make([]uuid.UUID, len(o))

	for i, o := range o {
		IDPK[i] = o.ID
	}

	mods = append(mods,
		SelectWhere.Tournaments.ID.In(IDPK...),
	)

	o2, err := Tournaments.Query(ctx, exec, mods...).All()
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

func tournamentsJoinMatches[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, Matches.NameAs(ctx)).On(
			MatchColumns.TournamentID.EQ(TournamentColumns.ID),
		),
	}
}

func tournamentsJoinPlayers[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, TournamentParticipations.NameAs(ctx)).On(
			TournamentParticipationColumns.TournamentID.EQ(TournamentColumns.ID),
		),
		dialect.Join[Q](typ, Players.NameAs(ctx)).On(
			PlayerColumns.ID.EQ(TournamentParticipationColumns.ParticipantID),
		),
	}
}

// Matches starts a query for related objects on match
func (o *Tournament) Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) MatchesQuery {
	return Matches.Query(ctx, exec, append(mods,
		sm.Where(MatchColumns.TournamentID.EQ(psql.Arg(o.ID))),
	)...)
}

func (os TournamentSlice) Matches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) MatchesQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = psql.ArgGroup(o.ID)
	}

	return Matches.Query(ctx, exec, append(mods,
		sm.Where(psql.Group(MatchColumns.TournamentID).In(PKArgs...)),
	)...)
}

// Players starts a query for related objects on player
func (o *Tournament) Players(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) PlayersQuery {
	return Players.Query(ctx, exec, append(mods,
		sm.InnerJoin(TournamentParticipations.NameAs(ctx)).On(
			PlayerColumns.ID.EQ(TournamentParticipationColumns.ParticipantID)),
		sm.Where(TournamentParticipationColumns.TournamentID.EQ(psql.Arg(o.ID))),
	)...)
}

func (os TournamentSlice) Players(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) PlayersQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = psql.ArgGroup(o.ID)
	}

	return Players.Query(ctx, exec, append(mods,
		sm.InnerJoin(TournamentParticipations.NameAs(ctx)).On(
			PlayerColumns.ID.EQ(TournamentParticipationColumns.ParticipantID),
		),
		sm.Where(psql.Group(TournamentParticipationColumns.TournamentID).In(PKArgs...)),
	)...)
}

func (o *Tournament) Preload(name string, retrieved any) error {
	if o == nil {
		return nil
	}

	switch name {
	case "Matches":
		rels, ok := retrieved.(MatchSlice)
		if !ok {
			return fmt.Errorf("tournament cannot load %T as %q", retrieved, name)
		}

		o.R.Matches = rels

		for _, rel := range rels {
			if rel != nil {
				rel.R.Tournament = o
			}
		}
		return nil
	case "Players":
		rels, ok := retrieved.(PlayerSlice)
		if !ok {
			return fmt.Errorf("tournament cannot load %T as %q", retrieved, name)
		}

		o.R.Players = rels

		for _, rel := range rels {
			if rel != nil {
				rel.R.Tournaments = TournamentSlice{o}
			}
		}
		return nil
	default:
		return fmt.Errorf("tournament has no relationship %q", name)
	}
}

func ThenLoadTournamentMatches(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadTournamentMatches(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load TournamentMatches", retrieved)
		}

		err := loader.LoadTournamentMatches(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadTournamentMatches loads the tournament's Matches into the .R struct
func (o *Tournament) LoadTournamentMatches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Matches = nil

	related, err := o.Matches(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, rel := range related {
		rel.R.Tournament = o
	}

	o.R.Matches = related
	return nil
}

// LoadTournamentMatches loads the tournament's Matches into the .R struct
func (os TournamentSlice) LoadTournamentMatches(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	matches, err := os.Matches(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, o := range os {
		o.R.Matches = nil
	}

	for _, o := range os {
		for _, rel := range matches {
			if o.ID != rel.TournamentID {
				continue
			}

			rel.R.Tournament = o

			o.R.Matches = append(o.R.Matches, rel)
		}
	}

	return nil
}

func ThenLoadTournamentPlayers(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadTournamentPlayers(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load TournamentPlayers", retrieved)
		}

		err := loader.LoadTournamentPlayers(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadTournamentPlayers loads the tournament's Players into the .R struct
func (o *Tournament) LoadTournamentPlayers(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Players = nil

	related, err := o.Players(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, rel := range related {
		rel.R.Tournaments = TournamentSlice{o}
	}

	o.R.Players = related
	return nil
}

// LoadTournamentPlayers loads the tournament's Players into the .R struct
func (os TournamentSlice) LoadTournamentPlayers(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	// since we are changing the columns, we need to check if the original columns were set or add the defaults
	sq := dialect.SelectQuery{}
	for _, mod := range mods {
		mod.Apply(&sq)
	}

	if len(sq.SelectList.Columns) == 0 {
		mods = append(mods, sm.Columns(Players.Columns()))
	}

	q := os.Players(ctx, exec, append(
		mods,
		sm.Columns(TournamentParticipationColumns.TournamentID.As("related_tournament.ID")),
	)...)

	IDSlice := []uuid.UUID{}

	mapper := scan.Mod(scan.StructMapper[*Player](), func(ctx context.Context, cols []string) (scan.BeforeFunc, func(any, any) error) {
		return func(row *scan.Row) (any, error) {
				IDSlice = append(IDSlice, *new(uuid.UUID))
				row.ScheduleScan("related_tournament.ID", &IDSlice[len(IDSlice)-1])

				return nil, nil
			},
			func(any, any) error {
				return nil
			}
	})

	players, err := bob.Allx[*Player, PlayerSlice](ctx, exec, q, mapper)
	if err != nil {
		return err
	}

	for _, o := range os {
		o.R.Players = nil
	}

	for _, o := range os {
		for i, rel := range players {
			if o.ID != IDSlice[i] {
				continue
			}

			rel.R.Tournaments = append(rel.R.Tournaments, o)

			o.R.Players = append(o.R.Players, rel)
		}
	}

	return nil
}

func insertTournamentMatches0(ctx context.Context, exec bob.Executor, matches1 []*MatchSetter, tournament0 *Tournament) (MatchSlice, error) {
	for i := range matches1 {
		matches1[i].TournamentID = omit.From(tournament0.ID)
	}

	ret, err := Matches.InsertMany(ctx, exec, matches1...)
	if err != nil {
		return ret, fmt.Errorf("insertTournamentMatches0: %w", err)
	}

	return ret, nil
}

func attachTournamentMatches0(ctx context.Context, exec bob.Executor, count int, matches1 MatchSlice, tournament0 *Tournament) (MatchSlice, error) {
	setter := &MatchSetter{
		TournamentID: omit.From(tournament0.ID),
	}

	err := Matches.Update(ctx, exec, setter, matches1...)
	if err != nil {
		return nil, fmt.Errorf("attachTournamentMatches0: %w", err)
	}

	return matches1, nil
}

func (tournament0 *Tournament) InsertMatches(ctx context.Context, exec bob.Executor, related ...*MatchSetter) error {
	if len(related) == 0 {
		return nil
	}

	matches1, err := insertTournamentMatches0(ctx, exec, related, tournament0)
	if err != nil {
		return err
	}

	tournament0.R.Matches = append(tournament0.R.Matches, matches1...)

	for _, rel := range matches1 {
		rel.R.Tournament = tournament0
	}
	return nil
}

func (tournament0 *Tournament) AttachMatches(ctx context.Context, exec bob.Executor, related ...*Match) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	matches1 := MatchSlice(related)

	_, err = attachTournamentMatches0(ctx, exec, len(related), matches1, tournament0)
	if err != nil {
		return err
	}

	tournament0.R.Matches = append(tournament0.R.Matches, matches1...)

	for _, rel := range related {
		rel.R.Tournament = tournament0
	}

	return nil
}

func attachTournamentPlayers0(ctx context.Context, exec bob.Executor, count int, tournament0 *Tournament, players2 PlayerSlice) (TournamentParticipationSlice, error) {
	setters := make([]*TournamentParticipationSetter, count)
	for i := 0; i < count; i++ {
		setters[i] = &TournamentParticipationSetter{
			TournamentID:  omit.From(tournament0.ID),
			ParticipantID: omit.From(players2[i].ID),
		}
	}

	tournamentParticipations1, err := TournamentParticipations.InsertMany(ctx, exec, setters...)
	if err != nil {
		return nil, fmt.Errorf("attachTournamentPlayers0: %w", err)
	}

	return tournamentParticipations1, nil
}

func (tournament0 *Tournament) InsertPlayers(ctx context.Context, exec bob.Executor, related ...*PlayerSetter) error {
	if len(related) == 0 {
		return nil
	}

	inserted, err := Players.InsertMany(ctx, exec, related...)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}
	players2 := PlayerSlice(inserted)

	_, err = attachTournamentPlayers0(ctx, exec, len(related), tournament0, players2)
	if err != nil {
		return err
	}

	tournament0.R.Players = append(tournament0.R.Players, players2...)

	for _, rel := range players2 {
		rel.R.Tournaments = append(rel.R.Tournaments, tournament0)
	}
	return nil
}

func (tournament0 *Tournament) AttachPlayers(ctx context.Context, exec bob.Executor, related ...*Player) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	players2 := PlayerSlice(related)

	_, err = attachTournamentPlayers0(ctx, exec, len(related), tournament0, players2)
	if err != nil {
		return err
	}

	tournament0.R.Players = append(tournament0.R.Players, players2...)

	for _, rel := range related {
		rel.R.Tournaments = append(rel.R.Tournaments, tournament0)
	}

	return nil
}
