package middleware

import (
	"api/internal/model"
	"api/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
)

func MustAdministrator(r *ghttp.Request) {
	user, err := service.User.GetCtxUser(r.Context())
	res := model.ApiRes{
		Code:    gcode.CodeSecurityReason.Code(),
		Message: "Unauthorized!",
	}
	if err != nil {
		r.Response.WriteJson(res)
		r.Exit()
	}
	if user.Role != "admin" {
		r.Response.WriteJson(res)
		r.Exit()
	}
	r.Middleware.Next()
}
