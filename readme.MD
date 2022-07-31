# Kubernetes templates

# Run on minikube, with docker.

Minikube is local Kubernetes, focusing on making it easy to learn and develop for Kubernetes.

## Install:
- [docker](https://docs.docker.com/engine/install/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
- [minikube](https://minikube.sigs.k8s.io/docs/start/#installation)
  
NOTE: minikube and kubectl are cli tools that can be installed using your package manager, Docker Desktop is GUI, that comes with a cli tool.

## Run:
```
minikube start --driver="docker"
```
```
minikube dashboard
```

# Recommended vscode extensions:
- [Kubernetes](https://marketplace.visualstudio.com/items?itemName=ms-kubernetes-tools.vscode-kubernetes-tools)

# Useful links:
- [Kubernetes interactive tutorial](https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/cluster-interactive/)
- [AWS EKS workshop](https://www.eksworkshop.com/)
- [Terraform EKS module](https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest)