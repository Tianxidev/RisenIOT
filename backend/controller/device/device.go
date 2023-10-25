package device

import (
	"RisenIOT/backend/agreement/unisound/lamp"
	"RisenIOT/backend/controller/response"
	"RisenIOT/backend/global"
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
		response.Error(context, 400, "参数错误, 无法解析json")
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 创建协议识别结果管道
	protocolChan := make(chan string)

	// 判断是否是 [云知声灯控] 协议
	go lamp.NewUnisoundLamp().TopicHandler(root, protocolChan)

	// 读取管道
	protocol := <-protocolChan

	global.Logger.LOG("EMQX", "%s", protocol)

	response.Success(context, "接收成功", nil)

	return
}

// DeviceList 设备列表
func (dc *Controller) DeviceList(context *gin.Context) {
	if list, err := global.Device.DeviceList(); err == nil {
		global.Logger.INFO("获取设备列表成功")
		context.JSON(200, gin.H{
			"code":  200,
			"msg":   "获取设备列表成功",
			"total": len(list),
			"data":  list,
		})
		return
	} else {
		global.Logger.ERROR("获取设备列表失败: %s", err)
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

	// 解析命令参数 qos
	qos_val, ok := json["qos"].(float64)
	if !ok {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误, 无法读取QOS",
		})
		global.Logger.ERROR("参数错误: %s", err)
		return
	}

	// 解析命令参数 qos
	device_type, ok := json["device_type"].(string)
	if !ok {
		context.JSON(200, gin.H{
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
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  fmt.Sprintf("推送命令失败: %s", err),
		})
		global.Logger.ERROR("推送命令失败: %s", err)
		return
	}

	global.Logger.INFO("推送 %s 设备 %s 指令, QOS等级 %d", device_id, command, qos)
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "推送命令成功",
	})
}
