// Code generated by BobGen psql v0.25.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aarondl/opt/omit"
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
	"github.com/stephenafamo/bob/orm"
)

// TournamentParticipation is an object representing the database table.
type TournamentParticipation struct {
	TournamentID  uuid.UUID `db:"tournament_id,pk" `
	ParticipantID uuid.UUID `db:"participant_id,pk" `

	R tournamentParticipationR `db:"-" `
}

// TournamentParticipationSlice is an alias for a slice of pointers to TournamentParticipation.
// This should almost always be used instead of []*TournamentParticipation.
type TournamentParticipationSlice []*TournamentParticipation

// TournamentParticipations contains methods to work with the tournament_participation table
var TournamentParticipations = psql.NewTablex[*TournamentParticipation, TournamentParticipationSlice, *TournamentParticipationSetter]("", "tournament_participation")

// TournamentParticipationsQuery is a query on the tournament_participation table
type TournamentParticipationsQuery = *psql.ViewQuery[*TournamentParticipation, TournamentParticipationSlice]

// TournamentParticipationsStmt is a prepared statment on tournament_participation
type TournamentParticipationsStmt = bob.QueryStmt[*TournamentParticipation, TournamentParticipationSlice]

// tournamentParticipationR is where relationships are stored.
type tournamentParticipationR struct {
	ParticipantPlayer *Player     // tournament_participation.tournament_participation_participant_id_fkey
	Tournament        *Tournament // tournament_participation.tournament_participation_tournament_id_fkey
}

// TournamentParticipationSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type TournamentParticipationSetter struct {
	TournamentID  omit.Val[uuid.UUID] `db:"tournament_id,pk"`
	ParticipantID omit.Val[uuid.UUID] `db:"participant_id,pk"`
}

func (s TournamentParticipationSetter) SetColumns() []string {
	vals := make([]string, 0, 2)
	if !s.TournamentID.IsUnset() {
		vals = append(vals, "tournament_id")
	}

	if !s.ParticipantID.IsUnset() {
		vals = append(vals, "participant_id")
	}

	return vals
}

func (s TournamentParticipationSetter) Overwrite(t *TournamentParticipation) {
	if !s.TournamentID.IsUnset() {
		t.TournamentID, _ = s.TournamentID.Get()
	}
	if !s.ParticipantID.IsUnset() {
		t.ParticipantID, _ = s.ParticipantID.Get()
	}
}

func (s TournamentParticipationSetter) InsertMod() bob.Mod[*dialect.InsertQuery] {
	vals := make([]bob.Expression, 2)
	if s.TournamentID.IsUnset() {
		vals[0] = psql.Raw("DEFAULT")
	} else {
		vals[0] = psql.Arg(s.TournamentID)
	}

	if s.ParticipantID.IsUnset() {
		vals[1] = psql.Raw("DEFAULT")
	} else {
		vals[1] = psql.Arg(s.ParticipantID)
	}

	return im.Values(vals...)
}

func (s TournamentParticipationSetter) Apply(q *dialect.UpdateQuery) {
	um.Set(s.Expressions()...).Apply(q)
}

func (s TournamentParticipationSetter) Expressions(prefix ...string) []bob.Expression {
	exprs := make([]bob.Expression, 0, 2)

	if !s.TournamentID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "tournament_id")...),
			psql.Arg(s.TournamentID),
		}})
	}

	if !s.ParticipantID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			psql.Quote(append(prefix, "participant_id")...),
			psql.Arg(s.ParticipantID),
		}})
	}

	return exprs
}

type tournamentParticipationColumnNames struct {
	TournamentID  string
	ParticipantID string
}

type tournamentParticipationRelationshipJoins[Q dialect.Joinable] struct {
	ParticipantPlayer bob.Mod[Q]
	Tournament        bob.Mod[Q]
}

func buildTournamentParticipationRelationshipJoins[Q dialect.Joinable](ctx context.Context, typ string) tournamentParticipationRelationshipJoins[Q] {
	return tournamentParticipationRelationshipJoins[Q]{
		ParticipantPlayer: tournamentParticipationsJoinParticipantPlayer[Q](ctx, typ),
		Tournament:        tournamentParticipationsJoinTournament[Q](ctx, typ),
	}
}

func tournamentParticipationsJoin[Q dialect.Joinable](ctx context.Context) joinSet[tournamentParticipationRelationshipJoins[Q]] {
	return joinSet[tournamentParticipationRelationshipJoins[Q]]{
		InnerJoin: buildTournamentParticipationRelationshipJoins[Q](ctx, clause.InnerJoin),
		LeftJoin:  buildTournamentParticipationRelationshipJoins[Q](ctx, clause.LeftJoin),
		RightJoin: buildTournamentParticipationRelationshipJoins[Q](ctx, clause.RightJoin),
	}
}

