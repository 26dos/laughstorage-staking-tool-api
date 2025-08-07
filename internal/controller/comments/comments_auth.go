package comments

import (
	"api/api/comments/auth"
	"api/internal/logic"
	"context"
)

func (c *ControllerAuth) DelComments(ctx context.Context, req *auth.DelCommentsReq) (res *auth.DelCommentsRes, err error) {
	err = logic.Comments.DelComments(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &auth.DelCommentsRes{}
	return res, nil
}

func (c *ControllerAuth) PubComments(ctx context.Context, req *auth.PubCommentsReq) (res *auth.PubCommentsRes, err error) {
	err = logic.Comments.PubComments(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &auth.PubCommentsRes{}
	return res, nil
}
