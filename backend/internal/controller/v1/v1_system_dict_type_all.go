package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictTypeAll(ctx context.Context, req *system.DictTypeAllReq) (res *system.DictTypeAllRes, err error) {
	res = new(system.DictTypeAllRes)
	res.DictType, err = service.SysDictType().GetAllDictType(ctx)
	return
}
