# Use the official Golang image as the base image
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and code into the container
COPY go.mod go.sum ./
COPY main.go ./

# Download the Go module dependencies
RUN go mod download

# Build the Go CLI application
RUN go build -o toolbox-cli

# Create a smaller, final image
FROM alpine:latest

# Set the working directory inside the final image
WORKDIR /app

# Copy the built binary from the previous stage into the final image
COPY --from=build /app/toolbox-cli ./toolbox-cli

# Expose any necessary ports (if your CLI application uses network ports)
# EXPOSE 8080

# Define the command to run your CLI application
CMD ["./toolbox-cli"]
