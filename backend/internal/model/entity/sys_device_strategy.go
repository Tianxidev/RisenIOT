// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysDeviceStrategy is the golang structure for table sys_device_strategy.
type SysDeviceStrategy struct {
	Id       int    `json:"id"       orm:"id"        ` // 策略ID
	Name     string `json:"name"     orm:"name"      ` // 策略名称
	Type     int    `json:"type"     orm:"type"      ` // 策略类型
	Cron     string `json:"cron"     orm:"cron"      ` // Cron时间
	DeviceId int    `json:"deviceId" orm:"device_id" ` // 设备ID
	AuthorId int    `json:"authorId" orm:"author_id" ` // 创建用户ID
	Action   string `json:"action"   orm:"action"    ` // 策略动作
}
