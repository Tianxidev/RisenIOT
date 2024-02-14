package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Gateway = &gcmd.Command{
		Name:  "gateway",
		Brief: "启动 web 服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := make(chan os.Signal, 1)

			// 设置日志异步输出
			g.Log().SetAsync(true)

			g.Log().Info(ctx, "正在启动 gateway 服务, PID = "+gconv.String(gproc.Pid()))
			// service.Gateway().Start(ctx)
			signal.Notify(c, os.Interrupt)

			// for {
			// 	select {
			// 	case <-c:
			// 		g.Log().Info(ctx, "gateway 服务已停止")
			// 		return nil
			// 	}
			// }

			return nil
		},
	}
)
