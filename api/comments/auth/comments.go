package auth

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type PubCommentsReq struct {
	g.Meta   `path:"/{p_id}" method:"POST" tags:"comments" summary:"publish comment"`
	PId      string `v:"required#proposal id is required" dc:"post id" in:"path"`
	Content  string `v:"required#comment content is required" dc:"comment content" json:"content"`
	ParentId int64  `dc:"parent id" json:"parent_id" in:"query"`
}

type PubCommentsRes struct {
	model.ApiRes
}

type DelCommentsReq struct {
	g.Meta `path:"/{c_id}" method:"DELETE" tags:"comments" summary:"delete comment"`
	CId    int64 `v:"required#comment id is required" dc:"comment id" in:"path"`
}

type DelCommentsRes struct {
	model.ApiRes
}
