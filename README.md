# gobash

A simple tool that runs multiple bash scripts in parallel and waits for all of them to finish. It requires at least one parent script that will wait for the children to finish successfully.

## Usage

```bash
go run main.go parent1.sh parent2.sh --children child1.sh child2.sh evil_child.sh
```

The first arguments are the parent scripts that will wait for the children to finish. At least one parent script is required.

`--children` is optional and can be used to specify multiple scripts that will be run in parallel.

### Docker

```bash
docker run --rm \
  -v $(pwd):/scripts \
  -v /var/run/docker.sock:/var/run/docker.sock  mcklmo/gobash \
  /scripts/parent1.sh /scripts/parent2.sh --children /scripts/child1.sh /scripts/child2.sh
```

## Maintenance

### Push new version

```bash
docker build -t mcklmo/gobash .
docker push mcklmo/gobash
```
