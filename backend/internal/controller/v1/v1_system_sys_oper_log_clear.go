package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/system"
)

func (c *ControllerSystem) SysOperLogClear(ctx context.Context, req *system.SysOperLogClearReq) (res *system.SysOperLogClearRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
