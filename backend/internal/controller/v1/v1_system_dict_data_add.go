package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictDataAdd(ctx context.Context, req *system.DictDataAddReq) (res *system.DictDataAddRes, err error) {
	err = service.SysDictData().Add(ctx, req, service.UserCtx().GetUserId(ctx))
	return
}
