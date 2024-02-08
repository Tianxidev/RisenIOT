package v1

import (
	"backend/internal/consts"
	"backend/internal/service"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) InfoSearch(ctx context.Context, req *device.InfoSearchReq) (res *device.InfoSearchRes, err error) {
	res = &device.InfoSearchRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	res.CurrentPage = req.PageNum
	res.Total, _, res.List, err = service.DeviceInfo().List(ctx, req)
	return
}
