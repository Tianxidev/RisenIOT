package unisound

import (
	"RisenIOT/backend/controller/response"
	"RisenIOT/backend/global"
	"RisenIOT/backend/logger"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/gin-gonic/gin"
)

type LampController struct {
}

var topicPrefix = "/Lamp/TransIn/"

// LampOpenOrClose 灯开关接口
func (dc *LampController) LampOpenOrClose(context *gin.Context) {

	var err error
	var root ast.Node

	// 读取请求体
	RawData, _ := context.GetRawData()

	// 解析请求体
	if root, err = sonic.GetFromString(string(RawData)); err != nil {
		response.Error(context, 400, "参数错误, 无法解析json")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 设备id
	deviceId, _ := root.Get("device_id").String()

	// 命令
	command, _ := root.Get("command").String()

	// 通道
	chanVal, err := root.Get("chan").Int64()
	if err != nil {
		response.Error(context, 400, "参数错误, 无法读取通道")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 拼接订阅
	topic := topicPrefix + deviceId

	logger.GlobalLogger.INFO("接收到开关灯命令到订阅 %s, 通道: %d", topic, chanVal)

	// 判断是否是开灯命令
	if string(command) == "on" {

		err := global.Device.SendHex(topic, "MQTT", global.Device.Get().LampOnCmd(int(chanVal)))
		if err != nil {
			logger.GlobalLogger.ERROR("下发指令异常: %d", err)
			response.Error(context, 400, "下发指令异常")
			return
		}
	}

	// 判断是否是关灯命令
	if string(command) == "off" {
		err := global.Device.SendHex(topic, "MQTT", global.Device.Get().LampOffCmd(int(chanVal)))
		if err != nil {
			logger.GlobalLogger.ERROR("下发指令异常: %s", err)
			response.Error(context, 400, "下发指令异常")
			return
		}
	}

	response.Success(context, "下发指令成功", nil)

}

// LampDimming 调光接口
func (dc *LampController) LampDimming(context *gin.Context) {

	var err error
	var root ast.Node

	// 读取请求体
	RawData, _ := context.GetRawData()

	// 解析请求体
	if root, err = sonic.GetFromString(string(RawData)); err != nil {
		response.Error(context, 400, "参数错误, 无法解析json")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 设备id
	deviceId, _ := root.Get("device_id").String()

	// 亮度
	brightness, err := root.Get("brightness").Int64()
	if err != nil {
		response.Error(context, 400, "参数错误, 无法读取亮度")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 通道
	chanVal, err := root.Get("chan").Int64()
	if err != nil {
		response.Error(context, 400, "参数错误, 无法读取通道")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 拼接订阅
	topic := topicPrefix + deviceId

	logger.GlobalLogger.INFO("接收到调光命令到订阅 %s, 通道: %d", topic, chanVal)

	// 下发指令
	if brightness != 0 {
		err := global.Device.SendHex(topic, "MQTT", global.Device.Get().LampBrightnessCmd(int(brightness), int(chanVal)))
		if err != nil {
			response.Error(context, 400, "下发指令异常")
			logger.GlobalLogger.ERROR("下发指令异常: %s", err)
			return
		}
	}

	response.Success(context, "下发指令成功", nil)

}

// LampStatus 查询灯状态接口
func (dc *LampController) LampStatus(context *gin.Context) {

	var err error
	var root ast.Node

	// 读取请求体
	RawData, _ := context.GetRawData()

	// 解析请求体
	if root, err = sonic.GetFromString(string(RawData)); err != nil {
		response.Error(context, 400, "参数错误, 无法解析json")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 读取设备id
	deviceId, _ := root.Get("device_id").String()

	logger.GlobalLogger.INFO("接收到查询灯状态命令到订阅 %s", topicPrefix+deviceId)

	// 查询灯光状态
	//if status, err := global.Device.Get().LampStatus(deviceId); err == nil {
	//	logger.GlobalLogger.INFO("查询灯光状态成功")
	//	context.JSON(200, gin.H{
	//		"code": 200,
	//		"msg":  "查询灯光状态成功",
	//		"data": status,
	//	})
	//	return
	//} else {
	//	logger.GlobalLogger.ERROR("查询灯光状态失败: %s", err)
	//	context.JSON(200, gin.H{
	//		"code": 400,
	//		"msg":  "查询灯光状态失败",
	//	})
	//}

}
