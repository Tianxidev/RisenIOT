package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PostDelete(ctx context.Context, req *system.PostDeleteReq) (res *system.PostDeleteRes, err error) {
	err = service.SysPost().Delete(ctx, req.Ids)
	return
}
