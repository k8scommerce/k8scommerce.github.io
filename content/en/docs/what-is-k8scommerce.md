---
title: "What is K8sCommerce?"
linkTitle: "What is K8sCommerce?"
weight: 5
date: 2022-01-12
description: >
  An overview of what K8sCommerce is and why it exists
---

### K8sCommerce Platform

K8sCommerce is a headless ecommerce platform architected from the ground-up to be cloud native 
and run on Kubernetes. It is comprised of 14 microservices, 2 public/web facing via a JSON REST API,
and 12 private/non-web facing RPC services.

### Why K8sCommerce?

As of January 2022 no production ready open source projects written in Go and designed for Kubernetes existed.

Today's most popular open-source ecommerce solutions (WooCommerce - 30% market share, Magento - 9% market share, Bagisto, Drupal commerce, Joomla, OpenCart, Prestashop, OsCommerce, Zencart, X-Cart, etc.) are written in PHP, with many founded 20+ years ago, and are monolithic (all-in-one) software programs that were never designed for the cloud and lack finite resource management

In 2022 the time for a highly-scalable, open source, microservice based ecommerce platform was past due. We need systems
that are resource conscientious that automatically adapt in size based on traffic and computing loads. Kubernetes enables 
adapting / scaling in size but if the majority of open source software available are monolithic leviathans K8s auto scaling 
cannot break those application into separate services. 

### JSON REST Gateway Services

There are currently 2 public-facing, RESTful Gateway Services. These external-facing 
services are intended to be exposed to the internet via a DMZ network. These communicate
via `JSON` over common protocols such as `http` and `https` through standardized 
`GET`, `POST`, `PUT`, `PATCH`, and `DELETE` HTTP methods.

### Client Gateway Service

The client gateway service exposes all necessary client REST endpoints to facilitate a headless storefront.

  - Microservice overview: [client gateway](/docs/microservices/gateway-services/client-gateway)
  - REST endpoints: [client REST endpoints](/docs/rest-gateway-endpoints/client-gateway)

Client gateway endpoints by service:

  - [cart endpoints](/docs/rest-gateway-endpoints/client-gateway/cart)
  - [customer endpoints](/docs/rest-gateway-endpoints/client-gateway/customer) 
  - [catalog endpoints](/docs/rest-gateway-endpoints/client-gateway/catalog) 
  - [payments endpoints](/docs/rest-gateway-endpoints/client-gateway/payment)

### Admin Gateway Service

In addition to all client gateway endpoints, the admin gateway exposes all endpoints needed for full administration of the platform and store(s).

  - Microservice overview: [admin gateway](/docs/microservices/gateway-services/admin-gateway)
  - REST endpoints: [admin REST endpoints](/docs/rest-gateway-endpoints/admin-gateway)

Admin gateway endpoints by service:

  - [cart endpoints](/docs/rest-gateway-endpoints/admin-gateway/cart)
  - [catalog endpoints](/docs/rest-gateway-endpoints/admin-gateway/catalog) 
  - [customer endpoints](/docs/rest-gateway-endpoints/admin-gateway/customer) 
  - [email endpoints](/docs/rest-gateway-endpoints/admin-gateway/email) 
  - [inventory endpoints](/docs/rest-gateway-endpoints/admin-gateway/inventory) 
  - [others bought endpoints](/docs/rest-gateway-endpoints/admin-gateway/othersbought) 
  - [payment endpoints](/docs/rest-gateway-endpoints/admin-gateway/payment) 
  - [shipping endpoints](/docs/rest-gateway-endpoints/admin-gateway/shipping) 
  - [similar products endpoints](/docs/rest-gateway-endpoints/admin-gateway/similarproducts) 
  - [store endpoints](/docs/rest-gateway-endpoints/admin-gateway/store) 
  - [user endpoints](/docs/rest-gateway-endpoints/admin-gateway/user) 
  - [warehouse endpoints](/docs/rest-gateway-endpoints/admin-gateway/warehouse) 
  