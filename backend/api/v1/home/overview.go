package home

import "github.com/gogf/gf/v2/frame/g"

// OverviewReq 获取首页概览请求
type OverviewReq struct {
	g.Meta `path:"/home/overview" tags:"概览" method:"get" summary:"获取概览数据"`
}

// OverviewRes 获取首页概览响应
type OverviewRes struct {
	TotalDevice int `json:"totalDevice"` // 总设备数
	Online      int `json:"online"`      // 在线设备数
	Offline     int `json:"offline"`     // 离线设备数
}
