# Use an official Golang runtime as a parent image
FROM golang:1.21.3-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any dependencies
RUN go mod download

# Build the Go application
# RUN go build -o main ./...

# Expose port 8080 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["go", "run", "controller/main.go"]
