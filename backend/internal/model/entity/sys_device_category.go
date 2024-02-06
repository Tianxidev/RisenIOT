// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysDeviceCategory is the golang structure for table sys_device_category.
type SysDeviceCategory struct {
	Id        int    `json:"id"        ` //
	KindId    int    `json:"kindId"    ` // 类型ID
	Name      string `json:"name"      ` //
	Mark      string `json:"mark"      ` //
	DataType  int    `json:"dataType"  ` // 数据类型
	LogicType int    `json:"logicType" ` //
	Unit      string `json:"unit"      ` // 单位
	Ratio     string `json:"ratio"     ` // 比例
	Format    string `json:"format"    ` //
	HomeShow  int    `json:"homeShow"  ` //
	Remark    string `json:"remark"    ` //
	CreatedAt string `json:"createdAt" ` //
	UpdatedAt string `json:"updatedAt" ` //
}
