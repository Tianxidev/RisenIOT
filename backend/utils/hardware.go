package utils

import (
	"encoding/json"
)

// CpuUsage CPU使用率
type CpuUsage struct {
	Name  string `json:"name"`
	Usage uint64 `json:"usage"`
}

// String CpuUsage 转换为 json 字符串
func (m CpuUsage) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

// DiskUsage 磁盘使用率
type DiskUsage struct {
	DeviceID  string  `json:"deviceID"`
	FreeSpace float64 `json:"freeSpace"`
	Size      float64 `json:"size"`
}

// String DiskUsage 转换为 json 字符串
func (m DiskUsage) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

// NetworkInterfaceUsage 网络接口使用率
type NetworkInterfaceUsage struct {
	Name                string
	CurrentBandwidth    uint64
	BytesTotalPerSec    uint64
	BytesReceivedPerSec uint64
	BytesSentPerSec     uint64
	PacketsPerSec       uint64
}

// String NetworkInterfaceUsage 转换为 json 字符串
func (m NetworkInterfaceUsage) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

// GetCpuUsage 获取系统 CPU 使用率
func GetCpuUsage() ([]CpuUsage, error) {

	return []CpuUsage{}, nil
}

// GetDiskUsage 获取系统磁盘使用率
func GetDiskUsage() ([]DiskUsage, error) {
	return []DiskUsage{}, nil
}
