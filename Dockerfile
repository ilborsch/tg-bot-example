# Use the official Golang image as a build stage
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

COPY ./config/prod.yaml ./config/prod.yaml

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tg-bot ./cmd/tg-bot/main.go

# Use a smaller base image for the final stage
FROM alpine:latest

# Install necessary libraries (if needed)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/tg-bot .

# Copy the configuration file from the builder stage
COPY --from=builder /app/config/prod.yaml ./config/prod.yaml

# Ensure the binary has executable permissions
RUN chmod +x ./tg-bot

# Command to run the application
CMD ["./tg-bot", "--config=./config/prod.yaml"]
