# Use the official Go image as base
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/socialai ./main.go

# Use a minimal image for runtime
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/socialai .

# Copy configuration files
COPY conf/ ./conf/

# Expose port (App Engine will set PORT env var)
EXPOSE 8080

# Run the application
CMD ["./socialai"]


