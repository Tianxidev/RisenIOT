package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) LoginLogDel(ctx context.Context, req *system.LoginLogDelReq) (res *system.LoginLogDelRes, err error) {
	err = service.SysLoginLog().DeleteLoginLogByIds(ctx, req.Ids)
	return
}
