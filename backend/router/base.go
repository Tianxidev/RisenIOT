package router

import (
	"RisenIOT/backend/controller"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

// InitBaseRouter 初始化基础路由
func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	r := Router.Group("base")
	base := controller.NewBaseController()
	r.GET("/sys/version", base.GetVersion)
	r.GET("/auth/reload", base.CasbinReload)
	r.GET("/log/consoleLogWS", base.ConsoleLogWS)
}
