package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) SysOperLogClear(ctx context.Context, req *system.SysOperLogClearReq) (res *system.SysOperLogClearRes, err error) {
	err = service.OperateLog().ClearLog(ctx)
	return
}
