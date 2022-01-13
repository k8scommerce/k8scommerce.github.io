---
title: "What is K8sCommerce?"
linkTitle: "What is K8sCommerce?"
weight: 10
date: 2022-01-12
description: >
  An overview of K8sCommerce and why it exists
---


K8sCommerce is a headless ecommerce platform architected from the ground-up to be cloud native 
and run on Kubernetes. It is comprised of 14 microservices, 2 public/web facing via a JSON REST API,
and 12 private/non-web facing RPC services.

### JSON REST API Facing Services:
- There are 2 [Public-Facing JSON Services](/docs/microservices/public-facing-json). 
  - A [client](/docs/microservices/public-facing-json/client) service that exposes all necessary [client endpoints](/docs/rest-endpoints/client) to facilitate a headless storefront 
    - [Cart](/docs/rest-endpoints/client/cart)
    - [Customer](/docs/rest-endpoints/client/customer) 
    - [Catalog](/docs/rest-endpoints/client/catalog) 
    - [Payments](/docs/rest-endpoints/client/payment)
  - An [admin](/docs/microservices/public-facing-json/admin) service that exposes all [administration endpoints](/docs/rest-endpoints/admin) required to manage the K8sCommerce platform:
    - [Cart](/docs/rest-endpoints/admin/cart)
    - [Catalog](/docs/rest-endpoints/admin/catalog) 
    - [Customer](/docs/rest-endpoints/admin/customer) 
    - [Email](/docs/rest-endpoints/admin/email) 
    - [Inventory](/docs/rest-endpoints/admin/inventory) 
    - [Others Bought](/docs/rest-endpoints/admin/othersbought) 
    - [Payment](/docs/rest-endpoints/admin/payment) 
    - [Shipping](/docs/rest-endpoints/admin/shipping) 
    - [Similar Products](/docs/rest-endpoints/admin/similarproducts) 
    - [Store](/docs/rest-endpoints/admin/store) 
    - [User](/docs/rest-endpoints/admin/user) 
    - [Warehouse](/docs/rest-endpoints/admin/warehouse) 
  