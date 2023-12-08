# Use the golang image version 1.21.3 based on Alpine
FROM golang:1.21.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local files to the container's working directory
COPY . /app

# Set environment variable for DB_URI
ENV DB_URI="mongodb+srv://admin:admin123@taller-admins.ez0xrnf.mongodb.net/?retryWrites=true&w=majority"

# Expose port 8080
EXPOSE 8080

# Command to run the Go application
CMD ["go", "run", "controller/main.go"]
