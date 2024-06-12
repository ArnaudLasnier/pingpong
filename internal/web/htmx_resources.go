package web

import "path"

type htmxResource string

func (r htmxResource) String() string {
	return string(r)
}

func (r htmxResource) IDSelector() string {
	return "#" + string(r)
}

func (r htmxResource) Endpoint() string {
	return "/" + string(r)
}

func (r htmxResource) GetEndpoint() string {
	return "GET /" + string(r)
}

func (r htmxResource) GetEndpointWithPathValues(values ...pathValue) string {
	var dynamicSegments []string
	for _, value := range values {
		dynamicSegments = append(dynamicSegments, value.DynamicSegment())
	}
	suffix := path.Join(dynamicSegments...)
	return "GET /" + string(r) + "/" + suffix
}

func (r htmxResource) PostEndpoint() string {
	return "POST /" + string(r)
}

const (
	createPlayerModalResource      htmxResource = "create-player-modal"
	createPlayerFormResource       htmxResource = "create-player-form"
	deletePlayerModalResource      htmxResource = "delete-player-modal"
	createTournamentModalResource  htmxResource = "create-tournament-modal"
	deleteTournamentModalResource  htmxResource = "delete-tournament-modal"
	createTournamentFormResource   htmxResource = "create-tournament-form"
	registerPlayerModalResource    htmxResource = "register-player-modal"
	registerPlayerFormResource     htmxResource = "register-player-form"
	registerPlayerButtonResource   htmxResource = "register-player-button"
	deregisterPlayerButtonResource htmxResource = "deregister-player-button"
	addParticipantModalResource    htmxResource = "add-participant-modal"
	addParticipantFormResource     htmxResource = "add-participant-form"
)
