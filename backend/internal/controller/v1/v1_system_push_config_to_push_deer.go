package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PushConfigToPushDeer(ctx context.Context, req *system.PushConfigToPushDeerReq) (res *system.PushConfigToPushDeerRes, err error) {
	if service.SysPush().SavePushDeerConfig(ctx, req.Key) {
		service.UserCtx().GetCtx(ctx).Message = "保存成功"
	}
	return
}
