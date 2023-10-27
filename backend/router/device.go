package router

import (
	"RisenIOT/backend/controller/ApiDevice"
	"RisenIOT/backend/controller/ApiUnisound"
	"github.com/gin-gonic/gin"
)

// New 初始化相关路由
func (s *DeviceRouter) New(Router *gin.RouterGroup) {

	// 路由分组
	router := Router.Group("device")

	// 创建控制器
	deviceController := new(ApiDevice.Controller)
	lampController := new(ApiUnisound.LampController)

	// 设备管理 - 查询设备列表
	router.GET("/list", deviceController.DeviceList)

	// 设备管理 - 接收设备数据 EMQX WebHook
	router.POST("/data/receive/emqx/webhook", deviceController.ReceiveDataFromEmqxWebHook)

	// 设备管理 - 设备命令透传推送
	router.POST("/cmd/push", deviceController.DeviceCmdPush)

	// 设备管理 - 云知声灯控 - 灯开关接口
	router.POST("/unisound/lamp/execute", lampController.LampOpenOrClose)

	// 设备管理 - 云知声灯控 - 灯调光接口
	router.POST("/unisound/lamp/dimming", lampController.LampDimming)

	// 设备管理 - 云知声灯控 - 查询灯状态接口
	router.POST("/unisound/lamp/status", lampController.LampStatus)

}
