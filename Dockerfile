# Build stage
ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Final stage
FROM debian:bookworm-slim

# Install CA certificates so TLS works with MongoDB Atlas
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binary
COPY --from=builder /app/server /usr/local/bin/server

# Expose app port
EXPOSE 8080

# Start app
CMD ["server"]