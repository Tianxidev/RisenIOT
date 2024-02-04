package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PersonalResetPwd(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
