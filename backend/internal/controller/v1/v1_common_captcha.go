package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/common"
)

func (c *ControllerCommon) Captcha(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
