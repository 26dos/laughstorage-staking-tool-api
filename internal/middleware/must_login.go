package middleware

import (
	"api/internal/model"
	"api/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
)

func MustLogin(r *ghttp.Request) {
	ctx := r.Context()
	res := model.ApiRes{
		Code:    gcode.CodeNotAuthorized.Code(),
		Message: "Please login!",
	}
	_, err := service.User.GetCtxUser(ctx)
	if err != nil {
		r.Response.WriteJson(res)
		r.Exit()
	}
	r.Middleware.Next()
}
