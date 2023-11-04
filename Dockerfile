# Use an official Golang runtime as the base image
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image to run the application
FROM scratch

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Set the working directory inside the container
WORKDIR /app

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
