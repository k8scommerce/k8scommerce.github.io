---
title: "Catalog"
linkTitle: "Catalog"
weight: 20
date: 2022-01-12
description: >
  Catalog
---


### returns all categories

#### Route definition

- Url: /v1/categories
- Method: GET
- Request: `GetAllCategoriesRequest`
- Response: `GetAllCategoriesResponse`

#### Request definition


```golang
type GetAllCategoriesRequest struct {
}
```


#### Response definition


```golang
type GetAllCategoriesResponse struct {
	Categories []Category `json:"categories,omitempty"`
	StatusCode int64 `json:"statusCode,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}
```
  


### returns a category by url slug

#### Route definition

- Url: /v1/category/slug/:slug
- Method: GET
- Request: `GetCategoryBySlugRequest`
- Response: `GetCategoryBySlugResponse`

#### Request definition


```golang
type GetCategoryBySlugRequest struct {
	Slug string `json:"slug,omitempty"`
}
```


#### Response definition


```golang
type GetCategoryBySlugResponse struct {
	Category Category `json:"category,omitempty"`
	StatusCode int64 `json:"statusCode,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}

type Category struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Description string `json:"description"`
	MetaTitle string `json:"metaTitle"`
	MetaDescription string `json:"metaDescription"`
	MetaKeywords string `json:"metaKeywords"`
	SortOrder int32 `json:"sortOrder"`
}
```
  


### returns a category by id

#### Route definition

- Url: /v1/category/:id
- Method: GET
- Request: `GetCategoryByIdRequest`
- Response: `GetCategoryByIdResponse`

#### Request definition


```golang
type GetCategoryByIdRequest struct {
	Id int64 `json:"id,omitempty"`
}
```


#### Response definition


```golang
type GetCategoryByIdResponse struct {
	Category Category `json:"category,omitempty"`
	StatusCode int64 `json:"statusCode,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}

type Category struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Description string `json:"description"`
	MetaTitle string `json:"metaTitle"`
	MetaDescription string `json:"metaDescription"`
	MetaKeywords string `json:"metaKeywords"`
	SortOrder int32 `json:"sortOrder"`
}
```
  


### returns a product by sku

#### Route definition

- Url: /v1/product/sku/:sku
- Method: GET
- Request: `GetProductBySkuRequest`
- Response: `Product`

#### Request definition


```golang
type GetProductBySkuRequest struct {
	Sku string `path:"sku"`
}
```


#### Response definition


```golang
type Product struct {
	Id int64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	Description string `json:"description"`
	MetaTitle string `json:"metaTitle,omitempty"`
	MetaDescription string `json:"metaDescription,omitempty"`
	MetaKeywords string `json:"metaKeywords,omitempty"`
	Variants []Variant `json:"variants,omitempty"`
}
```
  


### returns a product by url slug

#### Route definition

- Url: /v1/product/slug/:slug
- Method: GET
- Request: `GetProductBySlugRequest`
- Response: `Product`

#### Request definition


```golang
type GetProductBySlugRequest struct {
	Slug string `path:"slug"`
}
```


#### Response definition


```golang
type Product struct {
	Id int64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	Description string `json:"description"`
	MetaTitle string `json:"metaTitle,omitempty"`
	MetaDescription string `json:"metaDescription,omitempty"`
	MetaKeywords string `json:"metaKeywords,omitempty"`
	Variants []Variant `json:"variants,omitempty"`
}
```
  


### returns a product by id

#### Route definition

- Url: /v1/product/:id
- Method: GET
- Request: `GetProductByIdRequest`
- Response: `Product`

#### Request definition


```golang
type GetProductByIdRequest struct {
	Id int64 `path:"id"`
}
```


#### Response definition


```golang
type Product struct {
	Id int64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	Description string `json:"description"`
	MetaTitle string `json:"metaTitle,omitempty"`
	MetaDescription string `json:"metaDescription,omitempty"`
	MetaKeywords string `json:"metaKeywords,omitempty"`
	Variants []Variant `json:"variants,omitempty"`
}
```
  


### returns all products belonging to a category id

#### Route definition

- Url: /v1/products/category/:categoryId/:currentPage/:pageSize
- Method: GET
- Request: `GetProductsByCategoryIdRequest`
- Response: `GetProductsByCategoryIdResponse`

#### Request definition


```golang
type GetProductsByCategoryIdRequest struct {
	CategoryId int64 `path:"categoryId"`
	CurrentPage int64 `path:"currentPage"`
	PageSize int64 `path:"pageSize"`
	SortOn string `form:"sortOn,optional"`
}
```


#### Response definition


```golang
type GetProductsByCategoryIdResponse struct {
	Products []Product `json:"products"`
	TotalRecords int64 `json:"totalRecords"`
	TotalPages int64 `json:"totalPages"`
}
```
  


### returns all products

#### Route definition

- Url: /v1/products/:currentPage/:pageSize
- Method: GET
- Request: `GetAllProductsRequest`
- Response: `GetAllProductsResponse`

#### Request definition


```golang
type GetAllProductsRequest struct {
	CurrentPage int64 `path:"currentPage"`
	PageSize int64 `path:"pageSize"`
	SortOn string `form:"sortOn"`
}
```


#### Response definition


```golang
type GetAllProductsResponse struct {
	Products []Product `json:"products"`
	TotalRecords int64 `json:"totalRecords"`
	TotalPages int64 `json:"totalPages"`
}
```
  