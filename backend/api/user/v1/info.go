package v1

import "github.com/gogf/gf/v2/frame/g"

// UserInfoReq 获取用户信息 请求参数
type UserInfoReq struct {
	g.Meta `path:"/user/info" tags:"用户" method:"get" summary:"获取用户信息"`
}

// UserInfoRes 获取用户信息 响应参数
type UserInfoRes struct {
	g.Meta `mime:"application/json"`
}
