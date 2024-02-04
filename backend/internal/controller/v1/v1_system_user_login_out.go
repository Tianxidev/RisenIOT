package v1

import (
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserLoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error) {
	err = service.Token().Get().RemoveToken(ctx, service.Token().Get().GetRequestToken(g.RequestFromCtx(ctx)))
	res = &system.UserLoginOutRes{
		Ok: err == nil,
	}
	return
}
