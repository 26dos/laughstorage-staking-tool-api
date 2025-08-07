package proposals

import (
	"context"

	"api/api/proposals/manage"
	"api/internal/logic"
)

func (c *ControllerManage) AuditProposal(ctx context.Context, req *manage.AuditProposalReq) (res *manage.AuditProposalRes, err error) {
	proposal, err := logic.Proposals.AuditProposal(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &manage.AuditProposalRes{}
	res.Data = proposal
	return res, nil
}

func (c *ControllerManage) Rejection(ctx context.Context, req *manage.RejectionReq) (res *manage.RejectionRes, err error) {
	proposal, err := logic.Proposals.RejectionProposal(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &manage.RejectionRes{}
	res.Data = proposal
	return res, nil
}

func (c *ControllerManage) CreatePlan(ctx context.Context, req *manage.CreatePlanReq) (res *manage.CreatePlanRes, err error) {
	proposal, err := logic.Proposals.CreatePlan(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &manage.CreatePlanRes{}
	res.Data = proposal
	return res, nil
}
