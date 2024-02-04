package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserResetPwd(ctx context.Context, req *system.UserResetPwdReq) (res *system.UserResetPwdRes, err error) {
	err = service.SysUser().ResetUserPwd(ctx, req)
	return
}
