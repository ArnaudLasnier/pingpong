// Code generated by the local DBGEN tool. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"

	models "github.com/ArnaudLasnier/pingpong/internal/database/models"
)

type contextKey string

var (
	matchCtx                   = newContextual[*models.Match]("match")
	migrationCtx               = newContextual[*models.Migration]("migration")
	playerCtx                  = newContextual[*models.Player]("player")
	tournamentCtx              = newContextual[*models.Tournament]("tournament")
	tournamentParticipationCtx = newContextual[*models.TournamentParticipation]("tournamentParticipation")
)

// Contextual is a convienience wrapper around context.WithValue and context.Value
type contextual[V any] struct {
	key contextKey
}

func newContextual[V any](key string) contextual[V] {
	return contextual[V]{key: contextKey(key)}
}

func (k contextual[V]) WithValue(ctx context.Context, val V) context.Context {
	return context.WithValue(ctx, k.key, val)
}

func (k contextual[V]) Value(ctx context.Context) (V, bool) {
	v, ok := ctx.Value(k.key).(V)
	return v, ok
}
