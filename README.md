# Gobash

A tool to run shell scripts with Docker.

It runs all children in parallel and waits for all of them to finish before exiting.

If any child script exits with a non-zero exit code, all other children and the entire script will exit immediately with the same exit code.

## Usage

```bash
PARENT_SCRIPT=docker.sh
CHILD_SCRIPT=go.sh

docker run \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v "$(pwd)":"$(pwd)" \
  -w "$(pwd)" \
  -e HOST_PWD="$(pwd)" \
  mcklmo/gobash \
  $PARENT_SCRIPT --children $CHILD_SCRIPT
```

## Deployment

```bash
docker build -t mcklmo/gobash .
docker push mcklmo/gobash
```
