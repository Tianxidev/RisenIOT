package router

import (
	v1 "backend/internal/controller/v1"
	"backend/internal/model"
	"backend/internal/service"
	"backend/library/libRouter"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Router struct{}

var R = new(Router)

var (
	vContext context.Context
)

// BindController 绑定控制器
func (r *Router) BindController(ctx context.Context, router *ghttp.RouterGroup) {
	vContext = ctx

	// 中间件配置
	router.Middleware(service.Middleware().CORS)

	// v1 版本路由
	router.Group("/api/v1", func(group *ghttp.RouterGroup) {

		err := service.Token().Get().Middleware(group)
		if err != nil {
			g.Log().Error(ctx, "访问控制中间件异常", err)
			return
		}

		// 绑定路由
		group.Bind(
			v1.NewCommon(),
			v1.NewSystem(),
			v1.NewDevice(),
		)

		//自动绑定定义的模块
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})

	// 未匹配到路由时的处理 (统一返回 404)
	router.ALL("/*", r.DefaultHandler)
}

// DefaultHandler 未匹配到路由时的处理 (统一返回 404)
func (_ *Router) DefaultHandler(r *ghttp.Request) {
	g.Log().Notice(vContext, "用户访问未注册的路由: ", r.URL.Path)

	// 未匹配到路由时，统一返回: {"code":404,"message":"404 Not Found"}
	r.Response.WriteJson(model.DefaultHandlerResponse{
		Code:    404,
		Message: "Not Found",
	})
}
