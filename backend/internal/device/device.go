package device

import (
	"RisenIOT/backend/internal/agreement/unisound/lamp"
	"RisenIOT/backend/internal/emqx"
	"RisenIOT/backend/internal/redis"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/bytedance/sonic"
	"log"
	"strings"
)

type Device struct {
}

// Info 设备信息结构体
type Info struct {
	ProductID      string        `json:"productID"`      // 产品ID
	DeviceUser     string        `json:"deviceUser"`     // 设备用户名
	DeviceIP       string        `json:"deviceIP"`       // 设备IP
	DeviceConnType string        `json:"deviceConnType"` // 设备连接类型
	DeviceType     string        `json:"deviceType"`     // 设备类型
	DeviceInfo     []interface{} // 设备信息
}

// CreateDevice 创建设备实例
func CreateDevice() *Device {
	return &Device{}
}

// DeviceList 设备列表
func (d *Device) DeviceList() ([]Info, error) {
	var deviceList []Info

	// 读取 Redis 中的设备信息
	if DeviceListFromRedis, err := redis.NewRedis().Get("deviceList"); err == nil {
		err = sonic.Unmarshal([]byte(DeviceListFromRedis), &deviceList)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println(deviceList)

	// 读取 MQTT 设备列表
	if mqttDeviceList, err := emqx.Create().GetDeviceList(); err == nil {

		// 遍历设备列表
		for _, item := range *mqttDeviceList {
			deviceList = append(deviceList, Info{
				ProductID:      item.DeviceId,
				DeviceUser:     item.Username,
				DeviceIP:       item.IpAddress,
				DeviceConnType: "MQTT",
			})
		}
	}

	// list 转 json
	deviceListJson, _ := json.Marshal(deviceList)

	// 写入设备信息到 Redis
	redis.NewRedis().Set("deviceList", string(deviceListJson), 99999)

	return deviceList, nil

}

// Get 获取设备实例
func (d *Device) Get() *lamp.UnisoundLamp {
	return lamp.NewUnisoundLamp()
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
	if DeviceType == "unisound" && Agreement == "mqtt" {

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
