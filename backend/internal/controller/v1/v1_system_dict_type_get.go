package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictTypeGet(ctx context.Context, req *system.DictTypeGetReq) (res *system.DictTypeGetRes, err error) {
	res = new(system.DictTypeGetRes)
	res.DictType, err = service.SysDictType().Get(ctx, req)
	return
}
