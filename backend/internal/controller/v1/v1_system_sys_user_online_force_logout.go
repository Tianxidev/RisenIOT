package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) SysUserOnlineForceLogout(ctx context.Context, req *system.SysUserOnlineForceLogoutReq) (res *system.SysUserOnlineForceLogoutRes, err error) {
	err = service.SysUserOnline().ForceLogout(ctx, req.Ids)
	return
}
