package agreement

import (
	"RisenIOT/backend/utils"
	"fmt"
)

// UnisoundLamp 云之声灯控协议支持
type UnisoundLamp struct {
}

type unisoundLampCmd struct {
	Random []byte // 随机码, 2字节
	Addr   []byte // 设备地址, 1字节, 范围：01H～F7H(十进制 1～247)
	Cmd    []byte // 命令码, 1字节,不同的功能对应不同的值, 03H:读单个或多个字寄存器; 06H: 写单个字寄存器; 10H: 写连续N个寄存器
	Data   []byte // 数据, 具体和对应命令有关, 比如开关灯, 调节亮度, 读系统状态就是 2 个字节地址 + 2 字节数据组成, 不同的命令数据码不一样
	CRC    []byte // CRC-16/MODBUS 校验码, 2字节, 是对地址码+功能码+数据码的计算得到
}

// NewUnisoundLamp 创建云知声灯控协议实例
func NewUnisoundLamp() *UnisoundLamp {
	return &UnisoundLamp{}
}

// randomGenerate 随机码生成
func (u *UnisoundLamp) randomGenerate() []byte {
	return []byte{0xff, 0x00}
}

// LampOnCmd 开灯命令
func (u *UnisoundLamp) LampOnCmd() []byte {
	var ulc unisoundLampCmd
	var cmd []byte
	// 举例下发给设备的开灯命令：ff 00 01 06 DF 0B 01 A0 C2 34
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

	// 16 进制打印
	fmt.Printf("LampOnCmd: %X\r\n", cmd)

	return cmd
}

// LampOffCmd 关灯命令
func (u *UnisoundLamp) LampOffCmd() []byte {
	var ulc unisoundLampCmd
	var cmd []byte
	// 举例下发给设备的关灯命令：ff 00 01 06 DF 0B 01 5F 82 74
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

	// 16 进制打印
	fmt.Printf("LampOffCmd: %X\r\n", cmd)

	return cmd
}

// LampBrightnessCmd 调节亮度命令
func (u *UnisoundLamp) LampBrightnessCmd(brightness int) []byte {

	var ulc unisoundLampCmd
	var cmd []byte
	// 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0C 00 00 72 1D -> 亮度0%
	// 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0C 00 14 72 12 -> 亮度20%
	// 举例下发给设备的调节亮度命令：ff 00 01 06 DF 0C 00 64 73 F6 -> 亮度100%
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

	// 16 进制打印
	fmt.Printf("LampBrightnessCmd: %X\r\n", cmd)

	return cmd

}
