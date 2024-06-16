# Ping-Pong

A web application written in Go to organize table tennis tournaments.

## Getting Started

Run:

```sh
go run ./cmd/devdb
```

This will start a PostgreSQL container and update your `.envrc` configuration.

Then, in another shell <small>(make sure you have <a href="https://direnv.net"><code>direnv</code></a> installed so that it can load environment variables from `.envrc`)</small>:

```sh
go run ./cmd/pingpong
```

Or, if you want live reload:

```sh
air
```

## Building for Production

Run:

```sh
go build -o ./bin/pingpong ./cmd/pingpong
chmod +x ./bin/pingpong
```

and deploy `./bin/pingpong`.

## Regenerate Database Models & Factories

```sh
go generate ./internal/database
```
