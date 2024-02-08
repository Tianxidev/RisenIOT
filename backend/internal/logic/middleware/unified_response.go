package middleware

import (
	"backend/internal/model"
	"backend/internal/service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

// UnifiedResponseHandler 统一返回
func (s *sMiddleware) UnifiedResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if err != nil {
		r.Response.ClearBuffer()
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = gstr.Replace(err.Error(), "exception recovered: ", "")
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
			msg = service.UserCtx().GetCtx(r.Context()).Message
		}
	}

	// 设置响应码
	r.Response.WriteHeader(http.StatusOK)

	// 统一返回对象
	r.Response.WriteJson(model.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
