package cmd

import (
	"api/internal/controller/auth"
	"api/internal/controller/comments"
	"api/internal/controller/proposals"
	"api/internal/controller/user"
	"api/internal/service"
	"context"

	"api/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			go service.Staking.Start(ctx)
			go service.Allocator.Start(ctx)
			// cors config
			s.BindMiddlewareDefault(func(r *ghttp.Request) {
				r.Response.CORSDefault()
				r.Middleware.Next()
			})
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.Response)
				group.Middleware(middleware.JwtMiddleware)
				group.Group("/auth").Bind(auth.NewV1())
				group.Group("/proposals").Bind(proposals.NewUser())
				group.Group("/proposals", func(p *ghttp.RouterGroup) {
					p.Middleware(middleware.MustLogin)
					p.Bind(proposals.NewAuth())
					p.Bind(proposals.NewManage()).Middleware(middleware.MustAdministrator)
				})
				group.Group("/comments").Bind(comments.NewV1())
				group.Group("/comments").Bind(comments.NewAuth()).Middleware(middleware.MustLogin)
				group.Group("/user").Bind(user.NewV1()).Middleware(middleware.MustLogin)
			})
			s.Run()
			return nil
		},
	}
)
