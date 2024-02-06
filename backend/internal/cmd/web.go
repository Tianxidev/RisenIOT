package cmd

import (
	"backend/internal/router"
	"backend/internal/service"
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Web = &gcmd.Command{
		Name:  "web",
		Brief: "启动 web 服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			var (
				defaultPort = 10001
				cfg         = g.Cfg()
				s           = g.Server()
				oai         = s.GetOpenApi()
			)

			g.Log().Info(ctx, "正在启动 web 服务")
			cfg.GetAdapter().(*gcfg.AdapterFile).SetFileName("config.yaml")
			serviceName, _ := cfg.Get(ctx, "server.name")

			oai.Info.Title = serviceName.String()
			oai.Info.Description = "开放接口文档"
			oai.Config.CommonResponse = `Code`
			oai.Config.CommonResponseDataField = `Data`

			// 设置 web 端口
			serverPort, _ := cfg.Get(ctx, "server.port")
			if serverPort == nil {
				g.Log().Notice(ctx, "读取工作端口失败, 将使用默认端口: "+strconv.Itoa(defaultPort))
			} else {
				g.Log().Info(ctx, "设置端口 ["+strconv.Itoa(serverPort.Int())+"] 成功!")
				defaultPort = serverPort.Int()
			}

			// 注册根路由
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().Ctx)
				group.Middleware(service.Middleware().UnifiedResponseHandler)
				router.R.BindController(ctx, group)
			})

			// 关闭路由信息打印
			//s.SetDumpRouterMap(false)
			s.SetPort(defaultPort)
			s.Run()
			return nil
		},
	}
)
