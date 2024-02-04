package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RoleList(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error) {
	res, err = service.SysRole().GetRoleListSearch(ctx, req)
	return
}
