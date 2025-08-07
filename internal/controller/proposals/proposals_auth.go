package proposals

import (
	"context"

	"api/api/proposals/auth"
	"api/internal/logic"
)

func (c *ControllerAuth) Proposals(ctx context.Context, req *auth.ProposalsReq) (res *auth.ProposalsRes, err error) {
	pid, err := logic.Proposals.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &auth.ProposalsRes{}
	res.Data = *pid
	return res, nil
}

func (c *ControllerAuth) MyProposals(ctx context.Context, req *auth.MyProposalsReq) (res *auth.MyProposalsRes, err error) {
	data, err := logic.Proposals.GetMyProposals(ctx)
	if err != nil {
		return nil, err
	}
	res = &auth.MyProposalsRes{}
	res.Data = *data
	return res, nil
}
