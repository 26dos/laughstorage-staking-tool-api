package service

import (
	"api/internal/consts"
	"api/internal/dao"
	"api/internal/model/entity"
	"context"
	"database/sql"
	"slices"
	"strings"

	"errors"

	"github.com/gogf/gf/v2/frame/g"
)

var User = new(user)

type user struct{}

func (u *user) GetUserByAddress(ctx context.Context, address string, noJoin bool) (*entity.SliverUser, error) {
	user := entity.SliverUser{}
	db := dao.SliverUser.Ctx(ctx)
	err := db.Where("wallet", address).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) && noJoin {
			administrator := g.Cfg().MustGet(ctx, "app.administrator").String()
			administratorArr := strings.Split(administrator, ",")
			role := "user"
			if slices.Contains(administratorArr, address) {
				role = "admin"
			}
			_, err = db.Ctx(ctx).Insert(&entity.SliverUser{
				Wallet: address,
				Role:   role,
			})
			if err != nil {
				return nil, err
			}
			return u.GetUserByAddress(ctx, address, false)
		}
		return nil, err
	}
	return &user, nil
}

// GetCtxUser 获取当前登录用户
func (u *user) GetCtxUser(ctx context.Context) (*entity.SliverUser, error) {
	value := g.RequestFromCtx(ctx).Context().Value(consts.JwtUserKey)
	if value == nil {
		return nil, errors.New("please login")
	}
	return value.(*entity.SliverUser), nil
}
