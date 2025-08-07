package v1

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type NonceReq struct {
	g.Meta `path:"/nonce" method:"get" tags:"auth" summary:"nonce"`
}

type NonceRes struct {
	model.ApiRes
}

type LoginReq struct {
	g.Meta    `path:"/login" method:"post" tags:"auth" summary:"login"`
	Nonce     string `v:"required" dc:"nonce"`
	Message   string `v:"required" dc:"message"`
	Signature string `v:"required" dc:"signature"`
	Address   string `v:"required" dc:"address"`
}

type LoginRes struct {
	model.ApiRes
}
