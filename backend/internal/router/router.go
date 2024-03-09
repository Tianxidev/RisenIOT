package router

import (
	v1 "backend/internal/controller/v1"
	"backend/internal/service"
	"backend/library/libRouter"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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

		err := service.Token().Middleware(group)
		if err != nil {
			g.Log().Error(ctx, "访问控制中间件异常", err)
			return
		}

		// 绑定路由
		group.Bind(
			v1.NewCommon(),
			v1.NewSystem(),
			v1.NewDevice(),
			v1.NewHome(),
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
	panic(gerror.NewCode(gcode.New(404, "", nil), "错误的请求方法或地址"))
}
