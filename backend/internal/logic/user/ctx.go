package user

import (
	"backend/internal/consts"
	"backend/internal/model"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

// InitCtx 上下文初始化
func (s *sUser) InitCtx(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// GetCtx 获取上下文
func (s *sUser) GetCtx(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetCtx 上下文设置
func (s *sUser) SetCtx(ctx context.Context, ctxUser *model.Context) {
	s.GetCtx(ctx).Code = ctxUser.Code
	s.GetCtx(ctx).Message = ctxUser.Message
	s.GetCtx(ctx).Data = ctxUser.Data
}
