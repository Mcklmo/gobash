GO_VERSION=1.23
echo "Running Go program with Docker. Go version: $GO_VERSION"
ls
docker run --rm -v "$(pwd)":/app -w /app golang:$GO_VERSION go run test/main.go
