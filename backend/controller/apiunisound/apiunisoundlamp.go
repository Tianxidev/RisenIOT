package apiunisound

import (
	"RisenIOT/backend/agreement"
	"RisenIOT/backend/controller/apiresponse"
	"RisenIOT/backend/device"
	"RisenIOT/backend/pkg/logger"
	"RisenIOT/backend/utils"
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
		apiresponse.Error(context, 400, "参数错误, 无法解析json")
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
		apiresponse.Error(context, 400, "参数错误, 无法读取通道")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 拼接订阅
	topic := topicPrefix + deviceId

	logger.GlobalLogger.INFO("接收到开关灯命令到订阅 %s, 通道: %d", topic, chanVal)

	// 判断是否是开灯命令
	if string(command) == "on" {
		err := device.CreateDevice().SendHex(topic, "MQTT", agreement.NewUnisoundLamp().LampOnCmd(int(chanVal)))
		if err != nil {
			logger.GlobalLogger.ERROR("下发指令异常: %d", err)
			apiresponse.Error(context, 400, "下发指令异常")
			return
		}
	}

	// 判断是否是关灯命令
	if string(command) == "off" {
		err := device.CreateDevice().SendHex(topic, "MQTT", agreement.NewUnisoundLamp().LampOffCmd(int(chanVal)))
		if err != nil {
			logger.GlobalLogger.ERROR("下发指令异常: %s", err)
			apiresponse.Error(context, 400, "下发指令异常")
			return
		}
	}

	apiresponse.Success(context, nil, "下发指令成功")

}

// LampDimming 调光接口
func (dc *LampController) LampDimming(context *gin.Context) {

	var err error
	var root ast.Node

	// 读取请求体
	RawData, _ := context.GetRawData()

	// 解析请求体
	if root, err = sonic.GetFromString(string(RawData)); err != nil {
		apiresponse.Error(context, 400, "参数错误, 无法解析json")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 设备id
	deviceId, _ := root.Get("device_id").String()

	// 亮度
	brightness, err := root.Get("brightness").Int64()
	if err != nil {
		apiresponse.Error(context, 400, "参数错误, 无法读取亮度")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 通道
	chanVal, err := root.Get("chan").Int64()
	if err != nil {
		apiresponse.Error(context, 400, "参数错误, 无法读取通道")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 拼接订阅
	topic := topicPrefix + deviceId

	logger.GlobalLogger.INFO("接收到调光命令到订阅 %s, 通道: %d", topic, chanVal)

	// 下发指令
	err = device.CreateDevice().SendHex(topic, "MQTT", agreement.NewUnisoundLamp().LampBrightnessCmd(int(brightness), int(chanVal)))
	if err != nil {
		apiresponse.Error(context, 400, "下发指令异常")
		logger.GlobalLogger.ERROR("下发指令异常: %s", err)
		return
	}

	apiresponse.Success(context, nil, "下发指令成功")

}

// LampStatus 查询灯状态接口
func (dc *LampController) LampStatus(context *gin.Context) {

	var err error
	var root ast.Node

	// 读取请求体
	RawData, _ := context.GetRawData()

	// 解析请求体
	if root, err = sonic.GetFromString(string(RawData)); err != nil {
		apiresponse.Error(context, 400, "参数错误, 无法解析json")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 读取设备id
	deviceId, _ := root.Get("device_id").String()

	logger.GlobalLogger.INFO("接收到查询灯状态命令到订阅 %s", topicPrefix+deviceId)

	// 查询设备状态 redis
	if deviceInfo, err := device.CreateDevice().GetDeviceInfo(deviceId); err == nil {
		logger.GlobalLogger.INFO("查询设备状态成功")
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "查询设备状态成功",
			"data": deviceInfo,
		})
		return
	} else {
		logger.GlobalLogger.ERROR("查询设备状态失败: %s", err)
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "查询设备状态失败",
		})
	}

}

// SetLocation 设置经纬度接口
func (dc *LampController) SetLocation(context *gin.Context) {

	var err error
	var root ast.Node

	// 读取请求体
	RawData, _ := context.GetRawData()

	// 解析请求体
	if root, err = sonic.GetFromString(string(RawData)); err != nil {
		apiresponse.Error(context, 400, "参数错误, 无法解析json")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 读取设备id
	deviceId, _ := root.Get("device_id").String()

	// 读取经度
	longitude, err := root.Get("longitude").String()

	// 读取维度
	latitude, err := root.Get("latitude").String()

	// 查询设备状态 redis
	if deviceInfo, err := device.CreateDevice().GetDeviceInfo(deviceId); err == nil {

		// 更新经度
		deviceInfo["longitude"] = longitude

		// 更新纬度
		deviceInfo["latitude"] = latitude

		// 更新设备信息
		err := device.CreateDevice().UpdateDeviceInfo(deviceId, deviceInfo)
		if err != nil {
			logger.GlobalLogger.ERROR("更新设备信息失败: %s", err)
		}

		// 转换经度为 int
		longitudeInt := utils.StringToInt(longitude)

		// 转换纬度为 int
		latitudeInt := utils.StringToInt(latitude)

		// 下发经纬度到设备
		err = device.CreateDevice().SendHex(topicPrefix+deviceId, "MQTT", agreement.NewUnisoundLamp().SetLocationCmd(longitudeInt, latitudeInt))
		if err != nil {
			logger.GlobalLogger.ERROR("更新设备信息失败: %s", err)
		}

		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "更新经纬度状态成功",
			"data": deviceInfo,
		})
		return
	} else {
		logger.GlobalLogger.ERROR("查询设备状态失败: %s", err)
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "查询设备状态失败",
		})
	}

}
