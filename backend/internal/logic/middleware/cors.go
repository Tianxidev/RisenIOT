package middleware

import "github.com/gogf/gf/v2/net/ghttp"

// CORS 同源策略
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
