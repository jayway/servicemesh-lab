# Service mesh on Kubernetes lab
This repo contains files for deploying microservices to a Kubernetes cluster.

There are two microservices, `numbergen` and `namegen`.

## Compile
```bash
CGO_ENABLED=0 GOOS=linux go build -a -o ./namegen/app ./namegen/
CGO_ENABLED=0 GOOS=linux go build -a -o ./numbergen/app ./numbergen/
```

## Build docker images

### Install docker machine
Required if you use a Cloud Shell.

#### Azure
```bash
DOCKER_RESOURCEGROUP_NAME="sandbox-<firstname>.<lastname>"
docker-machine create \
 --driver azure \
 --azure-location ${ACC_LOCATION} \
 --azure-resource-group ${DOCKER_RESOURCEGROUP_NAME} \
 --azure-subscription-id "$(az account show --query id -o tsv)" \
servicemesh-dockerhost
```
The first time you try to create a machine, Azure driver will ask you to authenticate.
If the host already exist, remove it with `$ docker-machine rm servicemesh-dockerhost`
List docker machines with `$ docker-machine ls`

Configure the shell
```bash
eval $(docker-machine env servicemesh-dockerhost --shell bash)
```
### Build
```bash
docker build -t namegen-scratch ./namegen/
docker build -t numbergen-scratch ./numbergen/
```

## Test
If you use Azure, make sure you have an inboud rule for port 8080, since the Cloud Shell is not in the same vnet.
To list existing rules: `$ az network nsg rule list --nsg-name servicemesh-dockerhost-firewall -o table`
To create a rule:

```bash
az network nsg rule create \
 --nsg-name servicemesh-dockerhost-firewall \
 --name Port8080 \
 --priority 500 \
 --destination-port-ranges 8080
```

```bash
docker run -d -p 8080:8080 namegen-scratch
curl $(docker-machine ip servicemesh-dockerhost):81
```

## Troubleshooting
Problem: You get `Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock:`.
Solution: Current user does not have permission. Run command with `sudo`.