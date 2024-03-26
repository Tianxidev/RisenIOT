package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
)

// SDAP 海康设备发现协议
var (
	SDAP = &gcmd.Command{
		Name:  "sdap",
		Brief: "海康设备发现协议",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			return nil
		},
	}
)
