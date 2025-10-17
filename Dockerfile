FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o server ./cmd/server
RUN go build -o migrate ./cmd/migrate
RUN go build -o seed ./cmd/seed

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binaries
COPY --from=builder /app/server .
COPY --from=builder /app/migrate .
COPY --from=builder /app/seed .

# Copy web assets
COPY --from=builder /app/web ./web

EXPOSE 8080

CMD ["./server"]
