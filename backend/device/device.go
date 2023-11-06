package device

import (
	"RisenIOT/backend/emqx"
	"RisenIOT/backend/pkg/logger"
	"RisenIOT/backend/redis"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/goccy/go-json"
	"log"
	"strings"
)

type Device struct {
}

// Info 设备信息结构体
type Info struct {
	ProductID      string      `json:"productID"`      // 产品ID
	DeviceUser     string      `json:"deviceUser"`     // 设备用户名
	DeviceIP       string      `json:"deviceIP"`       // 设备IP
	DeviceConnType string      `json:"deviceConnType"` // 设备连接类型
	DeviceType     string      `json:"deviceType"`     // 设备类型
	DeviceInfo     interface{} // 设备信息
}

// CreateDevice 创建设备实例
func CreateDevice() *Device {
	return &Device{}
}

// DeviceList 设备列表
func (d *Device) DeviceList() ([]Info, error) {
	var deviceList []Info

	// 读取 MQTT 设备列表
	if mqttDeviceList, err := emqx.Create().GetDeviceList(); err == nil {

		// 遍历设备列表
		for _, item := range *mqttDeviceList {
			var device Info

			// 查询 Redis 中是否存在设备信息
			if deviceInfo, err := redis.NewRedis().Get(item.DeviceId); err == nil {
				if deviceInfo != "" {
					log.Printf("从 Redis 中获取设备信息: %s", deviceInfo)
					err := json.Unmarshal([]byte(deviceInfo), &device)
					if err != nil {
						log.Printf("解析设备信息失败: %s", err)
					}
				}
			}

			// 更新设备信息
			device.ProductID = item.DeviceId
			device.DeviceUser = item.Username
			device.DeviceIP = item.IpAddress
			device.DeviceConnType = "MQTT"

			deviceList = append(deviceList, device)

			// api_device 转 json
			deviceJson, _ := json.Marshal(device)

			// 写入设备信息到 Redis
			redis.NewRedis().Set(device.ProductID, string(deviceJson), -1)
		}
	}

	return deviceList, nil

}

// SendHex 发送 HEX 数据
func (d *Device) SendHex(Topic, DeviceConnType string, data []byte) error {

	if Topic == "" || DeviceConnType == "" || data == nil {
		return nil
	}

	if DeviceConnType == "MQTT" {
		Payload := base64.StdEncoding.EncodeToString(data)
		mqtt := emqx.Create()
		err := mqtt.SendDataToTopic(Payload, "base64", Topic, 1)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// DeviceCmdPush 设备命令推送
func (d *Device) DeviceCmdPush(Payload string, Agreement string, DeviceId string, DeviceType string, Qos int) error {

	// 去除空格
	Payload = strings.ReplaceAll(Payload, " ", "")

	// 判断设备类型是否是[云知声]
	if DeviceType == "ApiUnisound" && Agreement == "mqtt" {

		var data []byte

		// 将 HEX 编码的字符串转换为 HEX 数据
		data1, _ := hex.DecodeString(Payload)
		for _, v := range data1 {
			data = append(data, v)
		}

		// 将 HEX 数据转换为 Base64 编码的字符串
		sEnc := base64.StdEncoding.EncodeToString([]byte(data))
		Payload = string(sEnc)
		mqtt := emqx.Create()
		err := mqtt.SendDataToTopic(Payload, "base64", DeviceId, 1)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("不支持的设备类型")
}

// UpdateDeviceInfo 更新设备信息
func (d *Device) UpdateDeviceInfo(deviceId string, deviceInfo map[string]interface{}) error {
	var device Info

	// 查询 Redis 中是否存在设备信息
	if deviceInfo, err := redis.NewRedis().Get(deviceId); err == nil {
		if deviceInfo != "" {
			err := json.Unmarshal([]byte(deviceInfo), &device)
			if err != nil {
				logger.GlobalLogger.ERROR("解析设备信息失败: %s", err)
			}
		}
	}

	// 更新设备信息
	device.DeviceInfo = deviceInfo

	// api_device 转 json
	deviceJson, _ := json.Marshal(device)

	// 写入设备信息到 Redis
	redis.NewRedis().Set(device.ProductID, string(deviceJson), -1)

	return nil
}

// GetDeviceInfo 获取设备信息
func (d *Device) GetDeviceInfo(deviceId string) (map[string]interface{}, error) {
	var device Info

	// 查询 Redis 中是否存在设备信息
	if deviceInfo, err := redis.NewRedis().Get(deviceId); err == nil {
		if deviceInfo != "" {
			err := json.Unmarshal([]byte(deviceInfo), &device)
			if err != nil {
				logger.GlobalLogger.ERROR("解析设备信息失败: %s", err)
			}
		}
	}

	// 判空设备信息
	if device.DeviceInfo == nil {
		return nil, errors.New("设备信息为空")
	}

	return device.DeviceInfo.(map[string]interface{}), nil

}
