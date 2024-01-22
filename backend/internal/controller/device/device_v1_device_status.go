package device

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/device/v1"
)

// DeviceStatus 设备状态
func (c *ControllerV1) DeviceStatus(ctx context.Context, req *v1.DeviceStatusReq) (res *v1.DeviceStatusRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
