package auth

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ProposalsReq struct {
	g.Meta         `path:"/" method:"post" tags:"proposals" summary:"proposals"`
	PId            string `json:"pId"`
	Data           string `json:"data" v:"required#Please enter the form data"`
	SubmitType     string `json:"submit_type" v:"required#Please enter the submit type"`
	PName          string `json:"p_name" v:"required#Please enter the data owner name"`
	RequestDataCap string `json:"request_data_cap" v:"required#The total amount of DataCap you need to apply for must be filled in"`
}

type ProposalsRes struct {
	model.ApiRes
}

type MyProposalsReq struct {
	g.Meta `path:"/my" method:"get" tags:"proposals" summary:"proposals"`
}

type MyProposalsRes struct {
	model.ApiRes
}
