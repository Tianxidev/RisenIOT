package agreement

import (
	"RisenIOT/backend/device"
	"RisenIOT/backend/logger"
	"RisenIOT/backend/utils"
	"encoding/hex"
	"fmt"
	"github.com/bytedance/sonic/ast"
	"github.com/goccy/go-json"
	"log"
	"regexp"
)

// UnisoundLamp 云之声灯控协议支持
type UnisoundLamp struct {
}

type unisoundLampCmd struct {
	Random []byte // 随机码, 2字节
	Addr   []byte // 设备地址, 1字节, 范围：01H～F7H(十进制 1～247)
	Cmd    []byte // 功能码, 1字节,不同的功能对应不同的值, 03H:读单个或多个字寄存器; 06H: 写单个字寄存器; 10H: 写连续N个寄存器
	Data   []byte // 数据, 具体和对应命令有关, 比如开关灯, 调节亮度, 读系统状态就是 2 个字节地址 + 2 字节数据组成, 不同的命令数据码不一样
	CRC    []byte // CRC-16/MODBUS 校验码, 2字节, 是对地址码+功能码+数据码的计算得到
}

type UnisoundLampInfo struct {
	DeviceType string `json:"device_type"` // 设备类型
	ChannelA   int    `json:"channel_a"`   // A 通道亮度
	ChannelB   int    `json:"channel_b"`   // B 通道亮度
	Longitude  string `json:"longitude"`   // 经度
	Latitude   string `json:"latitude"`    // 纬度
}

// NewUnisoundLamp 创建云知声灯控协议实例
func NewUnisoundLamp() *UnisoundLamp {
	return &UnisoundLamp{}
}

// TopicHandler 订阅处理器
func (u *UnisoundLamp) TopicHandler(jsonRoot ast.Node) {

	topic, _ := jsonRoot.Get("topic").String()
	clientId, _ := jsonRoot.Get("clientid").String()

	// 示例订阅格式: /Lamp/TransOut/868247062445885
	re1 := regexp.MustCompile(`/Lamp/TransOut/(\d{15})`)

	// 匹配订阅
	match1 := re1.FindStringSubmatch(topic)

	// 检查匹配结果1
	if len(match1) > 1 && match1[1] == clientId {

		var ulc unisoundLampCmd

		data, _ := jsonRoot.Get("payload_hexstr").String()

		// 返回协议示例: 00FF 0103 4005 0101 145A AF
		dataHex, _ := hex.DecodeString(data)

		// 构建协议 末尾2字节为校验码
		ulc.Random = dataHex[0:2]
		ulc.Addr = dataHex[2:3]
		ulc.Cmd = dataHex[3:4]
		ulc.Data = dataHex[4 : len(dataHex)-2]
		ulc.CRC = dataHex[len(dataHex)-2:]

		// 检查 CRC 校验码
		crc := utils.CRC16(dataHex[2 : len(dataHex)-2])
		if ulc.CRC[0] != byte(crc&0xff) || ulc.CRC[1] == byte(crc>>8) {
			logger.GlobalLogger.INFO("[ 云知声灯控 ] CRC 校验通过: %X => %X", ulc.CRC, crc)
		} else {
			logger.GlobalLogger.INFO("[ 云知声灯控 ] CRC 校验失败: %X => %X", ulc.CRC, crc)
			return
		}

		// 协议识别 - 读取设备经纬度 - 示例返回: 43F6 0103 DF12 05 6340B380 1D253F60 0008 CB8B
		// 6340B380：经度
		// 1D253F60：纬度
		// 0008：时区
		if ulc.Cmd[0] == 0x03 && ulc.Data[0] == 0xDF && ulc.Data[1] == 0x12 && ulc.Data[2] == 0x05 {

			var uli UnisoundLampInfo

			// 读取设备信息
			deviceInfo, _ := device.CreateDevice().GetDeviceInfo(clientId)
			deviceInfoJson, _ := json.Marshal(deviceInfo)
			_ = json.Unmarshal(deviceInfoJson, &uli)

			// 转换出经度int
			longitudeInt := utils.BToU32(ulc.Data, 3, 10)

			// 转换出纬度int
			latitudeInt := utils.BToU32(ulc.Data, 11, 18)

			// 经度
			uli.Longitude = fmt.Sprintf("%d", longitudeInt)

			// 纬度
			uli.Latitude = fmt.Sprintf("%d", latitudeInt)

			// 转换为map
			uliJson, _ := json.Marshal(uli)
			var deviceInfoMap map[string]interface{}
			_ = json.Unmarshal(uliJson, &deviceInfoMap)

			// 更新设备信息
			err := device.CreateDevice().UpdateDeviceInfo(clientId, deviceInfoMap)
			if err != nil {
				logger.GlobalLogger.ERROR("更新设备信息失败")
			}

			logger.GlobalLogger.INFO("[ 云知声灯控 ] 设备 %s 经度: %d, 纬度: %d", clientId, longitudeInt, latitudeInt)

			return
		}

		// 协议识别 - 读取双灯开关亮度状态 - 03 功能码 - 4005 寄存器 - 01 长度
		if ulc.Cmd[0] == 0x03 && ulc.Data[0] == 0x40 && ulc.Data[1] == 0x05 && ulc.Data[2] == 0x01 {

			var uli UnisoundLampInfo

			// 读取设备信息
			deviceInfo, err := device.CreateDevice().GetDeviceInfo(clientId)
			deviceInfoJson, _ := json.Marshal(deviceInfo)
			_ = json.Unmarshal(deviceInfoJson, &uli)

			// 更新设备信息
			uli.DeviceType = "云知声灯控"
			uli.ChannelA = int(ulc.Data[3])
			uli.ChannelB = int(ulc.Data[4])

			// 转换为map
			uliJson, _ := json.Marshal(uli)
			var deviceInfoMap map[string]interface{}
			_ = json.Unmarshal(uliJson, &deviceInfoMap)

			logger.GlobalLogger.INFO("[ 云知声灯控 ] 读取设备 %s 到 A 通道亮度: %d , B 通道亮度: %d", clientId, ulc.Data[3], ulc.Data[4])

			// 更新设备信息
			err = device.CreateDevice().UpdateDeviceInfo(clientId, deviceInfoMap)
			if err != nil {
				logger.GlobalLogger.ERROR("更新设备信息失败")
			}

			return
		}

		// 未识别协议
		logger.GlobalLogger.INFO("[ 云知声灯控 ] 接收设备 %s 上报的未知数据: %s", clientId, data)
		return

	}

}

