// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"api/api/auth/v1"
)

type IAuthV1 interface {
	Nonce(ctx context.Context, req *v1.NonceReq) (res *v1.NonceRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
