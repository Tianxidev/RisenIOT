package device

import "github.com/gogf/gf/v2/frame/g"

type GatewayListReq struct {
	g.Meta `path:"/device/gateway/list" tags:"数据网关" method:"get" summary:"网关列表"`
}

type GatewayListRes struct {
}

type GatewayAddReq struct {
	g.Meta `path:"/device/gateway/add" tags:"数据网关" method:"get" summary:"添加网关"`
}

type GatewayAddRes struct {
}

type GatewayEditReq struct {
	g.Meta `path:"/device/gateway/edit" tags:"数据网关" method:"get" summary:"编辑网关"`
}

type GatewayEditRes struct {
}

type GatewayDelReq struct {
	g.Meta `path:"/device/gateway/del" tags:"数据网关" method:"get" summary:"删除网关"`
}

// GatewayDelRes 删除网关响应参数
type GatewayDelRes struct {
}
