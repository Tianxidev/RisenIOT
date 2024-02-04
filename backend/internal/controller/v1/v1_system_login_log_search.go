package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) LoginLogSearch(ctx context.Context, req *system.LoginLogSearchReq) (res *system.LoginLogSearchRes, err error) {
	res, err = service.SysLoginLog().List(ctx, req)
	return
}
