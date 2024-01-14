package v1

import "github.com/gogf/gf/v2/frame/g"

// PkReq 获取系统公钥 请求参数
type PkReq struct {
	g.Meta `path:"/system/pk" tags:"系统" method:"get" summary:"获取系统公钥"`
}

// PkRes 获取系统公钥 响应参数
type PkRes struct {
	g.Meta `mime:"application/json"`
	Pk     string `json:"pk"`
}
