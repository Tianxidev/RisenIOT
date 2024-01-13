package router

import (
	"backend/internal/model"
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Router struct{}

var R = new(Router)

func (r *Router) BindController(ctx context.Context, router *ghttp.RouterGroup) {

	// 中间件配置
	router.Middleware(
		service.Middleware().Ctx,
		service.Middleware().CORS,
		service.Middleware().UnifiedResponseHandler,
	)

	// v1 版本路由
	router.Group("/api/v1", func(group *ghttp.RouterGroup) {

	})

	// 未匹配到路由时的处理
	router.ALL("/*", func(r *ghttp.Request) {
		g.Log().Notice(ctx, "用户访问未注册的路由: ", r.URL.Path)

		// 未匹配到路由时，统一返回: {"code":404,"message":"404 Not Found"}
		r.Response.WriteJson(model.DefaultHandlerResponse{
			Code:    404,
			Message: "Not Found",
		})
	})
}
