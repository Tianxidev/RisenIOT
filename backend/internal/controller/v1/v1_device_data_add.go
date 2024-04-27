package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"backend/api/v1/device"
)

// DataAdd 添加设备数据
func (c *ControllerDevice) DataAdd(ctx context.Context, req *device.DataAddReq) (res *device.DataAddRes, err error) {
	decode, err := service.DataCodec().HttpDecode(ctx, req)
	liberr.ErrIsNil(ctx, err, "解码传感器数据失败")
	g.Log().Print(ctx, "解码传感器结果:", decode.DeviceInfo, decode.EventList, decode.DataList)
	err = service.DataCodec().Save(ctx, decode)
	liberr.ErrIsNil(ctx, err, "保存传感器数据失败")

	service.UserCtx().GetCtx(ctx).Message = "上报成功"

	return
}
