// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAuthRule is the golang structure for table sys_auth_rule.
type SysAuthRule struct {
	Id         uint        `json:"id"         orm:"id"          ` //
	Pid        uint        `json:"pid"        orm:"pid"         ` // 父ID
	Name       string      `json:"name"       orm:"name"        ` // 规则名称
	Title      string      `json:"title"      orm:"title"       ` // 规则名称
	Icon       string      `json:"icon"       orm:"icon"        ` // 图标
	Condition  string      `json:"condition"  orm:"condition"   ` // 条件
	Remark     string      `json:"remark"     orm:"remark"      ` // 备注
	MenuType   uint        `json:"menuType"   orm:"menu_type"   ` // 类型 0目录 1菜单 2按钮
	Weigh      int         `json:"weigh"      orm:"weigh"       ` // 权重
	IsHide     uint        `json:"isHide"     orm:"is_hide"     ` // 显示状态
	Path       string      `json:"path"       orm:"path"        ` // 路由地址
	Component  string      `json:"component"  orm:"component"   ` // 组件路径
	IsLink     uint        `json:"isLink"     orm:"is_link"     ` // 是否外链 1是 0否
	ModuleType string      `json:"moduleType" orm:"module_type" ` // 所属模块
	ModelId    uint        `json:"modelId"    orm:"model_id"    ` // 模型ID
	IsIframe   uint        `json:"isIframe"   orm:"is_iframe"   ` // 是否内嵌iframe
	IsCached   uint        `json:"isCached"   orm:"is_cached"   ` // 是否缓存
	Redirect   string      `json:"redirect"   orm:"redirect"    ` // 路由重定向地址
	IsAffix    uint        `json:"isAffix"    orm:"is_affix"    ` // 是否固定
	LinkUrl    string      `json:"linkUrl"    orm:"link_url"    ` // 链接地址
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // 创建日期
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  ` // 修改日期
}
