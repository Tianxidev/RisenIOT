package device

import "github.com/gogf/gf/v2/frame/g"

type DataGetReq struct {
	g.Meta    `path:"/device/data/get" tags:"设备数据API" method:"get" summary:"获取设备数据"`
	DeviceId  int    `p:"deviceId" `
	DeviceSn  string `p:"deviceSn" `
	DevicePwd string `p:"devicePwd" `
}

type DataGetRes struct {
	g.Meta `mime:"application/json"`
	g.MapStrStr
}

// DataAddReq 添加操作请求参数
type DataAddReq struct {
	g.Meta    `path:"/device/data/add" tags:"设备数据API" method:"post" summary:"添加设备数据"`
	DeviceId  int         `p:"deviceId" `
	DeviceSn  string      `p:"deviceSn" `
	DevicePwd string      `p:"devicePwd" `
	Time      string      `p:"time" `
	Event     interface{} `p:"event" `
	Property  interface{} `p:"property" `
}

type DataAddRes struct {
}
