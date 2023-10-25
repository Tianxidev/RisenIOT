package lamp

import (
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
	DeviceId   string `json:"device_id"`   // 设备ID
	DeviceType string `json:"device_type"` // 设备类型
	ChannelA   int    `json:"channel_a"`   // A 通道亮度
	ChannelB   int    `json:"channel_b"`   // B 通道亮度
}

// NewUnisoundLamp 创建云知声灯控协议实例
func NewUnisoundLamp() *UnisoundLamp {
	return &UnisoundLamp{}
}

// TopicHandler 订阅处理器
func (u *UnisoundLamp) TopicHandler(jsonRoot ast.Node, protocol chan string) {

	topic, _ := jsonRoot.Get("topic").String()
	clientId, _ := jsonRoot.Get("clientid").String()

	// 示例订阅格式: /Lamp/TransOut/868247062445885
	re1 := regexp.MustCompile(`/Lamp/TransOut/(\d{15})`)

	// 示例订阅格式: /Lamp/TransIn/868247062445885
	re2 := regexp.MustCompile(`/Lamp/TransIn/(\d{15})`)

	// 匹配订阅
	match1 := re1.FindStringSubmatch(topic)
	match2 := re2.FindStringSubmatch(topic)

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
			log.Printf("[ 云知声灯控 ] CRC 校验通过: %X => %X", ulc.CRC, crc)
		} else {
			log.Printf("[ 云知声灯控 ] CRC 校验失败: %X => %X", ulc.CRC, crc)
			return
		}

		// 协议识别 - 读取双灯开关亮度状态 - 03 功能码 - 4005 寄存器 - 01 长度
		if ulc.Cmd[0] == 0x03 && ulc.Data[0] == 0x40 && ulc.Data[1] == 0x05 && ulc.Data[2] == 0x01 {

			var uli UnisoundLampInfo

			// 更新设备信息
			uli.DeviceId = clientId
			uli.DeviceType = "UnisoundLamp"
			uli.ChannelA = int(ulc.Data[3])
			uli.ChannelB = int(ulc.Data[4])

			jsonStr, _ := json.Marshal(uli)

			log.Printf("[ 云知声灯控 ] 读取设备 %s 到 A 通道亮度: %d , B 通道亮度: %d", clientId, ulc.Data[3], ulc.Data[4])

			// 返回识别结果
			protocol <- string(jsonStr)
			return
		}

		// 未识别协议
		log.Printf("[ 云知声灯控 ] 接收设备 %s 上报的未知数据: %s", clientId, data)
		return

	}

	// 检查匹配结果2
	if len(match2) > 1 && match2[1] == clientId {
		data, _ := jsonRoot.Get("payload_hexstr").String()
		protocol <- fmt.Sprintf("[ 云知声灯控 ] 下发到设备 %s 的数据: %s", clientId, data)
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
		log.Printf("[ 云知声灯控 ] 通道不合法: %d", channel)
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

		log.Printf("[ 云知声灯控 ] 通道 a 开灯命令: %X", cmd)

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

		log.Printf("[ 云知声灯控 ] 通道 b 开灯命令: %X", cmd)

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
		log.Printf("[ 云知声灯控 ] 通道不合法: %d", channel)
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

		log.Printf("[ 云知声灯控 ] 通道 a 关灯命令: %X", cmd)

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

		log.Printf("[ 云知声灯控 ] 通道 b 关灯命令: %X", cmd)

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

		log.Printf("[ 云知声灯控 ] 通道 a 调光命令: %X", cmd)

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

		log.Printf("[ 云知声灯控 ] 通道 b 调光命令: %X", cmd)

		return cmd
	}

	return []byte{}
}
