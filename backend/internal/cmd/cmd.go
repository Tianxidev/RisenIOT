package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = &gcmd.Command{
		Name: "main",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return nil
		},
	}

	All = &gcmd.Command{
		Name:        "all",
		Brief:       "启动所有服务",
		Description: "启动所有服务时执行此命令",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("正在启动所有服务")

			// 启动网关服务
			go func() {
				err := Gateway.Func(ctx, parser)
				if err != nil {
					g.Log().Error(ctx, err)
				}
			}()

			// 启动 web 服务
			go func() {
				err := Web.Func(ctx, parser)
				if err != nil {
					g.Log().Error(ctx, err)
				}
			}()

			// 主线程循环检测服务停止
			for {

				// 检测服务是否停止
				select {
				case <-ctx.Done():
					fmt.Println("所有服务已停止")
					return
				}
			}

		},
	}
)
