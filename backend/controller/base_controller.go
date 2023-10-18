package controller

import (
	"RisenIOT/backend/config"
	"github.com/gin-gonic/gin"
)

// GetVersion 获取系统版本信息
func GetVersion(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":       200,
		"SysName":    config.SysName,
		"SysVersion": config.SysVersion,
	})
}

// ConsoleLogWS 控制台日志WebSocket
func ConsoleLogWS(c *gin.Context) {

}
