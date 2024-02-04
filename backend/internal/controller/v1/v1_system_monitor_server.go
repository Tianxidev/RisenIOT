package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/system"
)

func (c *ControllerSystem) MonitorServer(ctx context.Context, req *system.MonitorServerReq) (res *system.MonitorServerRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
