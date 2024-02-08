package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/device"
)

func (c *ControllerDevice) InfoDelete(ctx context.Context, req *device.InfoDeleteReq) (res *device.InfoDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
