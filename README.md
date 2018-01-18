# Service mesh on Kubernetes lab
This repo contains files for deploying microservices to a Kubernetes cluster.

There are two microservices, `numbergen` and `namegen`.

## Compile and build docker images
```bash
CGO_ENABLED=0 GOOS=linux go build -a -o ./namegen/app ./namegen/
docker build -t <name>-scratch .

CGO_ENABLED=0 GOOS=linux go build -a -o ./numbergen/app ./numbergen/
docker build -t <name>-scratch .
```
