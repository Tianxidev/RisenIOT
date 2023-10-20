package device

import (
	"RisenIOT/backend/internal/emqx"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Device struct {
}

type Core struct {
	ProductID  string `json:"productID"`  // 产品id
	DeviceName string `json:"deviceName"` // 设备名称
}

// 设备类型数组
var deviceTypeArray = []string{"unisound"}

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

// DeviceCmdPush 设备命令推送
func (d *Device) DeviceCmdPush(Payload string, Agreement string, DeviceId string, DeviceType string, Qos int) error {

	// 去除空格
	Payload = strings.ReplaceAll(Payload, " ", "")

	// 判断设备类型是否支持
	if !in(DeviceType, deviceTypeArray) {
		return errors.New("不支持的设备类型")
	}

	// 判断设备类型是否是[云知声]
	if DeviceType == "unisound" {

		fmt.Println("云知声设备")

		var data []byte

		// 将 HEX 编码的字符串转换为 HEX 数据
		data1, _ := hex.DecodeString(Payload)
		for _, v := range data1 {
			data = append(data, v)
		}

		fmt.Println("发送消息: ", data)

		// 将 HEX 数据转换为 Base64 编码的字符串
		sEnc := base64.StdEncoding.EncodeToString([]byte(data))
		Payload = string(sEnc)
	}

	// 判断传输协议类型
	if Agreement == "mqtt" {
		mqtt := emqx.Create()
		err := mqtt.SendDataToTopic(Payload, "base64", DeviceId, 1)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("不支持的设备类型")
}
