package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/system/v1"
)

func (c *ControllerV1) Pk(ctx context.Context, req *v1.PkReq) (res *v1.PkRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
