package v1

import (
	"backend/internal/consts"
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) CategorySearch(ctx context.Context, req *device.CategorySearchReq) (res *device.CategorySearchRes, err error) {
	res = &device.CategorySearchRes{}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	res.CurrentPage = req.PageNum
	res.Total, _, res.List, err = service.DeviceCategory().List(ctx, req)
	res.Kind, err = service.DeviceKind().Get(ctx, req.KindId)
	liberr.ErrIsNil(ctx, err, "获取数据类型失败")
	return
}
