package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserGetByIds(ctx context.Context, req *system.UserGetByIdsReq) (res *system.UserGetByIdsRes, err error) {
	res = new(system.UserGetByIdsRes)
	res.List, err = service.SysUser().GetUsers(ctx, req.Ids)
	return
}
