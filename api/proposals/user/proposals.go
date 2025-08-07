package user

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ProposalsReq struct {
	g.Meta `path:"/" method:"get" tags:"proposals" summary:"proposals"`
	model.PageReq
}

type ProposalsRes struct {
	model.ApiRes
}

type ProposalDetailReq struct {
	g.Meta `path:"/{p_id}" method:"get" tags:"proposals" summary:"proposals"`
	PId    string `json:"p_id" v:"required#Please enter the id"`
}

type ProposalDetailRes struct {
	model.ApiRes
}
