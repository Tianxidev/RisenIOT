package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserEdit(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
