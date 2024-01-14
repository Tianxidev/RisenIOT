package system

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/system/v1"
)

func (c *ControllerV1) Version(ctx context.Context, req *v1.VersionReq) (res *v1.VersionRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
