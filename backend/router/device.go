package router

import (
	"RisenIOT/backend/controller"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

// InitDeviceRouter 初始化设备相关路由
func (s *BaseRouter) InitDeviceRouter(Router *gin.RouterGroup) {
	router := Router.Group("device")

	var deviceController = new(controller.DeviceController)

	router.GET("/list", deviceController.DeviceList)
	router.POST("/data/receive/emqx/webhook", deviceController.DeviceDataReceiveFromEmqxWebHook)
	router.POST("/cmd/push", deviceController.DeviceCmdPush)
	router.POST("/cmd/execute/lamp", deviceController.DeviceExecute)
	router.POST("/cmd/execute/lamp/dimming", deviceController.DeviceLampDimming)
}
