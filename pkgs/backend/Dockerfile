# ---------- Build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# ---------- Runtime stage ----------
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

# Expose port (change if your app uses another port)
EXPOSE 7007

# Run binary
CMD ["./app"]