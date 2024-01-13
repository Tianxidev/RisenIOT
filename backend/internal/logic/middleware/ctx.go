package middleware

import (
	"backend/internal/model"
	"backend/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Ctx 上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{}
	service.UserCtx().Init(r, customCtx)
	r.Middleware.Next()
}
