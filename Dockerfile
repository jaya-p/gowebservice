# Build: docker build -t gowebservice . 
#         && docker images
# Run: docker run -p 8080:8080 gowebservice // docker run -p <host-port>:<container-port> <image-name>
# Test: curl http://localhost:8080/api/v1/helloworld

# Use golang-based image for container; golang version 1.12.4
FROM golang:1.12.4-alpine

# Enable go modules for downloading dependencies
ENV GO111MODULE=on

# Add git executable in container
RUN apk add --no-cache git

# Copy webserver go file in local computer to container
COPY ./example/3-restapiwebserver/mainrestapi.go /go/src/

# Set working directory in container
WORKDIR /go

# Build the application
RUN go build -o bin/main src/mainrestapi.go

# Expose listening port for application
EXPOSE 8080

# Run the application
ENTRYPOINT ["/go/bin/main"]
