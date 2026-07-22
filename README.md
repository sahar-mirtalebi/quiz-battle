# Game App

Backend for an online quiz game. MySQL stores persistent data and Redis stores
the matchmaking waiting lists.

## Structure

```text
cmd/server/                 application entrypoint
internal/
  adapter/                  MySQL and Redis adapters
  auth/                     authentication and authorization
  game/                     game domain
  matching/                 matchmaking feature
  quiz/                     question and category domain
  user/                     user feature and backoffice
  config/                   application configuration
  server/                   HTTP server wiring
  shared/                   shared error utilities
```

## Run locally

```bash
cp .env.example .env
docker compose up -d --wait
go run ./cmd/server
```

Check the application:

```bash
curl http://localhost:8080/health-check
```

## Verify

```bash
go test ./...
go vet ./...
docker compose config --quiet
```

Migrations run automatically when the server starts.
