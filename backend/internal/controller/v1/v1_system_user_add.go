package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserAdd(ctx context.Context, req *system.UserAddReq) (res *system.UserAddRes, err error) {
	err = service.SysUser().Add(ctx, req)
	return
}
