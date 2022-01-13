---
title: "Project Overview"
linkTitle: "Project Overview"
weight: 1
date: 2022-01-10
description: >
  This page tells you how to get started with K8sCommerce including deployment on Kubernetes and basic configuration.
---

## Service Overview
![K8sCommerce Overview](/docs/images/diagrams/K8sCommerceOverview-no-title.drawio.png)


## Current Status

The following shows what has been completed and what is yet to be done:

| Service          | DB Table(s)       | .proto           | .api             | logic            |  tests           |  release        |
| ---------------- |:-----------------:|:----------------:|:----------------:|:----------------:|:----------------:|:----------------:
| Cart             |<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|                  | alpha           |
| Catalog          |<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|                  | alpha           |
| Customer         |<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  | -                 |
| Email            | -                 |<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  |                  | -               |
| Inventory        |<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  |                  | -               |
| Others Bought    | -                 |<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  |                  | -               |
| Payment          |<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  |                  | -               |
| Shipping         | -                 |<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  |                  | -               |
| Similar Products |<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  |                  | -               |
| Store            |<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  |                  | -               |
| User             |<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  | alpha            |
| Warehouse        | -                |<i class="fas fa-check-square text-success fa-2x"></i>|                  |                  |                  | -               |



## Getting Started

K8sCommerce can be run in multiple ways:

- **Local single-node kubernetes kind cluster:** Recommended for evaluating K8sCommerce - [instructions]()

- **Deployment to your existing cluster:** Recommended for testing & production - [intructions]()

- **Running Clusterless:** Recommended for local development - [intructions]() 


### Quick Start
Deployment to a k8s cluster using the K8sCommerce Operator.

```sh
# Install RabbitMQ Operator for messaging
kubectl apply -f https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml

# Install K8sCommerce Operator
kubectl apply -f 
```

