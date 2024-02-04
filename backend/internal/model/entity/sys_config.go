// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	ConfigId    uint        `json:"configId"    ` // 参数主键
	ConfigName  string      `json:"configName"  ` // 参数名称
	ConfigKey   string      `json:"configKey"   ` // 参数键名
	ConfigValue string      `json:"configValue" ` // 参数键值
	ConfigType  int         `json:"configType"  ` // 系统内置（Y是 N否）
	CreateBy    uint        `json:"createBy"    ` // 创建者
	UpdateBy    uint        `json:"updateBy"    ` // 更新者
	Remark      string      `json:"remark"      ` // 备注
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 修改时间
}
