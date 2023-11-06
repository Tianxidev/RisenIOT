package ApiDevice

import (
	"RisenIOT/backend/agreement"
	"RisenIOT/backend/controller/ApiResponse"
	"RisenIOT/backend/device"
	"RisenIOT/backend/pkg/logger"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

// ReceiveDataFromEmqxWebHook 设备数据接收 EMQX WebHook
func (dc *Controller) ReceiveDataFromEmqxWebHook(context *gin.Context) {
	var err error
	var root ast.Node

	// 读取请求体
	RawData, _ := context.GetRawData()

	// 解析请求体
	if root, err = sonic.GetFromString(string(RawData)); err != nil {
		ApiResponse.Error(context, 400, "参数错误, 无法解析json")
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 判断是否是 [云知声灯控] 协议
	go agreement.NewUnisoundLamp().TopicHandler(root)

	// 打印协议报文
	data, _ := root.Get("payload_hexstr").String()
	topic, _ := root.Get("topic").String()
	clientid, _ := root.Get("clientid").String()
	logger.GlobalLogger.LOG("EMQX", "接收到设备 %s 发送到订阅 %s 的数据: %s", clientid, topic, data)

	ApiResponse.Success(context, "接收成功", nil)

	return
}

// DeviceList 设备列表
func (dc *Controller) DeviceList(context *gin.Context) {
	if list, err := device.CreateDevice().DeviceList(); err == nil {
		logger.GlobalLogger.INFO("获取设备列表成功")
		context.JSON(200, gin.H{
			"code":  200,
			"msg":   "获取设备列表成功",
			"total": len(list),
			"data":  list,
		})
		return
	} else {
		logger.GlobalLogger.ERROR("获取设备列表失败: %s", err)
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  fmt.Sprintf("获取设备列表失败: %s", err),
		})
	}
	return
}

// DeviceCmdPush 设备命令透传推送
func (dc *Controller) DeviceCmdPush(context *gin.Context) {

	// 读取post json
	var json map[string]interface{}

	err := context.BindJSON(&json)
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法解析json",
		})
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 解析设备id device_id
	device_id, ok := json["device_id"].(string)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取设备ID",
		})
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 解析命令 command
	command, ok := json["command"].(string)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取命令",
		})
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 解析命令参数 qos
	qos_val, ok := json["qos"].(float64)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取QOS",
		})
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 解析命令参数 qos
	device_type, ok := json["device_type"].(string)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取设备类型",
		})
		logger.GlobalLogger.ERROR("参数错误: %s", err)
		return
	}

	// 转换 QOS 为int
	qos := int(qos_val)

	// 推送命令
	if err := device.CreateDevice().DeviceCmdPush(command, "mqtt", device_id, device_type, qos); err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  fmt.Sprintf("推送命令失败: %s", err),
		})
		logger.GlobalLogger.ERROR("推送命令失败: %s", err)
		return
	}

	logger.GlobalLogger.INFO("推送 %s 设备 %s 指令, QOS等级 %d", device_id, command, qos)
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "推送命令成功",
	})
}
