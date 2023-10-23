package router

import "github.com/gin-gonic/gin"

type DeviceRouter struct{}
type BaseRouter struct{}

// InitRouter 路由初始化
func InitRouter(Router *gin.RouterGroup) {
	new(BaseRouter).New(Router)
	new(DeviceRouter).New(Router)
}
