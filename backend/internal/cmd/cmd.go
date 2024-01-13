package cmd

import (
	"backend/internal/router"
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcfg"
)

var (
	Main = gcmd.Command{
		Name:  "Web Service",
		Usage: "web",
		Brief: "启动 web 服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			// 默认端口
			var defaultPort int = 10001

			// 设置配置文件
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config.yaml")

			// 读取配置文件
			serverPort, _ := g.Cfg().Get(ctx, "server.port")
			if serverPort == nil {
				g.Log().Notice(ctx, "读取工作端口失败, 将使用默认端口: "+strconv.Itoa(defaultPort))
			} else {
				g.Log().Info(ctx, "设置端口 ["+strconv.Itoa(serverPort.Int())+"] 成功!")
				defaultPort = serverPort.Int()
			}

			// 创建 gServer
			s := g.Server()

			// 注册根路由
			s.Group("/", func(group *ghttp.RouterGroup) {
				router.R.BindController(ctx, group)
			})

			// 关闭路由信息打印
			s.SetDumpRouterMap(false)
			s.SetPort(defaultPort)
			s.Run()
			return nil
		},
	}
)
