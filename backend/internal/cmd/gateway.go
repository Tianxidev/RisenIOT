package cmd

import (
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Gateway = &gcmd.Command{
		Name:  "gateway",
		Brief: "启动 web 服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "正在启动 gateway 服务")
			service.Gateway().Start(ctx)
			return nil
		},
	}
)
