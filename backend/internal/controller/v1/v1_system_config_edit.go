package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) ConfigEdit(ctx context.Context, req *system.ConfigEditReq) (res *system.ConfigEditRes, err error) {
	err = service.SysConfig().Edit(ctx, req, service.UserCtx().GetUserId(ctx))
	return
}
