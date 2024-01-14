package device

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/device/v1"
)

func (c *ControllerV1) DeviceDataReceive(ctx context.Context, req *v1.DeviceDataReceiveReq) (res *v1.DeviceDataReceiveRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
