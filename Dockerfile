# Multi-stage build: build the binary in a larger image
FROM golang:latest AS builder

# Set working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o main cmd/api/main.go

# Use the scratch image for the final image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Set port
ENV PORT=80

# Expose the port
EXPOSE 80

# Command to run the binary
CMD ["/main"]

