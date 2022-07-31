# Hellokube
Hellokube is a simple demonstration on how to write services that communicate with each other. Build docker images, and deploy them to a Kubernetes cluster.
## Table of Contents
- [Hellokube](#hellokube)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [How to run](#how-to-run)
- [Project structure](#project-structure)
  - [./hello](#hello)
  - [./count](#count)
  - [Makefile](#makefile)
  - [kubernetes.yaml](#kubernetesyaml)
## Prerequisites
Running Kubernetes cluster, docker, kubectl.

## How to run
From this directory run following commands:

```
make docker-images
```

```
kubectl apply -f kubernetes.yaml
```

# Project structure
## ./hello
Service hello is exposed to the internet. With 1 public endpoint `[GET] http://hostname/hello/name`.

:point_right: Example usage 
```
curl http://hostname/hello/world
```
:point_right: Example response 
```
Hello world
```

Before returning response, hello service will call [count](#count) service and log response of the count service.

For simplicity - it will panic if it does not get a response from the count service.

## ./count
Service for internal use. With 1 private endpoint that is available only from the inside of the Kubernetes cluster. When called will return number of times it has been called since it started.

:point_right: Example usage 
```
curl http://count-service/count
```
:point_right: Example response 
```
Said hello 10 times
```
## Makefile
Contains commands to build docker images.

## kubernetes.yaml
Contains configurations for kubernetes cluster.