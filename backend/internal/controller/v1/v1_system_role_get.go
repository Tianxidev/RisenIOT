package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RoleGet(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error) {
	res = new(system.RoleGetRes)
	res.Role, err = service.SysRole().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.MenuIds, err = service.SysRole().GetFilteredNamedPolicy(ctx, req.Id)
	return
}
