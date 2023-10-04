package controller

import (
	"RisenIOT/backend/config"
	"github.com/gin-gonic/gin"
)

func GetVersion(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":       200,
		"SysName":    config.SysName,
		"SysVersion": config.SysVersion,
	})
}
