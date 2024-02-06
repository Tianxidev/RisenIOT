package v1

import (
	"backend/internal/consts"
	"backend/internal/service"
	"context"

	"backend/api/v1/device"
)

// KindSearch 获取设备种类列表
func (c *ControllerDevice) KindSearch(ctx context.Context, req *device.KindSearchReq) (res *device.KindSearchRes, err error) {
	res = &device.KindSearchRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	res.CurrentPage = req.PageNum
	res.Total, _, res.List, err = service.DeviceKind().List(ctx, req)
	return
}
