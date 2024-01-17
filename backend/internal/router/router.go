package router

import (
	"backend/internal/controller/device"
	"backend/internal/controller/system"
	"backend/internal/model"
	"backend/internal/service"
	"backend/utility/token"
	"context"
	"github.com/goflyfox/gtoken/gtoken"
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

		// 初始化 gtoken 服务
		tokenService := token.Init(ctx, func(r *ghttp.Request) (string, interface{}) {

			// 判断是否是POST请求
			if r.Method != "POST" {
				r.Response.WriteJson(gtoken.Fail("请求方式错误"))
				g.Log().Notice(vContext, "用户访问登录接口, 但是请求方式不是POST: ", r.Method)
				r.ExitAll()
			}

			// 获取请求参数
			username := r.Get("username").String()
			password := r.Get("password").String()

			// 判断参数是否为空
			if username == "" || password == "" {
				r.Response.WriteJson(gtoken.Fail("用户名或密码不能为空"))
				g.Log().Notice(vContext, "用户访问登录接口, 但是用户名或密码为空")
				r.ExitAll()
			}

			// 验证用户信息
			userInfo, err := service.User().UserLoginVerifyFn(vContext, username, password)
			if err != nil {
				g.Log().Notice(vContext, "查询用户异常", err)
				r.Response.WriteJson(gtoken.Fail(err.Error()))
				r.ExitAll()
			}

			// 返回用户信息
			return userInfo.Username, userInfo
		})

		// 认证中间件配置
		err := tokenService.Middleware(ctx, group)
		if err != nil {
			return
		}

		group.Bind(
			system.NewV1(),
			device.NewV1(),
		)
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
