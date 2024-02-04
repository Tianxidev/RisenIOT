package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) ConfigGet(ctx context.Context, req *system.ConfigGetReq) (res *system.ConfigGetRes, err error) {
	res, err = service.SysConfig().Get(ctx, req.Id)
	return
}
