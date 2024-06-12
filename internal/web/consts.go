package web

import "github.com/ArnaudLasnier/pingpong/internal/webutils"

const (
	pathKeytournamentID webutils.PathKey = "tournamentID"
	pathKeyPlayerID     webutils.PathKey = "playerID"
)

const (
	formKeyPlayerID        webutils.FormKey = "playerID"
	formKeyPlayerFirstName webutils.FormKey = "playerFirstName"
	formKeyPlayerLastName  webutils.FormKey = "playerLastName"
	formKeyPlayerEmail     webutils.FormKey = "playerEmail"
	formKeyTournamentID    webutils.FormKey = "tournamentID"
	formKeyTournamentTitle webutils.FormKey = "tournamentTitle"
)

const (
	fragmentCreatePlayerModal      webutils.Fragment = "create-player-modal"
	fragmentCreatePlayerForm       webutils.Fragment = "create-player-form"
	fragmentDeletePlayerModal      webutils.Fragment = "delete-player-modal"
	fragmentDeletePlayer           webutils.Fragment = "delete-player"
	fragmentCreateTournamentModal  webutils.Fragment = "create-tournament-modal"
	fragmentDeleteTournamentModal  webutils.Fragment = "delete-tournament-modal"
	fragmentCreateTournamentForm   webutils.Fragment = "create-tournament-form"
	fragmentRegisterPlayerModal    webutils.Fragment = "register-player-modal"
	fragmentRegisterPlayerForm     webutils.Fragment = "register-player-form"
	fragmentRegisterPlayerButton   webutils.Fragment = "register-player-button"
	fragmentDeregisterPlayerButton webutils.Fragment = "deregister-player-button"
	fragmentAddParticipantModal    webutils.Fragment = "add-participant-modal"
	fragmentAddParticipantForm     webutils.Fragment = "add-participant-form"
)
