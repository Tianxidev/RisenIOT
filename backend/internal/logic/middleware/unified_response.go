package middleware

import (
	"backend/internal/model"
	"backend/internal/service"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UnifiedResponseHandler 统一返回
func (s *sMiddleware) UnifiedResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
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
			msg = service.UserCtx().Get(r.Context()).Message
		}
	}

	// 统一返回对象
	r.Response.WriteJson(model.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
