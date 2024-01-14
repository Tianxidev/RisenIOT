package v1

import "github.com/gogf/gf/v2/frame/g"

// VersionReq 获取系统版本 请求参数
type VersionReq struct {
	g.Meta `path:"/system/version" tags:"系统" method:"get" summary:"获取系统版本"`
}

// VersionRes 获取系统版本 响应参数
type VersionRes struct {
	g.Meta  `mime:"application/json"`
	Version string `json:"version"`
}
