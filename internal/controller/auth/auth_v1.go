package auth

import (
	"context"

	v1 "api/api/auth/v1"
	"api/internal/logic"
)

func (c *ControllerV1) Nonce(ctx context.Context, req *v1.NonceReq) (res *v1.NonceRes, err error) {
	nonce, err := logic.Auth.Nonce(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.NonceRes{}
	res.Data = nonce
	return res, nil
}

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := logic.Auth.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.LoginRes{}
	res.Data = token
	return res, nil
}
