# Use the official GoLang base image
FROM golang:1.20.4

# Set the working directory inside the container
WORKDIR /target

# # Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
# Fetch dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the GoLang Web API; main.go in ./example
# CGO_ENABLED=1 GOOS=linux
RUN go build -o liftapp ./app

# Change ownership of the build files
RUN chmod +x liftapp

# Set the startup command for the container
CMD ["./liftapp"]
