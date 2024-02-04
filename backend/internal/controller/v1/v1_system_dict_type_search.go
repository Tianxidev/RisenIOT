package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DictTypeSearch(ctx context.Context, req *system.DictTypeSearchReq) (res *system.DictTypeSearchRes, err error) {
	res, err = service.SysDictType().List(ctx, req)
	return
}
