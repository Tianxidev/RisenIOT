package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserDelete(ctx context.Context, req *system.UserDeleteReq) (res *system.UserDeleteRes, err error) {
	err = service.SysUser().Delete(ctx, req.Ids)
	return
}
