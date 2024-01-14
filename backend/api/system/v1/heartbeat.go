package v1

import "github.com/gogf/gf/v2/frame/g"

// HeartbeatReq 心跳检测 请求参数
type HeartbeatReq struct {
	g.Meta `path:"/system/heartbeat" tags:"系统" method:"get" summary:"心跳检测"`
}

// HeartbeatRes 心跳检测 响应参数
type HeartbeatRes struct {
	g.Meta `mime:"application/json"`
}
