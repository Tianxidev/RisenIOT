package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictTypeEdit(ctx context.Context, req *system.DictTypeEditReq) (res *system.DictTypeEditRes, err error) {
	err = service.SysDictType().Edit(ctx, req, service.UserCtx().GetUserId(ctx))
	return
}
