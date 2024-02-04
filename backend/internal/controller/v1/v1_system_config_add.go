package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) ConfigAdd(ctx context.Context, req *system.ConfigAddReq) (res *system.ConfigAddRes, err error) {
	err = service.SysConfig().Add(ctx, req, service.UserCtx().GetUserId(ctx))
	return
}
