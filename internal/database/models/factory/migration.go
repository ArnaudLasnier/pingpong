// Code generated by the local DBGEN tool. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"

	models "github.com/ArnaudLasnier/pingpong/internal/database/models"
	"github.com/aarondl/opt/omit"
	"github.com/jaswdr/faker/v2"
	"github.com/stephenafamo/bob"
)

type MigrationMod interface {
	Apply(*MigrationTemplate)
}

type MigrationModFunc func(*MigrationTemplate)

func (f MigrationModFunc) Apply(n *MigrationTemplate) {
	f(n)
}

type MigrationModSlice []MigrationMod

func (mods MigrationModSlice) Apply(n *MigrationTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// MigrationTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type MigrationTemplate struct {
	Version func() int64
	Dirty   func() bool

	f *Factory
}

// Apply mods to the MigrationTemplate
func (o *MigrationTemplate) Apply(mods ...MigrationMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.Migration
// this does nothing with the relationship templates
func (o MigrationTemplate) toModel() *models.Migration {
	m := &models.Migration{}

	if o.Version != nil {
		m.Version = o.Version()
	}
	if o.Dirty != nil {
		m.Dirty = o.Dirty()
	}

	return m
}

// toModels returns an models.MigrationSlice
// this does nothing with the relationship templates
func (o MigrationTemplate) toModels(number int) models.MigrationSlice {
	m := make(models.MigrationSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.Migration
// according to the relationships in the template. Nothing is inserted into the db
func (t MigrationTemplate) setModelRels(o *models.Migration) {}

// BuildSetter returns an *models.MigrationSetter
// this does nothing with the relationship templates
func (o MigrationTemplate) BuildSetter() *models.MigrationSetter {
	m := &models.MigrationSetter{}

	if o.Version != nil {
		m.Version = omit.From(o.Version())
	}
	if o.Dirty != nil {
		m.Dirty = omit.From(o.Dirty())
	}

	return m
}

// BuildManySetter returns an []*models.MigrationSetter
// this does nothing with the relationship templates
func (o MigrationTemplate) BuildManySetter(number int) []*models.MigrationSetter {
	m := make([]*models.MigrationSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.Migration
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use MigrationTemplate.Create
func (o MigrationTemplate) Build() *models.Migration {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.MigrationSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use MigrationTemplate.CreateMany
func (o MigrationTemplate) BuildMany(number int) models.MigrationSlice {
	m := make(models.MigrationSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatableMigration(m *models.MigrationSetter) {
	if m.Version.IsUnset() {
		m.Version = omit.From(random[int64](nil))
	}
	if m.Dirty.IsUnset() {
		m.Dirty = omit.From(random[bool](nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.Migration
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *MigrationTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.Migration) (context.Context, error) {
	var err error

	return ctx, err
}

// Create builds a migration and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *MigrationTemplate) Create(ctx context.Context, exec bob.Executor) (*models.Migration, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// create builds a migration and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *MigrationTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.Migration, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatableMigration(opt)

	m, err := models.Migrations.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = migrationCtx.WithValue(ctx, m)

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple migrations and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o MigrationTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.MigrationSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// createMany builds multiple migrations and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o MigrationTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.MigrationSlice, error) {
	var err error
	m := make(models.MigrationSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// Migration has methods that act as mods for the MigrationTemplate
var MigrationMods migrationMods

type migrationMods struct{}

func (m migrationMods) RandomizeAllColumns(f *faker.Faker) MigrationMod {
	return MigrationModSlice{
		MigrationMods.RandomVersion(f),
		MigrationMods.RandomDirty(f),
	}
}

// Set the model columns to this value
func (m migrationMods) Version(val int64) MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		o.Version = func() int64 { return val }
	})
}

// Set the Column from the function
func (m migrationMods) VersionFunc(f func() int64) MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		o.Version = f
	})
}

// Clear any values for the column
func (m migrationMods) UnsetVersion() MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		o.Version = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m migrationMods) RandomVersion(f *faker.Faker) MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		o.Version = func() int64 {
			return random[int64](f)
		}
	})
}

func (m migrationMods) ensureVersion(f *faker.Faker) MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		if o.Version != nil {
			return
		}

		o.Version = func() int64 {
			return random[int64](f)
		}
	})
}

// Set the model columns to this value
func (m migrationMods) Dirty(val bool) MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		o.Dirty = func() bool { return val }
	})
}

// Set the Column from the function
func (m migrationMods) DirtyFunc(f func() bool) MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		o.Dirty = f
	})
}

// Clear any values for the column
func (m migrationMods) UnsetDirty() MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		o.Dirty = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m migrationMods) RandomDirty(f *faker.Faker) MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		o.Dirty = func() bool {
			return random[bool](f)
		}
	})
}

func (m migrationMods) ensureDirty(f *faker.Faker) MigrationMod {
	return MigrationModFunc(func(o *MigrationTemplate) {
		if o.Dirty != nil {
			return
		}

		o.Dirty = func() bool {
			return random[bool](f)
		}
	})
}