// randomGenerate 随机码生成
func (u *UnisoundLamp) randomGenerate() []byte {
	return []byte{0xff, 0x00}
}

// LampOnCmd 开灯命令
func (u *UnisoundLamp) LampOnCmd(channel int) []byte {
	var ulc unisoundLampCmd
	var cmd []byte

	// 检查通道是否合法
	if channel < 1 || channel > 3 {
		logger.GlobalLogger.ERROR("[ 云知声灯控 ] 通道不合法: %d", channel)
		return []byte{}
	}

	// 通道 a - 举例下发给设备的开灯命令：FF00 0106 DF09 01A0 63F4
	if channel == 1 {
		ulc.Random = u.randomGenerate()
		ulc.Addr = []byte{0x01}
		ulc.Cmd = []byte{0x06}
		ulc.Data = []byte{0xDF, 0x09, 0x01, 0xA0}

		// 构建命令
		cmd = append(cmd, ulc.Random...)
		cmd = append(cmd, ulc.Addr...)
		cmd = append(cmd, ulc.Cmd...)
		cmd = append(cmd, ulc.Data...)

		// 忽略随机码部分计算 CRC 校验码
		crc := utils.CRC16(cmd[2:])
		ulc.CRC = []byte{byte(crc & 0xff), byte(crc >> 8)}
		cmd = append(cmd, ulc.CRC...)

		logger.GlobalLogger.INFO("[ 云知声灯控 ] 通道 a 开灯命令: %X", cmd)

		return cmd
	}

	// 通道 b - 举例下发给设备的开灯命令：FF00 0106 DF0B 01A0 C234
	if channel == 2 {

		ulc.Random = u.randomGenerate()
		ulc.Addr = []byte{0x01}
		ulc.Cmd = []byte{0x06}
		ulc.Data = []byte{0xDF, 0x0B, 0x01, 0xA0}

		// 构建命令
		cmd = append(cmd, ulc.Random...)
		cmd = append(cmd, ulc.Addr...)
		cmd = append(cmd, ulc.Cmd...)
		cmd = append(cmd, ulc.Data...)

		// 忽略随机码部分计算 CRC 校验码
		crc := utils.CRC16(cmd[2:])
		ulc.CRC = []byte{byte(crc & 0xff), byte(crc >> 8)}
		cmd = append(cmd, ulc.CRC...)

		logger.GlobalLogger.INFO("[ 云知声灯控 ] 通道 b 开灯命令: %X", cmd)

		return cmd
	}

	return []byte{}

}

