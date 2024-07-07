# Use the official Go image as the base image
FROM golang:1.22.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# Copy the rest of the application
COPY . .

ENTRYPOINT CompileDaemon -build="go build -o app ./main.go" -command="./app"