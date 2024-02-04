package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) LoginLogClear(ctx context.Context, req *system.LoginLogClearReq) (res *system.LoginLogClearRes, err error) {
	err = service.SysLoginLog().ClearLoginLog(ctx)
	return
}
