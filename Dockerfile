# Build stage
FROM golang:1.20-alpine AS builder

# Install bash
RUN apk add --no-cache bash

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go .

# Build the Go application
RUN go build -o gobash main.go

# Runtime stage
FROM alpine:latest

# Install bash and Docker CLI
RUN apk add --no-cache bash docker-cli

# Set the working directory inside the container
WORKDIR /app

# Copy the Go executable from the builder stage
COPY --from=builder /app/gobash /app/gobash

# Set the entrypoint to run the Go application
ENTRYPOINT ["/app/gobash"]
