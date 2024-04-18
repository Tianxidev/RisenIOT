// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysDeviceGroup is the golang structure for table sys_device_group.
type SysDeviceGroup struct {
	Id       int    `json:"id"       orm:"id"        ` // 分组ID
	Name     string `json:"name"     orm:"name"      ` // 分组名称
	Remarks  string `json:"remarks"  orm:"remarks"   ` // 备注
	CreateBy int    `json:"createBy" orm:"create_by" ` // 创建人ID
}
