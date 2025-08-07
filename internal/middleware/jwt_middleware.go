package middleware

import (
	"api/internal/consts"
	"api/internal/service"
	"api/utility"
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func JwtMiddleware(r *ghttp.Request) {
	auth := r.Header.Get("Authorization")
	token := strings.Split(auth, " ")[1]
	ctx := r.Context()
	if utility.EmptyString(token) || token == "undefined" {
		_ = context.WithValue(ctx, consts.JwtUserKey, nil)
		r.Middleware.Next()
		return
	}
	pre, _ := g.Config().Get(ctx, "jwt.pre")
	saveToken, err := g.Redis().Get(ctx, pre.String()+token)
	if err != nil || saveToken.String() != "1" {
		_ = context.WithValue(ctx, consts.JwtUserKey, nil)
		r.Middleware.Next()
		return
	}
	//* parse token
	jwtTool := service.NewJwt(ctx)
	c, err := jwtTool.AuthToken(token)
	if err != nil {
		_ = context.WithValue(ctx, consts.JwtUserKey, nil)
		r.Middleware.Next()
		return
	}
	user, err := service.User.GetUserByAddress(ctx, c.Uid, false)
	if err != nil {
		_ = context.WithValue(ctx, consts.JwtUserKey, nil)
		r.Middleware.Next()
		return
	}
	ctx = context.WithValue(ctx, consts.JwtUserKey, user)
	r.SetCtx(ctx)
	r.Middleware.Next()
}
