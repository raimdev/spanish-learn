# CLAUDE.md

## Project Overview

Spanish learning web application. Go backend (chi router, SQLite/sqlc), React frontend (Vite). Features: grammar lessons, vocabulary quizzes, exercises (fill-blank, translation, conjugation), Anki-style flashcards with SM-2 spaced repetition. Multi-user with session auth.

## Tech Stack

- **Backend**: Go 1.24+, chi/v5 router, mattn/go-sqlite3 (CGO), sqlc for type-safe SQL
- **Frontend**: React 18, Vite 6, React Router v6
- **Database**: SQLite with WAL mode
- **Auth**: Session tokens in HttpOnly cookies, bcrypt passwords

## Repository Structure

```
backend/                  Go backend
  cmd/server/main.go      Entry point, migration runner
  db/migrations/           SQL schema + seed data (001-008)
  db/queries/              sqlc query definitions
  db/sqlc/                 Generated code (DO NOT EDIT)
  internal/handler/        HTTP handlers
  internal/middleware/      Auth middleware
  internal/service/        Business logic (auth, SM-2)
  internal/router/         Chi router setup
  sqlc.yaml                sqlc configuration
frontend/                 React frontend
  src/api/client.js        Fetch wrapper with credentials
  src/context/             Auth state (React Context)
  src/components/          Layout, ProtectedRoute
  src/pages/               All page components
  src/styles/              CSS
docker-compose.yml        Containerized deployment
```

## Development Commands

### Backend
```bash
cd backend
CGO_ENABLED=1 go build ./cmd/server/     # Build
CGO_ENABLED=1 go test ./...              # Run all tests
CGO_ENABLED=1 go test ./internal/...     # Run handler/service tests only
DB_PATH=data/spanish_learn.db go run ./cmd/server/  # Run server (port 8080)
```

### Frontend
```bash
cd frontend
npm install        # Install dependencies
npm run dev        # Dev server (port 5173, proxies /api to :8080)
npm run build      # Production build
```

### Docker
```bash
docker-compose up --build    # Start both services
# Backend: http://localhost:8080
# Frontend: http://localhost:5173
```

### sqlc (regenerate after changing queries)
```bash
cd backend
sqlc generate
```

## Test User

Login: `test` / `test` (pre-seeded via 008_seed_data.sql)

## Key Architecture Decisions

- **CGO_ENABLED=1** required for mattn/go-sqlite3
- Backend tests use in-memory SQLite (`:memory:`) with migrations applied per test
- sqlc generates `db/sqlc/` — never edit those files directly
- Frontend uses Vite dev proxy for `/api` → backend, no CORS needed in dev
- SM-2 algorithm in `internal/service/srs.go` — quality 0-5, same as Anki
- Session tokens are 64-char hex strings, 24h expiry, stored in DB

## Environment Variables

| Variable | Default | Description |
|---|---|---|
| DB_PATH | data/spanish_learn.db | SQLite database file path |
| PORT | 8080 | Backend server port |
| CORS_ORIGIN | http://localhost:5173 | Allowed CORS origin |
| MIGRATIONS_DIR | db/migrations | SQL migration files directory |
| VITE_API_URL | http://localhost:8080 | Backend URL for Vite proxy |
