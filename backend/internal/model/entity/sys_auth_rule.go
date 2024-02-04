// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAuthRule is the golang structure for table sys_auth_rule.
type SysAuthRule struct {
	Id         uint        `json:"id"         ` //
	Pid        uint        `json:"pid"        ` // 父ID
	Name       string      `json:"name"       ` // 规则名称
	Title      string      `json:"title"      ` // 规则名称
	Icon       string      `json:"icon"       ` // 图标
	Condition  string      `json:"condition"  ` // 条件
	Remark     string      `json:"remark"     ` // 备注
	MenuType   uint        `json:"menuType"   ` // 类型 0目录 1菜单 2按钮
	Weigh      int         `json:"weigh"      ` // 权重
	IsHide     uint        `json:"isHide"     ` // 显示状态
	Path       string      `json:"path"       ` // 路由地址
	Component  string      `json:"component"  ` // 组件路径
	IsLink     uint        `json:"isLink"     ` // 是否外链 1是 0否
	ModuleType string      `json:"moduleType" ` // 所属模块
	ModelId    uint        `json:"modelId"    ` // 模型ID
	IsIframe   uint        `json:"isIframe"   ` // 是否内嵌iframe
	IsCached   uint        `json:"isCached"   ` // 是否缓存
	Redirect   string      `json:"redirect"   ` // 路由重定向地址
	IsAffix    uint        `json:"isAffix"    ` // 是否固定
	LinkUrl    string      `json:"linkUrl"    ` // 链接地址
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建日期
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 修改日期
}
