package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) GroupList(ctx context.Context, req *device.GroupListReq) (res *device.GroupListRes, err error) {
	res, err = service.DeviceGroup().GroupList(ctx)

	return
}
