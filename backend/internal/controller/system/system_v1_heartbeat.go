package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/system/v1"
)

// Heartbeat 心跳检测
func (c *ControllerV1) Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (res *v1.HeartbeatRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
