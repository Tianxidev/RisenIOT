package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PostAdd(ctx context.Context, req *system.PostAddReq) (res *system.PostAddRes, err error) {
	err = service.SysPost().Add(ctx, req)
	return
}
