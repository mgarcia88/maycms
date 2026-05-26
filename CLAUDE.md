# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

For architecture details, patterns, and conventions, see [ARCHITECTURE.md](ARCHITECTURE.md).

## Commands

```bash
# Run the server (requires .env with DSN set)
go run cmd/main.go

# Build
go build -o server ./cmd/main.go

# Spin up the PostgreSQL database
docker compose up -d

# Tidy dependencies
go mod tidy

# Run all tests
go test ./...

# Run a specific package tests
go test ./internal/domain/usecases/...
```

## Environment

Requires a `.env` file at the project root:

```
DSN=postgres://user:password@localhost:5432/dbname?sslmode=disable
POSTGRES_USER=...
POSTGRES_PASSWORD=...
POSTGRES_DB=...
```

`docker-compose.yml` starts only the database (`may-db` on port 5432). The Go server runs separately via `go run`.
