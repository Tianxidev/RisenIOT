package user_ctx

import (
	"backend/internal/consts"
	"backend/internal/model"
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sUserCtx struct{}

func init() {
	service.RegisterUserCtx(New())
}

func New() *sUserCtx {
	return &sUserCtx{}
}

// InitCtx 上下文初始化
func (s *sUserCtx) InitCtx(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// GetCtx 获取上下文
func (s *sUserCtx) GetCtx(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sUserCtx) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.GetCtx(ctx).User = ctxUser
}

// GetLoginUser 获取当前登陆用户信息
func (s *sUserCtx) GetLoginUser(ctx context.Context) *model.ContextUser {
	context := s.GetCtx(ctx)
	if context == nil {
		return nil
	}
	return context.User
}

// GetUserId 获取当前登录用户id
func (s *sUserCtx) GetUserId(ctx context.Context) uint64 {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.Id
	}
	return 0
}
