package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) SysUserOnlineForceLogoutAll(ctx context.Context, req *system.SysUserOnlineForceLogoutAllReq) (res *system.SysUserOnlineForceLogoutAllRes, err error) {
	err = service.SysUserOnline().ForceLogoutAll(ctx)
	return
}
