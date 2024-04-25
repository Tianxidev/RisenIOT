// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceStrategy is the golang structure of table sys_device_strategy for DAO operations like Where/Data.
type SysDeviceStrategy struct {
	g.Meta   `orm:"table:sys_device_strategy, do:true"`
	Id       interface{} // 策略ID
	Name     interface{} // 策略名称
	Type     interface{} // 策略类型
	Cron     interface{} // Cron时间
	DeviceId interface{} // 设备ID
	AuthorId interface{} // 创建用户ID
	Action   interface{} // 策略动作
}
