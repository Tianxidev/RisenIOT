package device

import "github.com/gogf/gf/v2/frame/g"

type GroupListReq struct {
	g.Meta `path:"/device/group/list" tags:"设备分组" method:"get" summary:"获取设备分组列表"`
}

type GroupListRes struct {
	g.Meta `mime:"application/json"`
	List   []g.Map
}

type GroupAddReq struct {
	g.Meta `path:"/device/group/add" tags:"设备分组" method:"post" summary:"添加设备分组"`
	Name   string `p:"name" v:"required|length:1,30#请输入分组名称|分组名称长度为:min到:max位"`
	Remark string `p:"remark" v:"length:0,100#备注长度为:min到:max位"`
}

type GroupAddRes struct {
}
