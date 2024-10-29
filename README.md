# gobash

A simple tool to run multiple bash scripts in parallel and wait for all of them to finish.

## Usage

```bash
go run main.go parent1.sh parent2.sh --children child1.sh child2.sh
```

The first arguments are the parent scripts that will wait for the children to finish. At least one parent script is required.

`--children` is optional and can be used to specify multiple scripts that will be run in parallel.