var TournamentParticipationColumns = struct {
	TournamentID  psql.Expression
	ParticipantID psql.Expression
}{
	TournamentID:  psql.Quote("tournament_participation", "tournament_id"),
	ParticipantID: psql.Quote("tournament_participation", "participant_id"),
}

type tournamentParticipationWhere[Q psql.Filterable] struct {
	TournamentID  psql.WhereMod[Q, uuid.UUID]
	ParticipantID psql.WhereMod[Q, uuid.UUID]
}

func TournamentParticipationWhere[Q psql.Filterable]() tournamentParticipationWhere[Q] {
	return tournamentParticipationWhere[Q]{
		TournamentID:  psql.Where[Q, uuid.UUID](TournamentParticipationColumns.TournamentID),
		ParticipantID: psql.Where[Q, uuid.UUID](TournamentParticipationColumns.ParticipantID),
	}
}

// FindTournamentParticipation retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindTournamentParticipation(ctx context.Context, exec bob.Executor, TournamentIDPK uuid.UUID, ParticipantIDPK uuid.UUID, cols ...string) (*TournamentParticipation, error) {
	if len(cols) == 0 {
		return TournamentParticipations.Query(
			ctx, exec,
			SelectWhere.TournamentParticipations.TournamentID.EQ(TournamentIDPK),
			SelectWhere.TournamentParticipations.ParticipantID.EQ(ParticipantIDPK),
		).One()
	}

	return TournamentParticipations.Query(
		ctx, exec,
		SelectWhere.TournamentParticipations.TournamentID.EQ(TournamentIDPK),
		SelectWhere.TournamentParticipations.ParticipantID.EQ(ParticipantIDPK),
		sm.Columns(TournamentParticipations.Columns().Only(cols...)),
	).One()
}

// TournamentParticipationExists checks the presence of a single record by primary key
func TournamentParticipationExists(ctx context.Context, exec bob.Executor, TournamentIDPK uuid.UUID, ParticipantIDPK uuid.UUID) (bool, error) {
	return TournamentParticipations.Query(
		ctx, exec,
		SelectWhere.TournamentParticipations.TournamentID.EQ(TournamentIDPK),
		SelectWhere.TournamentParticipations.ParticipantID.EQ(ParticipantIDPK),
	).Exists()
}

// PrimaryKeyVals returns the primary key values of the TournamentParticipation
func (o *TournamentParticipation) PrimaryKeyVals() bob.Expression {
	return psql.ArgGroup(
		o.TournamentID,
		o.ParticipantID,
	)
}

// Update uses an executor to update the TournamentParticipation
func (o *TournamentParticipation) Update(ctx context.Context, exec bob.Executor, s *TournamentParticipationSetter) error {
	return TournamentParticipations.Update(ctx, exec, s, o)
}

// Delete deletes a single TournamentParticipation record with an executor
func (o *TournamentParticipation) Delete(ctx context.Context, exec bob.Executor) error {
	return TournamentParticipations.Delete(ctx, exec, o)
}

// Reload refreshes the TournamentParticipation using the executor
func (o *TournamentParticipation) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := TournamentParticipations.Query(
		ctx, exec,
		SelectWhere.TournamentParticipations.TournamentID.EQ(o.TournamentID),
		SelectWhere.TournamentParticipations.ParticipantID.EQ(o.ParticipantID),
	).One()
	if err != nil {
		return err
	}
	o2.R = o.R
	*o = *o2

	return nil
}

func (o TournamentParticipationSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals TournamentParticipationSetter) error {
	return TournamentParticipations.Update(ctx, exec, &vals, o...)
}

func (o TournamentParticipationSlice) DeleteAll(ctx context.Context, exec bob.Executor) error {
	return TournamentParticipations.Delete(ctx, exec, o...)
}

