package logic

import (
	v1 "api/api/user/v1"
	"api/internal/consts"
	"api/internal/dao"
	"api/internal/model"
	"api/internal/service"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var User = new(userLogic)

type userLogic struct {
}

func (l *userLogic) GetUser(ctx context.Context, req *v1.UserReq) (*model.UserInfo, error) {
	user, err := service.User.GetCtxUser(ctx)
	if err != nil {
		return nil, err
	}
	return &model.UserInfo{
		LoginName:   user.LoginName,
		Role:        user.Role,
		Wallet:      user.Wallet,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		KycStatus:   user.KycStatus,
		CreatedAt:   user.CreatedAt,
		LastLoginAt: user.LastLoginAt,
	}, nil
}

func (l *userLogic) KycVerify(ctx context.Context, req *v1.KycVerifyReq) (bool, error) {
	user, err := service.User.GetCtxUser(ctx)
	if err != nil {
		return false, err
	}
	if user.KycStatus == "verified" {
		return false, service.CustomError.ParameterError(ctx, "user kyc already verified")
	}
	proposals, err := dao.SliverProposals.Ctx(ctx).Where("p_user", user.Id).Count()
	if err != nil {
		return false, err
	}
	if proposals <= 0 {
		return false, service.CustomError.ParameterError(ctx, "You must create a proposal first!")
	}
	idName := fmt.Sprintf("%s%s", req.IdFirstName, req.IdLastName)
	err = service.Ali.InitCredential(ctx).CredentialVerify(ctx, req.IdNumber, idName, req.IdCardFront)
	if err != nil {
		return false, err
	}
	err = service.Ali.InitCredential(ctx).Id2MetaVerifyWithOCR(ctx, req.IdCardFront, req.IdCardBack)
	if err != nil {
		return false, err
	}
	err = dao.SliverUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.SliverUser.Ctx(ctx).Where("id", user.Id).Data(g.Map{
			"kyc_status": "verified",
		}).Update()
		if err != nil {
			g.Log().Error(ctx, "kyc verify error", err)
			return err
		}
		_, err = dao.SliverProposals.Ctx(ctx).Where("p_user", user.Id).Data(g.Map{
			"kyc_status":            "verified",
			"kyc_verification_time": time.Now(),
		}).Update()
		if err != nil {
			g.Log().Error(ctx, "kyc verify error", err)
			return err
		}
		return nil
	})
	if err != nil {
		return false, consts.ServerErr
	}
	return true, nil
}
