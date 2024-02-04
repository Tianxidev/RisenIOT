package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PostSearch(ctx context.Context, req *system.PostSearchReq) (res *system.PostSearchRes, err error) {
	res, err = service.SysPost().List(ctx, req)
	return
}
