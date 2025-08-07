package logic

import (
	commentsAuth "api/api/comments/auth"
	v1 "api/api/comments/v1"
	"api/internal/dao"
	"api/internal/model"
	"api/internal/model/entity"
	"api/internal/service"
	"api/utility"
	"context"
	"math"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var Comments = new(comments)

type comments struct {
}

func (c *comments) DelComments(ctx context.Context, req *commentsAuth.DelCommentsReq) error {
	comment := entity.SliverComments{}
	err := dao.SliverComments.Ctx(ctx).Where("ID", req.CId).Scan(&comment)
	if err != nil {
		return service.CustomError.NoData(ctx, "comment not found")
	}
	user, _ := service.User.GetCtxUser(ctx)
	if comment.WalletAddress != user.Wallet {
		return service.CustomError.NoData(ctx, "you are not the owner of the comment")
	}
	// transaction
	err = dao.SliverComments.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.SliverComments.Ctx(ctx).Where("ID", req.CId).Delete()
		if err != nil {
			return service.CustomError.ServerError(ctx, "delete comment failed")
		}
		_, err = dao.SliverComments.Ctx(ctx).Where("parent_id", req.CId).Delete()
		if err != nil {
			return service.CustomError.ServerError(ctx, "delete comment failed")
		}
		return nil
	})
	if err != nil {
		return service.CustomError.ServerError(ctx, "delete comment failed")
	}
	return nil
}

func (c *comments) PubComments(ctx context.Context, req *commentsAuth.PubCommentsReq) error {
	content := utility.StringSanitization(req.Content)
	contentLen := len(content)
	if contentLen < 15 || contentLen > 1000 {
		return service.CustomError.ParameterError(ctx, "comment content length must be between 15 and 1000")
	}
	p := entity.SliverProposals{}

	err := dao.SliverProposals.Ctx(ctx).Where("p_id", req.PId).Scan(&p)
	if err != nil {
		return service.CustomError.NoData(ctx, "proposal not found")
	}

	if req.ParentId != 0 {
		comment := entity.SliverComments{}
		err := dao.SliverComments.Ctx(ctx).Where(g.Map{
			"ID":   req.ParentId,
			"p_id": p.Id,
		}).Scan(&comment)
		if err != nil {
			return service.CustomError.NoData(ctx, "your reply to a comment that does not exist")
		}
	}

	user, _ := service.User.GetCtxUser(ctx)

	comment := model.Comment{
		WalletAddress:  user.Wallet,
		CommentContent: content,
		PId:            p.Id,
	}
	if req.ParentId != 0 {
		comment.ParentId = req.ParentId
	}

	_, err = dao.SliverComments.Ctx(ctx).Data(comment).OmitEmpty().Insert()
	if err != nil {
		return service.CustomError.ServerError(ctx, "publish comment failed")
	}
	return nil
}

func (c *comments) ListComments(ctx context.Context, req *v1.ListCommentsReq) (model.PagedData, error) {
	p := entity.SliverProposals{}
	err := dao.SliverProposals.Ctx(ctx).Where("p_id", req.PId).Scan(&p)
	if err != nil {
		return model.PagedData{}, service.CustomError.NoData(ctx, "proposal not found")
	}

	comments := []model.Comment{}
	total, err := dao.SliverComments.Ctx(ctx).Where("p_id", p.Id).Count()
	if err != nil {
		return model.PagedData{}, service.CustomError.ServerError(ctx, "list comments failed")
	}
	err = dao.SliverComments.Ctx(ctx).Where("p_id", p.Id).Order("created_at DESC").WithAll().Page(req.Page, req.Limit).Scan(&comments)
	if err != nil {
		return model.PagedData{}, service.CustomError.ServerError(ctx, "list comments failed")
	}
	totalPage := int(math.Ceil(float64(total) / float64(req.Limit)))
	return model.PagedData{
		Data: comments,
		Paged: model.Paged{
			Total:     total,
			Page:      req.Page,
			Size:      req.Limit,
			TotalPage: totalPage,
		},
	}, nil
}
