package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictDataDelete(ctx context.Context, req *system.DictDataDeleteReq) (res *system.DictDataDeleteRes, err error) {
	err = service.SysDictData().Delete(ctx, req.Ids)
	return
}
