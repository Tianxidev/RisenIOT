package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserStatus(ctx context.Context, req *system.UserStatusReq) (res *system.UserStatusRes, err error) {
	err = service.SysUser().ChangeUserStatus(ctx, req)
	return
}
