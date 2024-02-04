package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserEdit(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error) {
	err = service.SysUser().Edit(ctx, req)
	return
}
