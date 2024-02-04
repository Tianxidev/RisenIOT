package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictTypeAdd(ctx context.Context, req *system.DictTypeAddReq) (res *system.DictTypeAddRes, err error) {
	err = service.SysDictType().Add(ctx, req, service.UserCtx().GetUserId(ctx))
	return
}
