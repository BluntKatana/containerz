# Containerz - 𝕐
Practical assignment for the course: Software Containerization.

Authors: Sander van Werkhoven, Floris Bos, Maik Larooij (group 30)

# Contents
- [Project overview](#project-overview)
  - [Cluster overview](#cluster-overview)
- [Installation](#installation)
  - [Prequisites](#prerequisites)
  - [Install](#install)
  - [Scaling](#scaling)
  - [Uninstall](#uninstall)
  - [Video commands](#video-commands)
- [Upgrading](#upgrading)
  - [Rolling update](#rolling-update)
  - [Canary release](#canary-release)
- [Repo overview](#repo-overview)
  - [YAML files](#yaml-files)
  - [Dockerfiles](#dockerfiles)

## Project overview

This project provides a simple web application with a client, API and database. It is named 'Y', as a parody on 'X', the new Twitter. The client provides the user with a simple Web UI to read and post messages.

![screenshot](img/y-screenshot.png)

The client comunicates with the API to retrieve and post messages. The API is responsible for handling the POST, PUT and GET messages and interacting with the database. The database is a simple MySQL database. The picture below shows the connections between the different components.

![connections](img/connections.png)

The following technologies are used for building the application:
- Client: [svelte](https://svelte.dev/)
  - UI library: [shadcn-svelte](https://www.shadcn-svelte.com/)
- API server: [golang](https://go.dev/)
  - Web Framework: [echo](https://echo.labstack.com/)
- Database: [mysql](https://www.mysql.com/)

The application is containerized using Docker and deployed to a Kubernetes cluster using Helm. For the assignment, we deployed the application to a Google Kubernetes Engine (GKE) cluster. The service is publicly available through https://y.maiklarooij.nl (as long as the GKE cluster is running on the credits). 

### Cluster overview

To gain an understanding of the cluster, we present the following chart:

![overview](img/overview.png)

It shows that the API and client deployments both have 3 replicas running on different nodes (total 3 nodes). The database is not replicated. The API and client services are exposed through an ingress, which is secured using a certificate from our own created Certificate Authority (CA). The database is only accessible from the API pods.

The different Kubernetes objects and their relation to each other are further shown in the following diagram:

![flow](img/flow-diagram.png)

## Installation

### Prerequisites
- [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) `kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml`
- Namespaces: `kubectl create namespace y-dev && kubectl create namespace y-prod`
- A persistent volume claim for the database: `kubectl apply -f yaml-files/y-db-storage.yaml`

### Install
Edit `helm-chart/values.yaml` to you liking and then simply run `helm install y-dev ./helm-chart --namespace y-dev` to (for example) install the application in the development environment.

### Scaling
Examples:
- Edit `helm-chart/values.yaml` and change the `replicas` to 3 for one of the deployments and then run `helm upgrade y-dev ./helm-chart --namespace y-dev` to horizontally scale to 3 replicas.
- `kubectl scale deployment y-api-deployment --replicas=3 -n y-dev` to horizontally scale the API deployment to 3 replicas.
- Use an autoscaler: `kubectl autoscale deployment y-api-deployment --min=1 --max=3 --cpu-percent 60 -n y-dev` to horizontally scale the API deployment up to a maximum of 3 replicas when the average CPU utilization is above 60%.

### Uninstall
Run `helm uninstall y-dev --namespace y-dev` to uninstall the application.


### Video commands
All commands used in the video are listed in the [commands.txt](commands.txt) file.

## Upgrading
Run `helm upgrade y-dev ./helm-chart --namespace y-dev` to upgrade the application after making changes to the chart. In the video, we show upgrading a part of the application using a rolling update and the other part using a canary release. We will update the y-client deployment to [version 1.1.0](src/y-client-1.1.0/), which includes our mascot [Gopher](src/y-client-1.1.0/src/lib/assets/go-gopher.svg) to show in the bottom-left corner of the application.

We will describe how we updated our application in the video here.

### Rolling update
As the rolling update is the default strategy for Kubernetes deployments, we can use Helm for updating the client application. We simply replace the image tag for the y-client to be 1.1.0 and run `helm upgrade y-dev ./helm-chart --namespace y-dev` to upgrade the application. The rolling update will be performed automatically. We can track the progress of the rolling update by running `kubectl rollout status deployment y-client-deployment -n y-dev`. Afterwards, we can use `kubectl describe deployment y-client-deployment -n y-dev` to see the history of the deployment and see that the rolling update was performed, one pod at a time.

### Canary release
This one, we have to do manually. We create a new [deployment yaml file](yaml-files/y-client-deployment-v2.yaml). This one uses the new image tag, but uses the **same app label**, so the service will also distribute traffic to these pods. We apply this file by running `kubectl apply -f yaml-files/y-client-deployment-v2.yaml`. After verifying that the update was successful, we can scale the new (canary) deployment to 3 replicas and delete the old deployment. Now, the whole application is updated to the new version of the client.

## Repo overview
```
/
├── README.md
├── commands.txt                              # list of commands used in the video
├── src
│   ├── y-api-1.0.0                           # source code for the API (version 1.0.0)
│   ├── y-api-1.1.0                           # source code for the API (version 1.1.0)
│   ├── y-client                              # source code for the client (version 1.0.0)         
├── helm-chart                                # helm chart for the application
│   ├── Chart.yaml
│   ├── values.yaml                           # values for the helm chart  
│   └── templates                             # templates for the helm chart
│     ├── y-api-deployment.yaml               # deployment for the API
│     ├── y-api-service.yaml                  # service for the API
│     ├── y-ca-issuer.yaml                    # cert issuer for the application
│     ├── y-client-deployment.yaml            # deployment for the client
│     ├── y-client-service.yaml               # service for the client
│     ├── y-db-config.yaml                    # configmap for the database
│     ├── y-db-deployment.yaml                # deployment for the database
│     ├── y-db-init-config.yaml               # configmap for the database init script
│     ├── y-db-secret.yaml                    # secret for the database
│     ├── y-db-service.yaml                   # service for the database
│     ├── y-db-storage.yaml                   # persistent volume claim for the database
│     ├── y-ingress.yaml                      # ingress for the application
│     ├── y-issuer.yaml                       # issuer for the application
│     ├── y-network-policy.yaml               # network policy for the application
│     └── y-rbac.yaml                         # rbac for the application
├── yaml
    ├── video-yaml-files                      # extra yaml files used in the video
    │   ├── y-api-deployment-other.yaml       # api deployment with other label for network policy testing
    │   ├── y-client-deployment-v2.yaml       # deployment for the client (version 1.1.0) for upgrading
    │   └── y-db-storage.yaml                 # persistent volume claim for the database
    └── helm-raw-yaml                         # all yaml files used in helm chart without values
```

### YAML files
The following diagram shows the structure of the used yaml files:

![yaml-structure](img/file-structure.png)

The yaml files for Helm are found in the [helm-chart/templates](helm-chart/templates) folder. The raw yaml files used in the Helm chart are found in the [yaml](yaml/helm-raw-yaml) folder. The extra yaml files used in the video are found in the [video-yaml-files](yaml/video-yaml-files/) folder.

### Dockerfiles
All Dockerfiles used are in the [src](src/) folder.

- [y-api](src/y-api/Dockerfile)
- [y-client-1.0.0](src/y-client-1.0.0/Dockerfile)
- [y-client-1.1.0](src/y-client-1.1.0/Dockerfile)