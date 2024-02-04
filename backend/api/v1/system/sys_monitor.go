package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MonitorServerReq struct {
	g.Meta `path:"/system/monitor/server" tags:"服务监控" method:"get" summary:"服务监控"`
}

type MonitorServerRes g.Map
