.PHONY: run stop restart db-migrate db-seed build clean

# Set default PostgreSQL URL if not provided
PG_URL ?= postgres://postgres:postgres@localhost:5432/capm?sslmode=disable

start:
	go run ./cmd/server

stop:
	@SERVER_PID=$$(pgrep -f "go run ./cmd/server" | head -n 1); \
	if [ -n "$$SERVER_PID" ]; then \
		echo "Stopping server (PID $$SERVER_PID)"; \
		kill $$SERVER_PID; \
	else \
		echo "Server is not running via go run"; \
	fi; \
	if lsof -i :8080 >/dev/null 2>&1; then \
		PID=$$(lsof -t -i :8080 | head -n 1); \
		echo "Stopping process on port 8080 (PID $$PID)"; \
		kill $$PID; \
	fi

restart:
	$(MAKE) stop
	$(MAKE) run

build:
	go build -o bin/server ./cmd/server

db-migrate:
	go run ./cmd/migrate

db-seed:
	go run ./cmd/seed

clean:
	rm -rf bin/

test:
	go test ./...

dev: db-migrate db-seed run
