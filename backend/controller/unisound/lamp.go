package unisound

import (
	"RisenIOT/backend/controller/response"
	"RisenIOT/backend/global"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fastjson"
)

type LampController struct {
}

var topicPrefix = "/Lamp/TransIn/"

// LampOpenOrClose 灯开关接口
func (dc *LampController) LampOpenOrClose(context *gin.Context) {

	var err error
	var p fastjson.Parser
	var v *fastjson.Value

	// 反序列化请求体
	data, _ := context.GetRawData()
	if v, err = p.Parse(string(data)); err != nil {
		response.Error(context, 400, "参数错误, 无法解析json")
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 设备id
	deviceId := v.GetStringBytes("device_id")

	// 命令
	command := v.GetStringBytes("command")

	// 通道
	chanVal, err := v.Get("chan").Int()
	if err != nil {
		response.Error(context, 400, "参数错误, 无法读取通道")
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 拼接订阅
	topic := topicPrefix + string(deviceId)

	global.Logger.INFO("接收到开关灯命令到订阅 %s, 通道: %d", topic, chanVal)

	// 判断是否是开灯命令
	if string(command) == "on" {

		err := global.Device.SendHex(topic, "MQTT", global.Device.Get().LampOnCmd(chanVal))
		if err != nil {
			global.Logger.ERROR("下发指令异常: %d", err)
			response.Error(context, 400, "下发指令异常")
			return
		}
	}

	// 判断是否是关灯命令
	if string(command) == "off" {
		err := global.Device.SendHex(topic, "MQTT", global.Device.Get().LampOffCmd(chanVal))
		if err != nil {
			global.Logger.ERROR("下发指令异常: %s", err)
			response.Error(context, 400, "下发指令异常")
			return
		}
	}

	response.Success(context, "下发指令成功", nil)

}

// LampDimming 调光接口
func (dc *LampController) LampDimming(context *gin.Context) {

	var err error
	var p fastjson.Parser
	var v *fastjson.Value

	// 反序列化请求体
	data, _ := context.GetRawData()
	if v, err = p.Parse(string(data)); err != nil {
		response.Error(context, 400, "参数错误, 无法解析json")
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 设备id
	deviceId := v.GetStringBytes("device_id")

	// 亮度
	brightness, err := v.Get("brightness").Int()
	if err != nil {
		response.Error(context, 400, "参数错误, 无法读取亮度")
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 通道
	chanVal, err := v.Get("chan").Int()
	if err != nil {
		response.Error(context, 400, "参数错误, 无法读取通道")
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 拼接订阅
	topic := topicPrefix + string(deviceId)

	global.Logger.INFO("接收到调光命令到订阅 %s, 通道: %d", topic, chanVal)

	// 下发指令
	if brightness != 0 {
		err := global.Device.SendHex(topic, "MQTT", global.Device.Get().LampBrightnessCmd(brightness, chanVal))
		if err != nil {
			response.Error(context, 400, "下发指令异常")
			global.Logger.ERROR("下发指令异常: %s", err)
			return
		}
	}

	response.Success(context, "下发指令成功", nil)

}
