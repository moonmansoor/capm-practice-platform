# CAPM Practice Platform

Go-powered platform for Certified Associate in Project Management (CAPM) exam preparation. Includes realistic mock exams, hard drills, and targeted practice sections backed by PostgreSQL.

## Features
- Mock Exam (150 questions, PMI ECO aligned)
- Short Quiz, Hard Drill, PMP Mock
- Practice drills: Earned Value, PERT, Stakeholder Salience, Project/Program/Portfolio vs Operations, Team & Motivation Theories
- Narrative explanations, PDF reports, instant feedback

## Stack
- Go, Gorilla/Mux
- PostgreSQL (pgx)
- HTML/Bootstrap/Vanilla JS
- gofpdf for reports

## Quick Start
```bash
go mod tidy
export PG_URL=postgres://postgres:postgres@localhost:5432/capm?sslmode=disable
make db-migrate
make db-seed
make run
```

Open `http://localhost:8080` to access the UI.

## Structure
```
cmd/            Server, seed, migrate
internal/       Database, handlers, service, repository, pdf
web/            Templates, static assets
```

## API
- `POST /api/exams/start`
- `POST /api/hard/start`
- `GET /api/team-motivation/questions?count=20`
- `DELETE /api/attempts/{id}`

## Practice Pages
- `/earned-value-drill`
- `/pert-drill`
- `/stakeholder-salience`
- `/project-operations`
- `/team-motivation`

## Make Commands
- `make run`, `make build`
- `make db-migrate`, `make db-seed`
- `make clean`

## TODO
- Authentication & RBAC
- Hexagonal refactor
- Logging/metrics/tracing
- Unit & integration tests
- Graceful shutdown
- Secrets management
- Background job support

## License
MIT
