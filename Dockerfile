# Use the latest stable Go image
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o task-manager .

# Use a minimal base image to run the application
FROM gcr.io/distroless/base

# Copy the binary from the builder image
COPY --from=builder /app/task-manager /task-manager

# Command to run the binary
CMD ["/task-manager"]
