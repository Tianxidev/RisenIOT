package v1

import (
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"backend/api/v1/device"
)

func (c *ControllerDevice) DataAdd(ctx context.Context, req *device.DataAddReq) (res *device.DataAddRes, err error) {
	g.Log().Print(ctx, "传感器上报数据:", req)

	decode, err := service.DataCodec().HttpDecode(ctx, req)
	if err != nil {
		return nil, err
	}

	g.Log().Print(ctx, "解码传感器结果:", decode)
	//msg, err := httpParse.Encode(ctx, req)
	//err = httpParse.Save(ctx, msg)

	return
}
