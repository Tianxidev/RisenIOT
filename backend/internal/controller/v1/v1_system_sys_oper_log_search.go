package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) SysOperLogSearch(ctx context.Context, req *system.SysOperLogSearchReq) (res *system.SysOperLogSearchRes, err error) {
	res, err = service.OperateLog().List(ctx, req)
	return
}
