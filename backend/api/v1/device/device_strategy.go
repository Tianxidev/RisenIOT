package device

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// StrategySearchReq 查询策略列表
type StrategySearchReq struct {
	g.Meta `path:"/device/strategy" tags:"设备策略" method:"get" summary:"设备策略列表"`
}

// StrategySearchRes 查询策略列表响应
type StrategySearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.SysDeviceStrategy `json:"list"`
}

// StrategyAddReq 添加设备策略
type StrategyAddReq struct {
	g.Meta   `path:"/device/strategy" tags:"设备策略" method:"post" summary:"添加设备策略"`
	Name     string `p:"name" v:"required#策略名称不能为空"`
	Type     int    `p:"type" v:"required#策略类型不能为空"`
	Cron     string `p:"cron"`
	DeviceId int    `p:"deviceId" v:"required#设备ID不能为空"`
	Action   string `p:"action" v:"required#动作不能为空"`
}

// StrategyAddRes 添加设备策略响应
type StrategyAddRes struct {
}
