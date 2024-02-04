package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RoleAdd(ctx context.Context, req *system.RoleAddReq) (res *system.RoleAddRes, err error) {
	err = service.SysRole().AddRole(ctx, req)
	return
}
