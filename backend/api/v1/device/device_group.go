package device

import "github.com/gogf/gf/v2/frame/g"

// GroupListReq 查询设备分组列表
type GroupListReq struct {
	g.Meta `path:"/device/group/list" tags:"设备分组" method:"get" summary:"获取设备分组列表"`
}

// GroupListRes 查询设备分组列表响应
type GroupListRes struct {
	g.Meta `mime:"application/json"`
	List   []g.Map `json:"list"`
	Total  int     `json:"total"`
}

// GroupAddReq 添加设备分组
type GroupAddReq struct {
	g.Meta  `path:"/device/group/add" tags:"设备分组" method:"post" summary:"添加设备分组"`
	Name    string `p:"name" v:"required|length:1,30#请输入分组名称|分组名称长度为:min到:max位"`
	Remarks string `p:"remarks" v:"length:0,100#备注长度为:min到:max位"`
}

// GroupAddRes 添加设备分组响应
type GroupAddRes struct {
}

// GroupEditReq 编辑设备分组
type GroupEditReq struct {
	g.Meta  `path:"/device/group/edit" tags:"设备分组" method:"put" summary:"编辑设备分组"`
	Id      int    `p:"id" v:"required#主键ID不能为空"`
	Name    string `p:"name" v:"required|length:1,30#请输入分组名称|分组名称长度为:min到:max位"`
	Remarks string `p:"remarks" v:"length:0,100#备注长度为:min到:max位"`
}

// GroupEditRes 编辑设备分组响应
type GroupEditRes struct {
}

// GroupDelReq 删除设备分组
type GroupDelReq struct {
	g.Meta `path:"/device/group/delete" tags:"设备分组" method:"delete" summary:"删除设备分组"`
	Ids    []int `p:"ids" v:"required#主键ID不能为空"`
}

// GroupDelRes 删除设备分组响应
type GroupDelRes struct {
}
