package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictDataEdit(ctx context.Context, req *system.DictDataEditReq) (res *system.DictDataEditRes, err error) {
	err = service.SysDictData().Edit(ctx, req, service.UserCtx().GetUserId(ctx))
	return
}
