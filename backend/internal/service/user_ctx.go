// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/internal/model"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IUserCtx interface {
		// InitCtx 上下文初始化
		InitCtx(r *ghttp.Request, customCtx *model.Context)
		// GetCtx 获取上下文
		GetCtx(ctx context.Context) *model.Context
		// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
		SetUser(ctx context.Context, ctxUser *model.ContextUser)
		// GetLoginUser 获取当前登陆用户信息
		GetLoginUser(ctx context.Context) *model.ContextUser
		// GetUserId 获取当前登录用户id
		GetUserId(ctx context.Context) uint64
	}
)

var (
	localUserCtx IUserCtx
)

func UserCtx() IUserCtx {
	if localUserCtx == nil {
		panic("implement not found for interface IUserCtx, forgot register?")
	}
	return localUserCtx
}

func RegisterUserCtx(i IUserCtx) {
	localUserCtx = i
}
