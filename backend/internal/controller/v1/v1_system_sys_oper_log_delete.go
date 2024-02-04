package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) SysOperLogDelete(ctx context.Context, req *system.SysOperLogDeleteReq) (res *system.SysOperLogDeleteRes, err error) {
	err = service.OperateLog().DeleteByIds(ctx, req.OperIds)
	return
}
