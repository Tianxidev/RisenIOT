package router

import (
	"RisenIOT/backend/controller/apibase"
	"github.com/gin-gonic/gin"
)

// New 初始化路由
func (s *BaseRouter) New(Router *gin.RouterGroup) {

	// 路由分组
	router := Router.Group("base")

	// 创建控制器
	controller := new(apibase.Controller)

	// 注册接口
	router.GET("/sys/version", controller.GetVersion)
	router.GET("/auth/reload", controller.CasbinReload)
	//router.GET("/log/consoleLogWS", controller.ConsoleLogWS)
	router.POST("/login", controller.Login)
}
