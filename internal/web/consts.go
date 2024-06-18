package web

import (
	"net/http"

	"github.com/ArnaudLasnier/pingpong/internal/webutils"
)

const (
	pathKeyTournamentID webutils.PathKey = "tournamentID"
	pathKeyPlayerID     webutils.PathKey = "playerID"
)

const (
	formKeyPlayerID        webutils.FormKey = "playerID"
	formKeyPlayerUsername  webutils.FormKey = "playerUsername"
	formKeyTournamentID    webutils.FormKey = "tournamentID"
	formKeyTournamentTitle webutils.FormKey = "tournamentTitle"
)

const (
	fragmentCreatePlayerModal            webutils.Fragment = "create-player-modal"
	fragmentCreatePlayerForm             webutils.Fragment = "create-player-form"
	fragmentDeletePlayerModal            webutils.Fragment = "delete-player-modal"
	fragmentDeletePlayer                 webutils.Fragment = "delete-player"
	fragmentPlayersTable                 webutils.Fragment = "players-table"
	fragmentCreateTournamentModal        webutils.Fragment = "create-tournament-modal"
	fragmentDeleteTournamentModal        webutils.Fragment = "delete-tournament-modal"
	fragmentDeleteTournamentModalContent webutils.Fragment = "delete-tournament-modal-content"
	fragmentDeleteTournament             webutils.Fragment = "delete-tournament"
	fragmentTournamentsTable             webutils.Fragment = "tournaments-table"
	fragmentCreateTournamentForm         webutils.Fragment = "create-tournament-form"
	fragmentRegisterPlayerModal          webutils.Fragment = "register-player-modal"
	fragmentRegisterPlayerForm           webutils.Fragment = "register-player-form"
	fragmentRegisterPlayerButton         webutils.Fragment = "register-player-button"
	fragmentDeregisterPlayerButton       webutils.Fragment = "deregister-player-button"
	fragmentAddParticipantModal          webutils.Fragment = "add-participant-modal"
	fragmentAddParticipantForm           webutils.Fragment = "add-participant-form"
	fragmentToastContainer               webutils.Fragment = "toast-container"
	fragmentToast                        webutils.Fragment = "toast"
	fragmentToastHeader                  webutils.Fragment = "toast-header"
	fragmentToastHeaderTitle             webutils.Fragment = "toast-header-title"
	fragmentToastBody                    webutils.Fragment = "toast-body"
)

var (
	formActionStartTournament webutils.FormAction = webutils.NewFormAction(http.MethodPost, "/start-tournament")
)

const (
	eventShowSuccess       webutils.Event = "showSuccess"
	eventShowError         webutils.Event = "showError"
	eventPlayerCreated     webutils.Event = "playerCreated"
	eventPlayerDeleted     webutils.Event = "playerDeleted"
	eventTournamentCreated webutils.Event = "tournamentCreated"
	eventTournamentDeleted webutils.Event = "tournamentDeleted"
	eventTournamentStarted webutils.Event = "tournamentStarted"
)
