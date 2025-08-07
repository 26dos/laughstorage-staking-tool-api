package proposals

import (
	"context"

	"api/api/proposals/user"
	"api/internal/logic"
)

func (c *ControllerUser) Proposals(ctx context.Context, req *user.ProposalsReq) (res *user.ProposalsRes, err error) {
	proposals, err := logic.Proposals.GetProposals(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &user.ProposalsRes{}
	res.Data = *proposals
	return res, nil
}

func (c *ControllerUser) ProposalDetail(ctx context.Context, req *user.ProposalDetailReq) (res *user.ProposalDetailRes, err error) {
	proposal, err := logic.Proposals.GetProposalDetail(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &user.ProposalDetailRes{}
	res.Data = *proposal
	return res, nil
}
