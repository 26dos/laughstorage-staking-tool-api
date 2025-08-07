// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package proposals

import (
	"context"

	"api/api/proposals/auth"
	"api/api/proposals/manage"
	"api/api/proposals/user"
)

type IProposalsAuth interface {
	Proposals(ctx context.Context, req *auth.ProposalsReq) (res *auth.ProposalsRes, err error)
	MyProposals(ctx context.Context, req *auth.MyProposalsReq) (res *auth.MyProposalsRes, err error)
}

type IProposalsManage interface {
	AuditProposal(ctx context.Context, req *manage.AuditProposalReq) (res *manage.AuditProposalRes, err error)
	Rejection(ctx context.Context, req *manage.RejectionReq) (res *manage.RejectionRes, err error)
	CreatePlan(ctx context.Context, req *manage.CreatePlanReq) (res *manage.CreatePlanRes, err error)
}

type IProposalsUser interface {
	Proposals(ctx context.Context, req *user.ProposalsReq) (res *user.ProposalsRes, err error)
	ProposalDetail(ctx context.Context, req *user.ProposalDetailReq) (res *user.ProposalDetailRes, err error)
}
