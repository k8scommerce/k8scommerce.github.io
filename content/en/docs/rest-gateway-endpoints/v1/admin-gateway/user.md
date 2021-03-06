---
title: "User"
linkTitle: "User"
weight: 110
date: 2022-01-12
description: >
  Content negotiation
---

## User Endpoints

###  admin

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/user/login | [login](#login) | Login |
  


## Paths

### <span id="login"></span> Login (*login*)

```
POST /v1/user/login
```

login for administration users

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [UserLoginRequest](#user-login-request) | `models.UserLoginRequest` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#login-200) | OK | A successful response. |  | [schema](#login-200-schema) |

#### Responses


##### <span id="login-200"></span> 200 - A successful response.
Status: OK

###### <span id="login-200-schema"></span> Schema
   
  

[UserLoginResponse](#user-login-response)

## Models

### <span id="jwt-token"></span> JwtToken


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accessExpire | int64 (formatted integer)| `int64` | ✓ | |  |  |
| accessToken | string| `string` | ✓ | |  |  |
| refreshAfter | int64 (formatted integer)| `int64` | ✓ | |  |  |



### <span id="response-status"></span> ResponseStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| statusCode | int64 (formatted integer)| `int64` | ✓ | | RFC http status code, ie. 204, etc - https:go.dev/src/net/http/status.go |  |
| statusMessage | string| `string` | ✓ | | status message |  |



### <span id="user"></span> User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | | email address |  |
| firstName | string| `string` | ✓ | | first name |  |
| id | int64 (formatted integer)| `int64` | ✓ | | user id |  |
| lastName | string| `string` | ✓ | | last name |  |
| password | string| `string` | ✓ | | password |  |



### <span id="user-login-request"></span> UserLoginRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | | email address |  |
| password | string| `string` | ✓ | | password |  |



### <span id="user-login-response"></span> UserLoginResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| jwt | [JwtToken](#jwt-token)| `JwtToken` | ✓ | | JwtToken object |  |
| status | [ResponseStatus](#response-status)| `ResponseStatus` | ✓ | | a ResponseStatus object |  |
| user | [User](#user)| `User` | ✓ | | User object |  |
