package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PushTextToPushDeer(ctx context.Context, req *system.PushTextToPushDeerReq) (res *system.PushTextToPushDeerRes, err error) {
	if service.SysPush().TextMsgSendPushDeer(ctx, req.Key, req.Text) {
		service.UserCtx().GetCtx(ctx).Message = "发送成功"
	}
	return
}
