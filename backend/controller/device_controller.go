package controller

import (
	"RisenIOT/backend/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

// DeviceCmdPush 设备命令推送
func DeviceCmdPush(c *gin.Context) {

	// 读取post json
	var json map[string]interface{}

	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法解析json",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析设备id device_id
	device_id, ok := json["device_id"].(string)
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取设备ID",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析命令 command
	command, ok := json["command"].(string)
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取命令",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析命令参数 qos
	qos_val, ok := json["qos"].(float64)
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取QOS",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析命令参数 qos
	device_type, ok := json["device_type"].(string)
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取设备类型",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 转换 QOS 为int
	qos := int(qos_val)

	// 推送命令
	if err := global.Device.DeviceCmdPush(command, "mqtt", device_id, device_type, qos); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  fmt.Sprintf("推送命令失败: %s", err),
		})
		global.Logger.ERROR("推送命令失败: %s", err)
		return
	}

	global.Logger.INFO("推送 %s 设备 %s 指令, QOS等级 %d", device_id, command, qos)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "推送命令成功",
	})
}
