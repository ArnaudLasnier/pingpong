// Code generated by the local DBGEN tool. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"

	models "github.com/ArnaudLasnier/pingpong/internal/tournamentdatabase/models"
	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"github.com/stephenafamo/bob"
)

type PlayerMod interface {
	Apply(*PlayerTemplate)
}

type PlayerModFunc func(*PlayerTemplate)

func (f PlayerModFunc) Apply(n *PlayerTemplate) {
	f(n)
}

type PlayerModSlice []PlayerMod

func (mods PlayerModSlice) Apply(n *PlayerTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// PlayerTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type PlayerTemplate struct {
	ID        func() uuid.UUID
	FirstName func() string
	LastName  func() string
	Email     func() string

	r playerR
	f *Factory
}

type playerR struct {
	Opponent1Matches []*playerROpponent1MatchesR
	Opponent2Matches []*playerROpponent2MatchesR
	Tournaments      []*playerRTournamentsR
}

type playerROpponent1MatchesR struct {
	number int
	o      *MatchTemplate
}
type playerROpponent2MatchesR struct {
	number int
	o      *MatchTemplate
}
type playerRTournamentsR struct {
	number int
	o      *TournamentTemplate
}

// Apply mods to the PlayerTemplate
func (o *PlayerTemplate) Apply(mods ...PlayerMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.Player
// this does nothing with the relationship templates
func (o PlayerTemplate) toModel() *models.Player {
	m := &models.Player{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.FirstName != nil {
		m.FirstName = o.FirstName()
	}
	if o.LastName != nil {
		m.LastName = o.LastName()
	}
	if o.Email != nil {
		m.Email = o.Email()
	}

	return m
}

// toModels returns an models.PlayerSlice
// this does nothing with the relationship templates
func (o PlayerTemplate) toModels(number int) models.PlayerSlice {
	m := make(models.PlayerSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.Player
// according to the relationships in the template. Nothing is inserted into the db
func (t PlayerTemplate) setModelRels(o *models.Player) {
	if t.r.Opponent1Matches != nil {
		rel := models.MatchSlice{}
		for _, r := range t.r.Opponent1Matches {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.Opponent1ID = null.From(o.ID)
				rel.R.Opponent1Player = o
			}
			rel = append(rel, related...)
		}
		o.R.Opponent1Matches = rel
	}

	if t.r.Opponent2Matches != nil {
		rel := models.MatchSlice{}
		for _, r := range t.r.Opponent2Matches {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.Opponent2ID = null.From(o.ID)
				rel.R.Opponent2Player = o
			}
			rel = append(rel, related...)
		}
		o.R.Opponent2Matches = rel
	}

	if t.r.Tournaments != nil {
		rel := models.TournamentSlice{}
		for _, r := range t.r.Tournaments {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.R.Players = append(rel.R.Players, o)
			}
			rel = append(rel, related...)
		}
		o.R.Tournaments = rel
	}
}

// BuildSetter returns an *models.PlayerSetter
// this does nothing with the relationship templates
func (o PlayerTemplate) BuildSetter() *models.PlayerSetter {
	m := &models.PlayerSetter{}

	if o.ID != nil {
		m.ID = omit.From(o.ID())
	}
	if o.FirstName != nil {
		m.FirstName = omit.From(o.FirstName())
	}
	if o.LastName != nil {
		m.LastName = omit.From(o.LastName())
	}
	if o.Email != nil {
		m.Email = omit.From(o.Email())
	}

	return m
}

// BuildManySetter returns an []*models.PlayerSetter
// this does nothing with the relationship templates
func (o PlayerTemplate) BuildManySetter(number int) []*models.PlayerSetter {
	m := make([]*models.PlayerSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.Player
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use PlayerTemplate.Create
func (o PlayerTemplate) Build() *models.Player {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.PlayerSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use PlayerTemplate.CreateMany
func (o PlayerTemplate) BuildMany(number int) models.PlayerSlice {
	m := make(models.PlayerSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatablePlayer(m *models.PlayerSetter) {
	if m.FirstName.IsUnset() {
		m.FirstName = omit.From(random[string](nil))
	}
	if m.LastName.IsUnset() {
		m.LastName = omit.From(random[string](nil))
	}
	if m.Email.IsUnset() {
		m.Email = omit.From(random[string](nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.Player
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *PlayerTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.Player) (context.Context, error) {
	var err error

	if o.r.Opponent1Matches != nil {
		for _, r := range o.r.Opponent1Matches {
			var rel0 models.MatchSlice
			ctx, rel0, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachOpponent1Matches(ctx, exec, rel0...)
			if err != nil {
				return ctx, err
			}
		}
	}

	if o.r.Opponent2Matches != nil {
		for _, r := range o.r.Opponent2Matches {
			var rel1 models.MatchSlice
			ctx, rel1, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachOpponent2Matches(ctx, exec, rel1...)
			if err != nil {
				return ctx, err
			}
		}
	}

	if o.r.Tournaments != nil {
		for _, r := range o.r.Tournaments {
			var rel2 models.TournamentSlice
			ctx, rel2, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachTournaments(ctx, exec, rel2...)
			if err != nil {
				return ctx, err
			}
		}
	}

	return ctx, err
}

// Create builds a player and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *PlayerTemplate) Create(ctx context.Context, exec bob.Executor) (*models.Player, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// create builds a player and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *PlayerTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.Player, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatablePlayer(opt)

	m, err := models.Players.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = playerCtx.WithValue(ctx, m)

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple players and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o PlayerTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.PlayerSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// createMany builds multiple players and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o PlayerTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.PlayerSlice, error) {
	var err error
	m := make(models.PlayerSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// Player has methods that act as mods for the PlayerTemplate
var PlayerMods playerMods

type playerMods struct{}

func (m playerMods) RandomizeAllColumns(f *faker.Faker) PlayerMod {
	return PlayerModSlice{
		PlayerMods.RandomID(f),
		PlayerMods.RandomFirstName(f),
		PlayerMods.RandomLastName(f),
		PlayerMods.RandomEmail(f),
	}
}

// Set the model columns to this value
func (m playerMods) ID(val uuid.UUID) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.ID = func() uuid.UUID { return val }
	})
}

// Set the Column from the function
func (m playerMods) IDFunc(f func() uuid.UUID) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m playerMods) UnsetID() PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m playerMods) RandomID(f *faker.Faker) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.ID = func() uuid.UUID {
			return random[uuid.UUID](f)
		}
	})
}

func (m playerMods) ensureID(f *faker.Faker) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		if o.ID != nil {
			return
		}

		o.ID = func() uuid.UUID {
			return random[uuid.UUID](f)
		}
	})
}

