package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictDataGet(ctx context.Context, req *system.DictDataGetReq) (res *system.DictDataGetRes, err error) {
	res, err = service.SysDictData().Get(ctx, req.DictCode)
	return
}
