package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RoleEdit(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error) {
	err = service.SysRole().EditRole(ctx, req)
	return
}
