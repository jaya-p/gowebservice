# Build: docker build -t helloworldweb . 
#         && docker images
# Run: docker run -p 8080:8080 helloworldweb // docker run -p <host-port>:<container-port> <image-name>
# Test: curl http://localhost:8080/helloworld

# Use golang-based image for container; golang version 1.12.4
FROM golang:1.12.4-alpine

# Enable go modules for downloading dependencies
ENV GO111MODULE=on

# Add git executable in container
RUN apk add --no-cache git

# Copy webserver go file in local computer to container
COPY ./example/mainWeb.go /go/src/

# Set working directory in container
WORKDIR /go

# Build the application
RUN go build -o bin/main src/mainWeb.go

# Expose listening port for application
EXPOSE 8080

# Run the application
CMD ["/go/bin/main"]
