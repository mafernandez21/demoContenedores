# Fundamentos de Cloud Computing - AWS

## DEMO - CONTENEDORES

Este material complementa la clases del módulo 6, donde se vió el uso de contenedores, ECR y ECS

### Links útiles

- [VsCode](https://code.visualstudio.com/)
- [Instalación de Docker](https://docs.docker.com/engine/install/)
- [Docker Cheat Sheet](https://www.docker.com/wp-content/uploads/2022/03/docker-cheat-sheet.pdf)
- [Dockerfile Reference](https://docs.docker.com/engine/reference/builder/)
- [AWS-CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

## DEMO - EKS

### Autenticar kubectl en AWS

- aws eks update-kubeconfig --region us-east-1 --name MyClusteK8s

### Escalar pods

- kubectl scale --replicas=1 -f Deployment.yamlapiVersion: v1

### Aplicar manifiestos

- kubectl apply -f MANIFIESTO.YAML

### Borrar manifiestos

- kubectl delete -f MANIFIESTO.YAML

### Links  útiles

- [Conceptos de Kubernetes](https://kubernetes.io/docs/concepts/)
