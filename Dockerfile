# Step 1: Build the Go application
# Use an official Golang image as the build environment
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Step 2: Create a minimal final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8000 for the HTTP server
EXPOSE 8080

# Run the Go application binary
CMD ["./main"]