package controller

import (
	"RisenIOT/backend/global"
	"RisenIOT/backend/internal/casbin"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

// GetVersion 获取系统版本信息
func (nbc *BaseController) GetVersion(c *gin.Context) {
	global.Logger.INFO("请求获取系统版本信息")
	c.JSON(200, gin.H{
		"code":       200,
		"SysName":    global.SysName,
		"SysVersion": global.SysVersion,
	})
}

// CasbinReload 权限管理重启
func (nbc *BaseController) CasbinReload(c *gin.Context) {
	casbin.SetupCasbinEnforcer()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

// ConsoleLogWS 控制台日志WebSocket
func (nbc *BaseController) ConsoleLogWS(c *gin.Context) {

}
