package device

import (
	"RisenIOT/backend/internal/agreement/unisound/lamp"
	"RisenIOT/backend/internal/emqx"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"sort"
	"strings"
)

type Device struct {
}

// Info 设备信息结构体
type Info struct {
	ProductID      string `json:"productID"`      // 产品ID
	DeviceUser     string `json:"deviceUser"`     // 设备用户名
	DeviceIP       string `json:"deviceIP"`       // 设备IP
	DeviceConnType string `json:"deviceConnType"` // 设备连接类型
}

// in 判断字符串是否在数组中
func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

// CreateDevice 创建设备实例
func CreateDevice() *Device {
	return &Device{}
}

// DeviceList 设备列表
func (d *Device) DeviceList() ([]Info, error) {
	var deviceList []Info

	if mqttDeviceList, err := emqx.Create().GetDeviceList(); err == nil {
		for _, v := range *mqttDeviceList {
			deviceList = append(deviceList, Info{
				ProductID:      v.DeviceId,
				DeviceUser:     v.Username,
				DeviceIP:       v.IpAddress,
				DeviceConnType: "MQTT",
			})
		}
	}

	return deviceList, nil

}

// Get 获取设备实例
func (d *Device) Get() *lamp.UnisoundLamp {
	return lamp.NewUnisoundLamp()
}

// SendHex 发送 HEX 数据
func (d *Device) SendHex(Topic, DeviceConnType string, data []byte) error {
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
