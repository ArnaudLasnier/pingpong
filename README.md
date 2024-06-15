# Ping-Pong

A web application written in Go to organize table tennis tournaments.

## Getting Started

Run:

```
make localdeps
```

This will start a PostgreSQL container and update your `.envrc` configuration.

Then, in another shell <small>(make sure you have <a href="https://direnv.net"><code>direnv</code></a> installed so that it can load environment variables from `.envrc`)</small>:

```
make run
```

Or, if you want live reload:

```
make live_reload
```

## Building for Production

Run:

```
make build
```

and deploy `./bin/pingpong`.

## Regenerate Database Models & Factories

```sh
go generate ./internal/database
```
