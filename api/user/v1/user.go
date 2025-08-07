package v1

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta `path:"/" method:"get" tags:"user" summary:"user"`
}

type UserRes struct {
	g.Meta `mime:"application/json"`
	model.ApiRes
}

type KycVerifyReq struct {
	g.Meta       `path:"/kyc-verify" method:"post" tags:"user" summary:"kyc verify"`
	DocumentType string `json:"document_type" v:"required#Document type is required"`
	IdFirstName  string `json:"id_first_name" v:"required#Id first name is required"`
	IdLastName   string `json:"id_last_name" v:"required#Id last name is required"`
	IdNumber     string `json:"id_number" v:"required#Id number is required"`
	IdCardFront  string `json:"id_card_front" `
	IdCardBack   string `json:"id_card_back" `
}

type KycVerifyRes struct {
	g.Meta `mime:"application/json"`
	model.ApiRes
}
