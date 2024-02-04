package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RoleGetParams(ctx context.Context, req *system.RoleGetParamsReq) (res *system.RoleGetParamsRes, err error) {
	res = new(system.RoleGetParamsRes)
	res.Menu, err = service.SysAuthRule().GetMenuList(ctx)
	return
}
