FROM golang:1.16-alpine as builder

# Set destination for COPY
WORKDIR /console-application

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . /console-application/

# Build
RUN go build -o email-application /console-application/cmd/
RUN chmod +x email-application