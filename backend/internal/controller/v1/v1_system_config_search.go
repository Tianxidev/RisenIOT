package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/system"
)

func (c *ControllerSystem) ConfigSearch(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