func (o TournamentParticipationSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	var mods []bob.Mod[*dialect.SelectQuery]

	TournamentIDPK := make([]uuid.UUID, len(o))
	ParticipantIDPK := make([]uuid.UUID, len(o))

	for i, o := range o {
		TournamentIDPK[i] = o.TournamentID
		ParticipantIDPK[i] = o.ParticipantID
	}

	mods = append(mods,
		SelectWhere.TournamentParticipations.TournamentID.In(TournamentIDPK...),
		SelectWhere.TournamentParticipations.ParticipantID.In(ParticipantIDPK...),
	)

	o2, err := TournamentParticipations.Query(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, old := range o {
		for _, new := range o2 {
			if new.TournamentID != old.TournamentID {
				continue
			}
			if new.ParticipantID != old.ParticipantID {
				continue
			}
			new.R = old.R
			*old = *new
			break
		}
	}

	return nil
}

func tournamentParticipationsJoinParticipantPlayer[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, Players.NameAs(ctx)).On(
			PlayerColumns.ID.EQ(TournamentParticipationColumns.ParticipantID),
		),
	}
}

func tournamentParticipationsJoinTournament[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, Tournaments.NameAs(ctx)).On(
			TournamentColumns.ID.EQ(TournamentParticipationColumns.TournamentID),
		),
	}
}

// ParticipantPlayer starts a query for related objects on player
func (o *TournamentParticipation) ParticipantPlayer(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) PlayersQuery {
	return Players.Query(ctx, exec, append(mods,
		sm.Where(PlayerColumns.ID.EQ(psql.Arg(o.ParticipantID))),
	)...)
}

func (os TournamentParticipationSlice) ParticipantPlayer(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) PlayersQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = psql.ArgGroup(o.ParticipantID)
	}

	return Players.Query(ctx, exec, append(mods,
		sm.Where(psql.Group(PlayerColumns.ID).In(PKArgs...)),
	)...)
}

// Tournament starts a query for related objects on tournament
func (o *TournamentParticipation) Tournament(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) TournamentsQuery {
	return Tournaments.Query(ctx, exec, append(mods,
		sm.Where(TournamentColumns.ID.EQ(psql.Arg(o.TournamentID))),
	)...)
}

func (os TournamentParticipationSlice) Tournament(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) TournamentsQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = psql.ArgGroup(o.TournamentID)
	}

	return Tournaments.Query(ctx, exec, append(mods,
		sm.Where(psql.Group(TournamentColumns.ID).In(PKArgs...)),
	)...)
}

func (o *TournamentParticipation) Preload(name string, retrieved any) error {
	if o == nil {
		return nil
	}

	switch name {
	case "ParticipantPlayer":
		rel, ok := retrieved.(*Player)
		if !ok {
			return fmt.Errorf("tournamentParticipation cannot load %T as %q", retrieved, name)
		}

		o.R.ParticipantPlayer = rel

		return nil
	case "Tournament":
		rel, ok := retrieved.(*Tournament)
		if !ok {
			return fmt.Errorf("tournamentParticipation cannot load %T as %q", retrieved, name)
		}

		o.R.Tournament = rel

		return nil
	default:
		return fmt.Errorf("tournamentParticipation has no relationship %q", name)
	}
}

func PreloadTournamentParticipationParticipantPlayer(opts ...psql.PreloadOption) psql.Preloader {
	return psql.Preload[*Player, PlayerSlice](orm.Relationship{
		Name: "ParticipantPlayer",
		Sides: []orm.RelSide{
			{
				From: "tournament_participation",
				To:   TableNames.Players,
				ToExpr: func(ctx context.Context) bob.Expression {
					return Players.Name(ctx)
				},
				FromColumns: []string{
					ColumnNames.TournamentParticipations.ParticipantID,
				},
				ToColumns: []string{
					ColumnNames.Players.ID,
				},
			},
		},
	}, Players.Columns().Names(), opts...)
}

