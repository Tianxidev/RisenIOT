package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) ConfigDelete(ctx context.Context, req *system.ConfigDeleteReq) (res *system.ConfigDeleteRes, err error) {
	err = service.SysConfig().Delete(ctx, req.Ids)
	return
}
