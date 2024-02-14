package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = &gcmd.Command{
		Name: "main",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "即将以单体模式启动...")

			// 启动所有服务
			All.Func(ctx, parser)

			return nil
		},
	}

	All = &gcmd.Command{
		Name:        "all",
		Brief:       "启动所有服务",
		Description: "启动所有服务时执行此命令",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			in1 := make(chan string)
			in2 := make(chan string)

			g.Log().Info(ctx, "正在启动所有服务")

			// 设置日志异步输出
			g.Log().SetAsync(true)

			// 启动网关服务
			go func() {
				defer close(in1)
				err := Gateway.Func(context.TODO(), parser)
				if err != nil {
					g.Log().Error(ctx, err)
				}
			}()

			// 启动 web 服务
			go func() {
				defer close(in2)
				err := Web.Func(context.TODO(), parser)
				if err != nil {
					g.Log().Error(ctx, err)
				}
			}()

			// 主线程循环检测服务停止
			for {
				select {
				case _, ok := <-in1:
					if !ok {
						g.Log().Info(ctx, "网关服务已停止")
						in1 = nil
					}
				case _, ok := <-in2:
					if !ok {
						g.Log().Info(ctx, "web 服务已停止")
						in2 = nil
					}
				}

				// 检查协程管道是否都已经关闭
				if in1 == nil && in2 == nil {
					return nil
				}
			}

			// for {
			// 	// 检测服务是否停止
			// 	select {
			// 	case <-ctx.Done():
			// 		fmt.Println("所有服务已停止")
			// 		return
			// 	}
			// }

		},
	}
)
