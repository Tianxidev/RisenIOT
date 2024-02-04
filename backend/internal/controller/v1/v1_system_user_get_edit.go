package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserGetEdit(ctx context.Context, req *system.UserGetEditReq) (res *system.UserGetEditRes, err error) {
	res, err = service.SysUser().GetEditUser(ctx, req.Id)
	return
}
