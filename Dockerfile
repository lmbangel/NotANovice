# Build stage
FROM golang:1.24.6-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS and SQLite
RUN apk --no-cache add ca-certificates sqlite

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy any necessary files (migrations, etc.)
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/queries ./queries

# Create directory for SQLite database
RUN mkdir -p /root/data

# Expose port 8000
EXPOSE 8000

# Command to run
CMD ["./main"]