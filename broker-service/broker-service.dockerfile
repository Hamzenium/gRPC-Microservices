# Use the official Golang image as a base image for building the application
# 'golang:1.18-alpine' is a lightweight version of Golang with Alpine Linux
FROM golang:1.18-alpine as builder

# Create a directory in the container for the application
RUN mkdir /app

# Copy the entire project directory from the host to the /app directory in the container
COPY . /app

# Set the working directory inside the container to /app
WORKDIR /app

# Build the Go application
# - CGO_ENABLED=0 disables CGO to ensure the build is fully static (no external dependencies)
# - The output binary is named 'brokerApp' and is built from the ./cmd/api directory
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# Ensure the built binary has executable permissions
RUN chmod +x /app/brokerApp

# Use a minimal base image for the final Docker image to reduce size
FROM alpine:latest

# Create a directory in the final image for the application
RUN mkdir /app

# Copy the built binary from the builder stage to the final image
COPY --from=builder /app/brokerApp /app

# Set the command to run the application when the container starts
CMD [ "/app/brokerApp" ]
