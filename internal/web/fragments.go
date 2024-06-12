package web

import (
	"path"

	"github.com/ArnaudLasnier/pingpong/internal/webutils"
)

type Fragment string

func (fragment Fragment) String() string {
	return string(fragment)
}

func (fragment Fragment) IDSelector() string {
	return "#" + string(fragment)
}

func (fragment Fragment) Endpoint() string {
	return "/" + string(fragment)
}

func (fragment Fragment) GetEndpoint() string {
	return "GET /" + string(fragment)
}

func (fragment Fragment) GetEndpointWithPathValues(values ...webutils.PathKey) string {
	var dynamicSegments []string
	for _, value := range values {
		dynamicSegments = append(dynamicSegments, value.DynamicSegment())
	}
	suffix := path.Join(dynamicSegments...)
	return "GET /" + string(fragment) + "/" + suffix
}

func (fragment Fragment) PostEndpoint() string {
	return "POST /" + string(fragment)
}

const (
	fragmentCreatePlayerModal      Fragment = "create-player-modal"
	fragmentCreatePlayerForm       Fragment = "create-player-form"
	fragmentDeletePlayerModal      Fragment = "delete-player-modal"
	fragmentCreateTournamentModal  Fragment = "create-tournament-modal"
	fragmentDeleteTournamentModal  Fragment = "delete-tournament-modal"
	fragmentCreateTournamentForm   Fragment = "create-tournament-form"
	fragmentRegisterPlayerModal    Fragment = "register-player-modal"
	fragmentRegisterPlayerForm     Fragment = "register-player-form"
	fragmentRegisterPlayerButton   Fragment = "register-player-button"
	fragmentDeregisterPlayerButton Fragment = "deregister-player-button"
	fragmentAddParticipantModal    Fragment = "add-participant-modal"
	fragmentAddParticipantForm     Fragment = "add-participant-form"
)