func ThenLoadTournamentParticipationParticipantPlayer(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadTournamentParticipationParticipantPlayer(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load TournamentParticipationParticipantPlayer", retrieved)
		}

		err := loader.LoadTournamentParticipationParticipantPlayer(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadTournamentParticipationParticipantPlayer loads the tournamentParticipation's ParticipantPlayer into the .R struct
func (o *TournamentParticipation) LoadTournamentParticipationParticipantPlayer(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.ParticipantPlayer = nil

	related, err := o.ParticipantPlayer(ctx, exec, mods...).One()
	if err != nil {
		return err
	}

	o.R.ParticipantPlayer = related
	return nil
}

// LoadTournamentParticipationParticipantPlayer loads the tournamentParticipation's ParticipantPlayer into the .R struct
func (os TournamentParticipationSlice) LoadTournamentParticipationParticipantPlayer(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	players, err := os.ParticipantPlayer(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, o := range os {
		for _, rel := range players {
			if o.ParticipantID != rel.ID {
				continue
			}

			o.R.ParticipantPlayer = rel
			break
		}
	}

	return nil
}

func PreloadTournamentParticipationTournament(opts ...psql.PreloadOption) psql.Preloader {
	return psql.Preload[*Tournament, TournamentSlice](orm.Relationship{
		Name: "Tournament",
		Sides: []orm.RelSide{
			{
				From: "tournament_participation",
				To:   TableNames.Tournaments,
				ToExpr: func(ctx context.Context) bob.Expression {
					return Tournaments.Name(ctx)
				},
				FromColumns: []string{
					ColumnNames.TournamentParticipations.TournamentID,
				},
				ToColumns: []string{
					ColumnNames.Tournaments.ID,
				},
			},
		},
	}, Tournaments.Columns().Names(), opts...)
}

func ThenLoadTournamentParticipationTournament(queryMods ...bob.Mod[*dialect.SelectQuery]) psql.Loader {
	return psql.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadTournamentParticipationTournament(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load TournamentParticipationTournament", retrieved)
		}

		err := loader.LoadTournamentParticipationTournament(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadTournamentParticipationTournament loads the tournamentParticipation's Tournament into the .R struct
func (o *TournamentParticipation) LoadTournamentParticipationTournament(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Tournament = nil

	related, err := o.Tournament(ctx, exec, mods...).One()
	if err != nil {
		return err
	}

	o.R.Tournament = related
	return nil
}

// LoadTournamentParticipationTournament loads the tournamentParticipation's Tournament into the .R struct
func (os TournamentParticipationSlice) LoadTournamentParticipationTournament(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	tournaments, err := os.Tournament(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, o := range os {
		for _, rel := range tournaments {
			if o.TournamentID != rel.ID {
				continue
			}

			o.R.Tournament = rel
			break
		}
	}

	return nil
}

func attachTournamentParticipationParticipantPlayer0(ctx context.Context, exec bob.Executor, count int, tournamentParticipation0 *TournamentParticipation, player1 *Player) (*TournamentParticipation, error) {
	setter := &TournamentParticipationSetter{
		ParticipantID: omit.From(player1.ID),
	}

	err := TournamentParticipations.Update(ctx, exec, setter, tournamentParticipation0)
	if err != nil {
		return nil, fmt.Errorf("attachTournamentParticipationParticipantPlayer0: %w", err)
	}

	return tournamentParticipation0, nil
}

func (tournamentParticipation0 *TournamentParticipation) InsertParticipantPlayer(ctx context.Context, exec bob.Executor, related *PlayerSetter) error {
	player1, err := Players.Insert(ctx, exec, related)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}

	_, err = attachTournamentParticipationParticipantPlayer0(ctx, exec, 1, tournamentParticipation0, player1)
	if err != nil {
		return err
	}

	tournamentParticipation0.R.ParticipantPlayer = player1

	return nil
}

func (tournamentParticipation0 *TournamentParticipation) AttachParticipantPlayer(ctx context.Context, exec bob.Executor, player1 *Player) error {
	var err error

	_, err = attachTournamentParticipationParticipantPlayer0(ctx, exec, 1, tournamentParticipation0, player1)
	if err != nil {
		return err
	}

	tournamentParticipation0.R.ParticipantPlayer = player1

	return nil
}

func attachTournamentParticipationTournament0(ctx context.Context, exec bob.Executor, count int, tournamentParticipation0 *TournamentParticipation, tournament1 *Tournament) (*TournamentParticipation, error) {
	setter := &TournamentParticipationSetter{
		TournamentID: omit.From(tournament1.ID),
	}

	err := TournamentParticipations.Update(ctx, exec, setter, tournamentParticipation0)
	if err != nil {
		return nil, fmt.Errorf("attachTournamentParticipationTournament0: %w", err)
	}

	return tournamentParticipation0, nil
}

func (tournamentParticipation0 *TournamentParticipation) InsertTournament(ctx context.Context, exec bob.Executor, related *TournamentSetter) error {
	tournament1, err := Tournaments.Insert(ctx, exec, related)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}

	_, err = attachTournamentParticipationTournament0(ctx, exec, 1, tournamentParticipation0, tournament1)
	if err != nil {
		return err
	}

	tournamentParticipation0.R.Tournament = tournament1

	return nil
}

func (tournamentParticipation0 *TournamentParticipation) AttachTournament(ctx context.Context, exec bob.Executor, tournament1 *Tournament) error {
	var err error

	_, err = attachTournamentParticipationTournament0(ctx, exec, 1, tournamentParticipation0, tournament1)
	if err != nil {
		return err
	}

	tournamentParticipation0.R.Tournament = tournament1

	return nil
}
