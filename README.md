# otusk8s2

## Requirements

* Docker 18.06 or later
* Kubernetes 1.19.7 or later
    * NGINX Ingress Controller 0.45.0 or later

## How to run with Docker

Build a Docker image:

```
make build-docker
```

It makes an ELF-binary and prepares a Docker image `adalekin/otusk8s2:latest` based on Scratch.

For more information see the [Dockerfile](build/docker/Dockerfile).

## How to run with Kubernetes

Apply Kubernetes manifests:

```
kubectl apply -f deployments/k8s
```

**NOTE:** you don't need to build a Docker image cause it already published at Docker Hub (https://hub.docker.com/repository/docker/adalekin/otusk8s2/).
