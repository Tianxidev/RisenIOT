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
		// Init 上下文初始化
		Init(r *ghttp.Request, customCtx *model.Context)
		// Get 获取上下文
		Get(ctx context.Context) *model.Context
		// Set 上下文设置
		Set(ctx context.Context, ctxUser *model.Context)
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
