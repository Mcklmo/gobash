docker run \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v $(pwd):/app \
    -w /app \
    gobash-runner \
    docker.sh --children go.sh
