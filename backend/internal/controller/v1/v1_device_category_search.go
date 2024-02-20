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

	// 查询产品信息
	res.Kind, err = service.DeviceKind().Get(ctx, req.KindId)
	liberr.ErrIsNil(ctx, err, "不存在的产品类型")

	// 查询数据类型列表
	res.Total, _, res.List, err = service.DeviceCategory().List(ctx, req)
	liberr.ErrIsNil(ctx, err, "获取数据类型列表失败")

	return
}
