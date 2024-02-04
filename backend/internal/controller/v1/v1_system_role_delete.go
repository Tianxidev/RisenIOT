package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RoleDelete(ctx context.Context, req *system.RoleDeleteReq) (res *system.RoleDeleteRes, err error) {
	err = service.SysRole().DeleteByIds(ctx, req.Ids)
	return
}
