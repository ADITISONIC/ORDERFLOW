# Build Stage
FROM golang:1.26.5-alpine AS builder

WORKDIR /app

# Copy go module files
COPY go.mod go.sum ./

RUN go mod download

# Copy source code
COPY . .

# Build application
RUN go build -o orderflow ./cmd/api

# Runtime Stage
FROM alpine:latest

WORKDIR /app

# Copy binary
COPY --from=builder /app/orderflow .

EXPOSE 8080

CMD ["./orderflow"]