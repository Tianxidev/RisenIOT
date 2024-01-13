package userctx

import (
	"backend/internal/consts"
	"backend/internal/model"
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sUserCtx struct{}
)

func init() {
	service.RegisterUserCtx(New())
}

func New() service.IUserCtx {
	return &sUserCtx{}
}

// Init 上下文初始化
func (s *sUserCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获取上下文
func (s *sUserCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// Set 上下文设置
func (s *sUserCtx) Set(ctx context.Context, ctxUser *model.Context) {
	s.Get(ctx).Code = ctxUser.Code
	s.Get(ctx).Message = ctxUser.Message
	s.Get(ctx).Data = ctxUser.Data
}