// LampOffCmd 关灯命令
func (u *UnisoundLamp) LampOffCmd(channel int) []byte {
	var ulc unisoundLampCmd
	var cmd []byte

	// 检查通道是否合法
	if channel < 1 || channel > 3 {
		logger.GlobalLogger.ERROR("[ 云知声灯控 ] 通道不合法: %d", channel)
		return []byte{}
	}

	// 通道 a - 举例下发给设备的关灯命令：FF 00 01 06 DF 09 01 5F 23 B4
	if channel == 1 {
		ulc.Random = u.randomGenerate()
		ulc.Addr = []byte{0x01}
		ulc.Cmd = []byte{0x06}
		ulc.Data = []byte{0xDF, 0x09, 0x01, 0x5F}

		// 构建命令
		cmd = append(cmd, ulc.Random...)
		cmd = append(cmd, ulc.Addr...)
		cmd = append(cmd, ulc.Cmd...)
		cmd = append(cmd, ulc.Data...)

		// 忽略随机码部分计算 CRC 校验码
		crc := utils.CRC16(cmd[2:])
		ulc.CRC = []byte{byte(crc & 0xff), byte(crc >> 8)}
		cmd = append(cmd, ulc.CRC...)

		logger.GlobalLogger.INFO("[ 云知声灯控 ] 通道 a 关灯命令: %X", cmd)

		return cmd
	}

	// 通道 b - 举例下发给设备的关灯命令：FF 00 01 06 DF 0B 01 5F 82 74
	if channel == 2 {

		ulc.Random = u.randomGenerate()
		ulc.Addr = []byte{0x01}
		ulc.Cmd = []byte{0x06}
		ulc.Data = []byte{0xDF, 0x0B, 0x01, 0x5F}

		// 构建命令
		cmd = append(cmd, ulc.Random...)
		cmd = append(cmd, ulc.Addr...)
		cmd = append(cmd, ulc.Cmd...)
		cmd = append(cmd, ulc.Data...)

		// 忽略随机码部分计算 CRC 校验码
		crc := utils.CRC16(cmd[2:])
		ulc.CRC = []byte{byte(crc & 0xff), byte(crc >> 8)}
		cmd = append(cmd, ulc.CRC...)

		logger.GlobalLogger.INFO("[ 云知声灯控 ] 通道 b 关灯命令: %X", cmd)

		return cmd
	}

	return []byte{}
}

// LampBrightnessCmd 调节亮度命令
func (u *UnisoundLamp) LampBrightnessCmd(brightness int, channel int) []byte {

	var ulc unisoundLampCmd
	var cmd []byte

	// 检查通道是否合法
	if channel < 1 || channel > 3 {
		log.Printf("[ 云知声灯控 ] 通道不合法: %d", channel)
		return []byte{}
	}

	// 通道 a - 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0A 00 00 92 1C -> 亮度0%
	// 通道 a - 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0A 00 14 92 13 -> 亮度20%
	// 通道 a - 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0A 00 64 93 F7 -> 亮度100%
	if channel == 1 {

		ulc.Random = u.randomGenerate()
		ulc.Addr = []byte{0x01}
		ulc.Cmd = []byte{0x06}
		ulc.Data = []byte{0xDF, 0x0A, 0x00, byte(brightness)}

		// 构建命令
		cmd = append(cmd, ulc.Random...)
		cmd = append(cmd, ulc.Addr...)
		cmd = append(cmd, ulc.Cmd...)
		cmd = append(cmd, ulc.Data...)

		// 忽略随机码部分计算 CRC 校验码
		crc := utils.CRC16(cmd[2:])
		ulc.CRC = []byte{byte(crc & 0xff), byte(crc >> 8)}
		cmd = append(cmd, ulc.CRC...)

		logger.GlobalLogger.INFO("[ 云知声灯控 ] 通道 a 调光命令: %X", cmd)

		return cmd
	}

	// 通道 b - 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0C 00 00 72 1D -> 亮度0%
	// 通道 b - 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0C 00 14 72 12 -> 亮度20%
	// 通道 b - 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0C 00 64 73 F6 -> 亮度100%
	if channel == 2 {

		ulc.Random = u.randomGenerate()
		ulc.Addr = []byte{0x01}
		ulc.Cmd = []byte{0x06}
		ulc.Data = []byte{0xDF, 0x0C, 0x00, byte(brightness)}

		// 构建命令
		cmd = append(cmd, ulc.Random...)
		cmd = append(cmd, ulc.Addr...)
		cmd = append(cmd, ulc.Cmd...)
		cmd = append(cmd, ulc.Data...)

		// 忽略随机码部分计算 CRC 校验码
		crc := utils.CRC16(cmd[2:])
		ulc.CRC = []byte{byte(crc & 0xff), byte(crc >> 8)}
		cmd = append(cmd, ulc.CRC...)

		logger.GlobalLogger.INFO("[ 云知声灯控 ] 通道 b 调光命令: %X", cmd)

		return cmd
	}

	return []byte{}
}

// SetLocationCmd 设置经纬度
func (u *UnisoundLamp) SetLocationCmd(longitude []byte, latitude []byte) []byte {

	var ulc unisoundLampCmd
	var cmd []byte

	// 举例下发给设备的经纬度命令：5A0A 0110 DF12 05 6340B380 1D253F60 0008 D958
	// 6340B380: 经度
	// 1D253F60: 纬度
	// 0008: 时区
	ulc.Random = u.randomGenerate()
	ulc.Addr = []byte{0x01}
	ulc.Cmd = []byte{0x10}
	ulc.Data = []byte{0xDF, 0x12, 0x05}

	ulc.Data = append(ulc.Data, longitude...)
	ulc.Data = append(ulc.Data, latitude...)
	ulc.Data = append(ulc.Data, 0x00)
	ulc.Data = append(ulc.Data, 0x08)

	// 忽略随机码部分计算 CRC 校验码
	crc := utils.CRC16(cmd[2:])
	ulc.CRC = []byte{byte(crc & 0xff), byte(crc >> 8)}
	cmd = append(cmd, ulc.CRC...)

	log.Printf("[ 云知声灯控 ] 设置经纬度命令: %X", cmd)

	return cmd

}
