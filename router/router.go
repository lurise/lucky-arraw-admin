package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	api2 "luckyarrow/app/api"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.ALL("/hello", api2.Hello)
		group.ALL("/prize", api2.PrizeApi)
		group.ALL("/qq",api2.QQApi)
	})
}
