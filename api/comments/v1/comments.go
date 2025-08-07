package v1

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)


type ListCommentsReq struct {
	g.Meta `path:"/{p_id}" method:"GET" tags:"comments" summary:"list comments"`
	PId    string `v:"required#proposal id is required" dc:"post id" in:"path"`
	model.PageReq
}

type ListCommentsRes struct {
	model.ApiRes
}

