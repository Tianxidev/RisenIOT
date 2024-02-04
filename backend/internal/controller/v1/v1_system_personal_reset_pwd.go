package v1

import (
	"backend/internal/dao"
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"backend/library/libUtils"
	"backend/library/liberr"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PersonalResetPwd(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error) {
	userId := service.UserCtx().GetUserId(ctx)
	salt := grand.S(10)
	password := libUtils.EncryptPassword(req.Password, salt)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(userId).Update(g.Map{
			dao.SysUser.Columns().UserSalt:     salt,
			dao.SysUser.Columns().UserPassword: password,
		})
		liberr.ErrIsNil(ctx, err, "重置用户密码失败")
	})
	return
}
