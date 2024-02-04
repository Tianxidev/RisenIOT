package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) ConfigSearch(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error) {
	res, err = service.SysConfig().List(ctx, req)
	return
}
