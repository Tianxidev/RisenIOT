package v1

import "github.com/gogf/gf/v2/frame/g"

// DeviceStatusReq 设备状态API 请求参数
type DeviceStatusReq struct {
	g.Meta   `path:"/device/status" tags:"设备数据API" method:"post" summary:"设备API"`
	DeviceId int `p:"deviceId" summary:"设备ID"`
}

// DeviceStatusRes 设备状态API 响应参数
type DeviceStatusRes struct {
	g.Meta `mime:"application/json"`
}
