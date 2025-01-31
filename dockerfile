# Step 1: Build the Go application
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download the dependencies first
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application binary
RUN go build -o myapp .

# Step 2: Create a smaller image for the final container
FROM alpine:latest

# Install CA certificates to allow HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port the Go app is running on
EXPOSE 8080

# Command to run the Go app inside the container
CMD ["./myapp"]
