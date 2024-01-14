// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/internal/model"
	"backend/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IUser interface {
		// InitCtx 上下文初始化
		InitCtx(r *ghttp.Request, customCtx *model.Context)
		// GetCtx 获取上下文
		GetCtx(ctx context.Context) *model.Context
		// SetCtx 上下文设置
		SetCtx(ctx context.Context, ctxUser *model.Context)
		// UserLoginVerifyFn 用户登录验证
		UserLoginVerifyFn(ctx context.Context, username string, password string) (user entity.RiseniotUser, err error)
		// GetByUsername 根据用户名获取用户信息
		GetByUsername(ctx context.Context, username string) (data entity.RiseniotUser, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
