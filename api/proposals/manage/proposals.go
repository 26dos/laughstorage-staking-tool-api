package manage

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/shopspring/decimal"
)

type AuditProposalReq struct {
	g.Meta  `path:"/{pid}/proposals-audit" method:"post" tags:"Proposals" summary:"Audit Proposal"`
	PID     string `json:"pid" in:"path" v:"required#proposal id is required"`
	DataCap string `json:"data_cap" v:"required#data cap is required"`
}

type AuditProposalRes struct {
	model.ApiRes
}

type RejectionReq struct {
	g.Meta `path:"/{pid}/proposals-rejection" method:"post" tags:"Proposals" summary:"Rejection Proposal"`
	PID    string `json:"pid" in:"path" v:"required#proposal id is required"`
	Reason string `json:"reason" v:"required#reason is required"`
}
type RejectionRes struct {
	model.ApiRes
}

type CreatePlanReq struct {
	g.Meta    `path:"/{pid}/create-plan" method:"post" tags:"Proposals" summary:"Add Plan"`
	PID       string          `json:"pid" in:"path" v:"required#proposal id is required"`
	DataCap   decimal.Decimal `json:"data_cap" v:"required#data cap is required"`
	Amount    decimal.Decimal `json:"amount" v:"required#amount is required"`
	StakeDays int             `json:"stake_days" v:"required#stake days is required"`
}

type CreatePlanRes struct {
	model.ApiRes
}
