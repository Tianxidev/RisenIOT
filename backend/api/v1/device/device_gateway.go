package device

import "github.com/gogf/gf/v2/frame/g"

// GatewayListReq 网关列表请求参数
type GatewayListReq struct {
	g.Meta `path:"/device/gateway/list" tags:"数据网关" method:"get" summary:"网关列表"`
}

// GatewayListRes 网关列表响应参数
type GatewayListRes struct {
}

// GatewayAddReq 添加网关请求参数
type GatewayAddReq struct {
	g.Meta `path:"/device/gateway/add" tags:"数据网关" method:"get" summary:"添加网关"`
}

// GatewayAddRes 添加网关响应参数
type GatewayAddRes struct {
}

// GatewayEditReq 修改网关请求参数
type GatewayEditReq struct {
	g.Meta `path:"/device/gateway/edit" tags:"数据网关" method:"get" summary:"编辑网关"`
}

// GatewayEditRes 修改网关响应参数
type GatewayEditRes struct {
}

// GatewayDelReq 删除网关请求参数
type GatewayDelReq struct {
	g.Meta `path:"/device/gateway/del" tags:"数据网关" method:"get" summary:"删除网关"`
}

// GatewayDelRes 删除网关响应参数
type GatewayDelRes struct {
}
