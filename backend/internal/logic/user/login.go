package user

import (
	"backend/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLoginVerifyFn 用户登录验证
func (s *sUser) UserLoginVerifyFn(ctx context.Context, username string, password string) (user entity.RiseniotUser, err error) {

	// 加密密码
	password = gmd5.MustEncryptString(password)

	g.Log().Info(ctx, "用户登录验证", username, password)

	// 查询用户信息
	user, err = s.GetByUsername(ctx, username)
	if err != nil {
		g.Log().Notice(ctx, "用户不存在", username)
		return entity.RiseniotUser{}, gerror.New("用户不存在")
	}

	// 判断用户是否存在
	if user.Username == "" {
		g.Log().Notice(ctx, "用户不存在", username)
		return entity.RiseniotUser{}, gerror.New("用户不存在")
	}

	// 判断密码是否正确
	if user.Password != password {
		g.Log().Notice(ctx, "密码错误", username)
		return entity.RiseniotUser{}, gerror.New("用户名或密码错误")
	}

	return user, err
}
