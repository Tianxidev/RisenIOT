package v1

import "github.com/gogf/gf/v2/frame/g"

// UserListReq 获取用户列表 请求参数
type UserListReq struct {
	g.Meta `path:"/user/list" tags:"用户" method:"get" summary:"获取用户列表"`
}

// UserListRes 获取用户列表 响应参数
type UserListRes struct {
	g.Meta `mime:"application/json"`
}
