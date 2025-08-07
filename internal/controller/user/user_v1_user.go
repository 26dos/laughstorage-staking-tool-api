package user

import (
	"context"

	v1 "api/api/user/v1"
	"api/internal/logic"
)

func (c *ControllerV1) User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	user, err := logic.User.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.UserRes{}
	res.Data = user
	return res, nil
}

func (c *ControllerV1) KycVerify(ctx context.Context, req *v1.KycVerifyReq) (res *v1.KycVerifyRes, err error) {
	ok, err := logic.User.KycVerify(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.KycVerifyRes{}
	res.Data = ok
	return res, nil
}
