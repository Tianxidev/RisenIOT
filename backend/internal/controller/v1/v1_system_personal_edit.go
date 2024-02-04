package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PersonalEdit(ctx context.Context, req *system.PersonalEditReq) (res *system.PersonalEditRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
