package v1

import (
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"backend/library/libUtils"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PersonalEdit(ctx context.Context, req *system.PersonalEditReq) (res *system.PersonalEditRes, err error) {
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	res = new(system.PersonalEditRes)
	res.UserInfo, err = service.Personal().EditPersonal(ctx, req)
	if err != nil {
		return
	}
	key := gconv.String(res.UserInfo.Id) + "-" + gmd5.MustEncryptString(res.UserInfo.UserName) + gmd5.MustEncryptString(res.UserInfo.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(res.UserInfo.Id) + "-" + gmd5.MustEncryptString(res.UserInfo.UserName) + gmd5.MustEncryptString(res.UserInfo.UserPassword+ip+userAgent)
	}
	res.UserInfo.UserPassword = ""
	res.Token, err = service.Token().Get().GenerateToken(ctx, key, res.UserInfo)
	return
}
