# Use the official Golang image as a build stage
FROM golang:1.19 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o receipt-processor .

# Use a compatible base image instead of Alpine (Alpine lacks some necessary libraries)
FROM debian:stable-slim

# Set working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/receipt-processor .

# Ensure the binary is executable
RUN chmod +x /root/receipt-processor

# Expose the port
EXPOSE 8080

# Start the application
CMD ["./receipt-processor"]
