package router

import (
	"RisenIOT/backend/controller"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

// InitDeviceRouter 初始化设备相关路由
func (s *BaseRouter) InitDeviceRouter(Router *gin.RouterGroup) {
	r := Router.Group("device")
	r.POST("/cmd/push", controller.DeviceCmdPush)
}
