---
title: "Customer"
linkTitle: "Customer"
weight: 30
date: 2022-01-12
description: >
  Customer
---


### 1. customer endpoints

#### Route definition

- Url: /v1/customer/login
- Method: POST
- Request: `CustomerLoginRequest`
- Response: `CustomerLoginResponse`

#### Request definition

```golang
type CustomerLoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
```


#### Response definition

```golang
type CustomerLoginResponse struct {
	JwtToken JwtToken `json:"jwt"`
	Customer Customer `json:"customer"`
	StatusCode int64 `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

type JwtToken struct {
	AccessToken string `json:"access_token"`
	AccessExpire int64 `json:"access_expire"`
	RefreshAfter int64 `json:"refresh_after"`
}

type Customer struct {
	Id int64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password,omitempty"`
}
```
  

