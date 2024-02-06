package model

import "github.com/gogf/gf/v2/os/gtime"

// DeviceEvent 设备事件
type DeviceEvent struct {
	Name string      // 事件名称
	Data interface{} // 事件数据
}

// DeviceData 设备数据
type DeviceData struct {
	CategoryId int         // 设备分类ID
	Name       string      // 设备名称
	Type       int         // 设备类型
	Data       interface{} // 设备数据
	Time       *gtime.Time // 时间
	Ratio      string      // 比率
}

// DeviceDecodeMsg 设备解码消息
type DeviceDecodeMsg struct {
	DataList   []*DeviceData  // 设备数据列表
	EventList  []*DeviceEvent // 设备事件列表
	DeviceInfo *DeviceAllInfo // 设备信息
}
