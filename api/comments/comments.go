// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package comments

import (
	"context"

	"api/api/comments/auth"
	"api/api/comments/v1"
)

type ICommentsAuth interface {
	PubComments(ctx context.Context, req *auth.PubCommentsReq) (res *auth.PubCommentsRes, err error)
	DelComments(ctx context.Context, req *auth.DelCommentsReq) (res *auth.DelCommentsRes, err error)
}

type ICommentsV1 interface {
	ListComments(ctx context.Context, req *v1.ListCommentsReq) (res *v1.ListCommentsRes, err error)
}
