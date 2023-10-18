package router

import (
	"RisenIOT/backend/controller"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

// InitBaseRouter 初始化基础路由
func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	r := Router.Group("base")
	r.GET("/version", controller.GetVersion)
	r.GET("/consoleLogWS", controller.ConsoleLogWS)
}
