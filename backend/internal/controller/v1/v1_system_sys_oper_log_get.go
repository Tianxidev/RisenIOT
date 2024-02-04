package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) SysOperLogGet(ctx context.Context, req *system.SysOperLogGetReq) (res *system.SysOperLogGetRes, err error) {
	res = new(system.SysOperLogGetRes)
	res.SysOperLogInfoRes, err = service.OperateLog().GetByOperId(ctx, req.OperId)
	return
}
