package middleware

import (
	"backend/internal/model"
	"backend/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// Ctx 上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := new(model.Context)
	data, err := service.Token().Get().ParseToken(r)
	if err == nil {
		_ = gconv.Struct(data.Data, &customCtx.User)
	}
	service.UserCtx().InitCtx(r, customCtx)
	r.Middleware.Next()
}
