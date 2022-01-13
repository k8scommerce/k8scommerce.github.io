---
title: "Development Philosophy"
weight: 2
date: 2022-01-10
description: >
  An overview of K8sCommerce's Development Philosophy
---

The K8sCommerce project embraces the development philosophy: *"Prefer tools over conventions."*

Employing code-generation tools as often as possible reduces repetitively writing foundational code such as CRUD operations, gRPC bindings, microservice boilerplate code. etc. This methodology shortens the developer's learning curve while reducing, and oftentimes eliminating, human-introduced bugs.

Incorporating the following code-generation tools in our project has helped speed up development, integrations, and logic refinements, creating a desirable, streamlined, high-velocity development approach.

- [Google's protocol buffers (protoc)](https://developers.google.com/protocol-buffers/): RPC client/server generation from .proto files with almost every language supported
- [go-zero's goctl](https://github.com/zeromicro/go-zero): microservice framework generation from .proto and .api files
- [xo's xo](https://github.com/xo/xo): Database CRUD object generation for relational databases, with primary key, foreign key, enum, and advanced composite primary key support.

Preferring tools over conventions, combined with ever-growing documentation, empowers businesses to successfully integrate K8sCommerce's projects into their existing coding ecosystem in condensed timeframes. 