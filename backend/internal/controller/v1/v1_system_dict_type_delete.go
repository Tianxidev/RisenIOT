package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictTypeDelete(ctx context.Context, req *system.DictTypeDeleteReq) (res *system.DictTypeDeleteRes, err error) {
	err = service.SysDictType().Delete(ctx, req.DictIds)
	return
}
