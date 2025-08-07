package model

import "github.com/golang-jwt/jwt/v5"

type ApiRes struct {
	Code    int         `json:"code" dc:"code"`
	Message string      `json:"message" dc:"message"`
	Data    interface{} `json:"data" dc:"data"`
	Paged   *Paged      `json:"paged,omitempty" dc:"paged"`
}

type PagedData struct {
	Paged Paged       `json:"paged" dc:"paged"`
	Data  interface{} `json:"data" dc:"data"`
}

type Paged struct {
	Total     int `json:"total" dc:"total"`
	Page      int `json:"page" dc:"page"`
	Size      int `json:"size" dc:"size"`
	TotalPage int `json:"total_page" dc:"total_page"`
}

type PageReq struct {
	Limit int    `json:"limit" d:"10" `
	Page  int    `json:"page"  d:"1" `
	Sort  string `json:"sort" dc:"排序方式" form:"sort"  d:"DESC"`
	Order string `json:"order" dc:"排序依据" form:"order" d:"created_at"`
}

type UserJwtClaims struct {
	Uid string `json:"uid" dc:"id"`
	jwt.RegisteredClaims
}
