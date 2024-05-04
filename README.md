# Ping Pong

A web application written in Go to organize table tennis tournaments.

## Goals

- basic authentication
- i18n in french
- creating a tournament
- registering users to the tournament
- pairing users randomly
- displaying past, present and future matches
- beginning the tournament and setting due dates
- entering match results

Out of scope:

- pairing users with scores
- letting users register themselves
- sending reminders to users

## Implementation Details

Creating a tournament

- create a tournament draft entity that you can edit later

## Service

PingPongService

- AddPlayer(form AddPlayerForm) *Player
- RemovePlayer(playerId uuid.UUID)
- CreateTournament(startDate time.Time, ...*Player) *Tournament
- RescheduleTournament(tournament *Tournament)
- AddParticipants(tournament *Tournament, ...*Player)
- RemoveParticipants(tournament *Tournament, ...*Player)
- StartTournament(tournament *Tournament)
- EnterMatchResult(match *Match, scoreOpponent1 int8, scoreOpponent2 int8)

- GetWinner(match *Match) *Player
- GetLooser(match *Match) *Player


## Entities and Relationships

Tournament

- id
- isDraft: bool
- startDate: civil.Date
- participants: []*Player

Player

- id
- firstName: string
- tournaments: []*Tournament

Match

- id
- dueDate: civil.Date
- tournament: *Tournament
- opponent1: *Player
- opponent2: *Player
- opponent1Score: sql.Null[int8]
- opponent2Score: sql.Null[int8]

## Dependencies

```sh
go get cloud.google.com/go github.com/caarlos0/env/v11 \
    github.com/jackc/pgx/v5 \
    github.com/golang-migrate/migrate/v4 \
    github.com/google/uuid \
    github.com/stephenafamo/bob \
    github.com/testcontainers/testcontainers-go \
    github.com/aarondl/opt
go install github.com/stephenafamo/bob/gen/bobgen-psql@v0.25.0
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

then

```sh
migrate create -ext sql -dir internal/pingpong/database/migrations -seq -digits 5 first_tables
```