// Set the model columns to this value
func (m playerMods) FirstName(val string) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.FirstName = func() string { return val }
	})
}

// Set the Column from the function
func (m playerMods) FirstNameFunc(f func() string) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.FirstName = f
	})
}

// Clear any values for the column
func (m playerMods) UnsetFirstName() PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.FirstName = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m playerMods) RandomFirstName(f *faker.Faker) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.FirstName = func() string {
			return random[string](f)
		}
	})
}

func (m playerMods) ensureFirstName(f *faker.Faker) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		if o.FirstName != nil {
			return
		}

		o.FirstName = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m playerMods) LastName(val string) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.LastName = func() string { return val }
	})
}

// Set the Column from the function
func (m playerMods) LastNameFunc(f func() string) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.LastName = f
	})
}

// Clear any values for the column
func (m playerMods) UnsetLastName() PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.LastName = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m playerMods) RandomLastName(f *faker.Faker) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.LastName = func() string {
			return random[string](f)
		}
	})
}

func (m playerMods) ensureLastName(f *faker.Faker) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		if o.LastName != nil {
			return
		}

		o.LastName = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m playerMods) Email(val string) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.Email = func() string { return val }
	})
}

// Set the Column from the function
func (m playerMods) EmailFunc(f func() string) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.Email = f
	})
}

// Clear any values for the column
func (m playerMods) UnsetEmail() PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.Email = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m playerMods) RandomEmail(f *faker.Faker) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.Email = func() string {
			return random[string](f)
		}
	})
}

func (m playerMods) ensureEmail(f *faker.Faker) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		if o.Email != nil {
			return
		}

		o.Email = func() string {
			return random[string](f)
		}
	})
}

func (m playerMods) WithOpponent1Matches(number int, related *MatchTemplate) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Opponent1Matches = []*playerROpponent1MatchesR{{
			number: number,
			o:      related,
		}}
	})
}

func (m playerMods) WithNewOpponent1Matches(number int, mods ...MatchMod) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		related := o.f.NewMatch(mods...)
		m.WithOpponent1Matches(number, related).Apply(o)
	})
}

func (m playerMods) AddOpponent1Matches(number int, related *MatchTemplate) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Opponent1Matches = append(o.r.Opponent1Matches, &playerROpponent1MatchesR{
			number: number,
			o:      related,
		})
	})
}

func (m playerMods) AddNewOpponent1Matches(number int, mods ...MatchMod) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		related := o.f.NewMatch(mods...)
		m.AddOpponent1Matches(number, related).Apply(o)
	})
}

func (m playerMods) WithoutOpponent1Matches() PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Opponent1Matches = nil
	})
}

func (m playerMods) WithOpponent2Matches(number int, related *MatchTemplate) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Opponent2Matches = []*playerROpponent2MatchesR{{
			number: number,
			o:      related,
		}}
	})
}

func (m playerMods) WithNewOpponent2Matches(number int, mods ...MatchMod) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		related := o.f.NewMatch(mods...)
		m.WithOpponent2Matches(number, related).Apply(o)
	})
}

func (m playerMods) AddOpponent2Matches(number int, related *MatchTemplate) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Opponent2Matches = append(o.r.Opponent2Matches, &playerROpponent2MatchesR{
			number: number,
			o:      related,
		})
	})
}

func (m playerMods) AddNewOpponent2Matches(number int, mods ...MatchMod) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		related := o.f.NewMatch(mods...)
		m.AddOpponent2Matches(number, related).Apply(o)
	})
}

func (m playerMods) WithoutOpponent2Matches() PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Opponent2Matches = nil
	})
}

func (m playerMods) WithTournaments(number int, related *TournamentTemplate) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Tournaments = []*playerRTournamentsR{{
			number: number,
			o:      related,
		}}
	})
}

func (m playerMods) WithNewTournaments(number int, mods ...TournamentMod) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		related := o.f.NewTournament(mods...)
		m.WithTournaments(number, related).Apply(o)
	})
}

func (m playerMods) AddTournaments(number int, related *TournamentTemplate) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Tournaments = append(o.r.Tournaments, &playerRTournamentsR{
			number: number,
			o:      related,
		})
	})
}

func (m playerMods) AddNewTournaments(number int, mods ...TournamentMod) PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		related := o.f.NewTournament(mods...)
		m.AddTournaments(number, related).Apply(o)
	})
}

func (m playerMods) WithoutTournaments() PlayerMod {
	return PlayerModFunc(func(o *PlayerTemplate) {
		o.r.Tournaments = nil
	})
}
