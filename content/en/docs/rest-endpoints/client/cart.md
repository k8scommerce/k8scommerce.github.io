---
title: "Cart"
linkTitle: "Cart"
weight: 10
date: 2022-01-12
description: >
  Cart
---


### Creates a shopping cart for the associated customerId

#### Route definition

- Url: `/v1/cart/:customerId`
- Method: POST
- Request: `CreateCartRequest`
- Response: `CreateCartResponse`

#### Request definition


```golang
type CreateCartRequest struct {
	CustomerId string `path:"customerId"`
}
```


#### Response definition


```golang
type CreateCartResponse struct {
	Cart Cart `json:"cart"`
	StatusCode int64 `json:"sessionId"`
	StatusMessage string `json:"sessionId"`
}

type Cart struct {
	Item []Item `json:"items"`
	TotalPrice float64 `json:"totalPrice"`
}
```
  


### returns a shopping cart if one exists

#### Route definition

- Url: /v1/cart/:customerId
- Method: GET
- Request: `GetCartRequest`
- Response: `GetCartResponse`

#### Request definition


```golang
type GetCartRequest struct {
	CustomerId string `path:"customerId"`
}
```


#### Response definition


```golang
type GetCartResponse struct {
	Cart Cart `json:"cart"`
	StatusCode int64 `json:"sessionId"`
	StatusMessage string `json:"sessionId"`
}

type Cart struct {
	Item []Item `json:"items"`
	TotalPrice float64 `json:"totalPrice"`
}
```
  

