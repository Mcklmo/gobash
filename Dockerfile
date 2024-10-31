# Use debian as base image since it's suitable for Docker-in-Docker
FROM debian:bullseye-slim

# Install required packages including Docker and dependencies
RUN apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    bash \
    docker.io \
    git \
    wget \
    --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

# Install Go
ENV GO_VERSION 1.23.2
RUN wget https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
    rm go$GO_VERSION.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Create a directory for the app
WORKDIR /app

# Copy the Go program and modules into the image
COPY main.go go.mod /app/

# Build the Go program
RUN go build -o /usr/local/bin/gobash /app/main.go

# Ensure the docker socket is accessible when running the container
VOLUME /var/run/docker.sock

# Make the container executable
ENTRYPOINT ["/usr/local/bin/gobash"]

# Default command (can be overridden when running the container)
CMD ["--help"]
