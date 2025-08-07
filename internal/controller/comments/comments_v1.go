package comments

import (
	"context"

	v1 "api/api/comments/v1"
	"api/internal/logic"
)



func (c *ControllerV1) ListComments(ctx context.Context, req *v1.ListCommentsReq) (res *v1.ListCommentsRes, err error) {
	comments, err := logic.Comments.ListComments(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.ListCommentsRes{}
	res.Data = comments
	return res, nil
}
