package v1

import "github.com/gogf/gf/v2/frame/g"

// DeviceDataReceiveReq 设备数据接收API 请求参数
type DeviceDataReceiveReq struct {
	g.Meta    `path:"/device/receive" tags:"设备数据API" method:"post" summary:"设备API"`
	DeviceId  int         `p:"deviceId" `
	DeviceSn  string      `p:"deviceSn" `
	DevicePwd string      `p:"devicePwd" `
	Time      string      `p:"time" `
	Event     interface{} `p:"event" `
	Property  interface{} `p:"property" `
}

// DeviceDataReceiveRes 设备数据接收API 响应参数
type DeviceDataReceiveRes struct {
	g.Meta `mime:"application/json"`
}
