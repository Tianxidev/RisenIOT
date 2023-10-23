package unisound

import (
	"RisenIOT/backend/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

type LampController struct {
}

// LampOpenOrClose 灯开关接口
func (dc *LampController) LampOpenOrClose(context *gin.Context) {

	// 读取post json
	var json map[string]interface{}

	err := context.BindJSON(&json)
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法解析json",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析设备id device_id
	device_id, ok := json["device_id"].(string)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取设备ID",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析命令 command
	command, ok := json["command"].(string)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取命令",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 判断时候是开灯命令
	if command == "on" {
		err := global.Device.SendHex(device_id, "MQTT", global.Device.Get().LampOnCmd())
		if err != nil {
			global.Logger.ERROR("下发指令异常: %s", err)
		}
	}

	if command == "off" {
		err := global.Device.SendHex(device_id, "MQTT", global.Device.Get().LampOffCmd())
		if err != nil {
			global.Logger.ERROR("下发指令异常: %s", err)
		}
	}

	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "下发指令成功",
	})

}

// LampDimming 调光接口
func (dc *LampController) LampDimming(context *gin.Context) {
	// 读取post json
	var json map[string]interface{}

	err := context.BindJSON(&json)
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法解析json",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析设备id device_id
	device_id, ok := json["device_id"].(string)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取设备ID",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析强度 rssi
	rssi, ok := json["rssi"].(string)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取命令",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	if rssi != "" {
		// 强度字符串转int
		RSSI, _ := strconv.Atoi(rssi)
		err := global.Device.SendHex(device_id, "MQTT", global.Device.Get().LampBrightnessCmd(RSSI))
		if err != nil {
			global.Logger.ERROR("下发指令异常: %s", err)
		}
	}

	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "下发指令成功",
	})

}
