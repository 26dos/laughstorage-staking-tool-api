// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"api/api/user/v1"
)

type IUserV1 interface {
	User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error)
	KycVerify(ctx context.Context, req *v1.KycVerifyReq) (res *v1.KycVerifyRes, err error)
}
