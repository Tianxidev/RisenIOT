package model

// SysDeviceStrategy 设备策略结构体
type SysDeviceStrategy struct {
	Id       int    `json:"id"`       // 策略ID
	Name     string `json:"name"`     // 策略名称
	Type     string `json:"type"`     // 策略类型
	TypeId   int    `json:"typeId"`   // 策略类型ID
	Cron     string `json:"cron"`     // cron表达式
	DeviceId int    `json:"deviceId"` // 设备ID
}
