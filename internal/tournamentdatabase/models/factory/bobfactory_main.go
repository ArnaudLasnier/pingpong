// Code generated by the local DBGEN tool. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

type Factory struct {
	baseMatchMods                   MatchModSlice
	basePlayerMods                  PlayerModSlice
	baseTournamentMods              TournamentModSlice
	baseTournamentParticipationMods TournamentParticipationModSlice
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) NewMatch(mods ...MatchMod) *MatchTemplate {
	o := &MatchTemplate{f: f}

	if f != nil {
		f.baseMatchMods.Apply(o)
	}

	MatchModSlice(mods).Apply(o)

	return o
}

func (f *Factory) NewPlayer(mods ...PlayerMod) *PlayerTemplate {
	o := &PlayerTemplate{f: f}

	if f != nil {
		f.basePlayerMods.Apply(o)
	}

	PlayerModSlice(mods).Apply(o)

	return o
}

func (f *Factory) NewTournament(mods ...TournamentMod) *TournamentTemplate {
	o := &TournamentTemplate{f: f}

	if f != nil {
		f.baseTournamentMods.Apply(o)
	}

	TournamentModSlice(mods).Apply(o)

	return o
}

func (f *Factory) NewTournamentParticipation(mods ...TournamentParticipationMod) *TournamentParticipationTemplate {
	o := &TournamentParticipationTemplate{f: f}

	if f != nil {
		f.baseTournamentParticipationMods.Apply(o)
	}

	TournamentParticipationModSlice(mods).Apply(o)

	return o
}

func (f *Factory) ClearBaseMatchMods() {
	f.baseMatchMods = nil
}

func (f *Factory) AddBaseMatchMod(mods ...MatchMod) {
	f.baseMatchMods = append(f.baseMatchMods, mods...)
}

func (f *Factory) ClearBasePlayerMods() {
	f.basePlayerMods = nil
}

func (f *Factory) AddBasePlayerMod(mods ...PlayerMod) {
	f.basePlayerMods = append(f.basePlayerMods, mods...)
}

func (f *Factory) ClearBaseTournamentMods() {
	f.baseTournamentMods = nil
}

func (f *Factory) AddBaseTournamentMod(mods ...TournamentMod) {
	f.baseTournamentMods = append(f.baseTournamentMods, mods...)
}

func (f *Factory) ClearBaseTournamentParticipationMods() {
	f.baseTournamentParticipationMods = nil
}

func (f *Factory) AddBaseTournamentParticipationMod(mods ...TournamentParticipationMod) {
	f.baseTournamentParticipationMods = append(f.baseTournamentParticipationMods, mods...)
}
