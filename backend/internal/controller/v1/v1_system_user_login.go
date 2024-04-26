package v1

import (
	"backend/api/v1/system"
	"backend/internal/model"
	"backend/internal/service"
	"backend/library/libUtils"
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerSystem) UserLogin(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error) {
	var (
		user        *model.LoginUserRes
		token       string
		permissions []string
		menuList    []*model.UserMenus
	)

	// 判断密码是否正确
	user, err = service.SysUser().GetAdminUserByUsernamePassword(ctx, req)
	if err != nil {
		return
	}

	// 更新登录信息
	err = service.SysUser().UpdateLoginInfo(ctx, user.Id, libUtils.GetClientIp(ctx))

	// 创建登录 token
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)

	// 获取客户端信息
	ip := libUtils.GetClientIp(ctx)

	// 获取客户端信息
	userAgent := libUtils.GetUserAgent(ctx)

	// 多端登录
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	user.UserPassword = ""

	// 生成 token
	token, err = service.Token().Get().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}

	// 获取用户菜单数据
	menuList, permissions, err = service.SysUser().GetAdminRules(ctx, user.Id)
	if err != nil {
		return
	}
	res = &system.UserLoginRes{
		UserInfo:    user,
		Token:       token,
		MenuList:    menuList,
		Permissions: permissions,
	}
	// 用户在线状态保存
	service.SysUserOnline().Invoke(gctx.New(), &model.SysUserOnlineParams{
		UserAgent: "",
		Uuid:      gmd5.MustEncrypt(token),
		Token:     token,
		Username:  user.UserName,
		Ip:        "",
	})
	return
}
