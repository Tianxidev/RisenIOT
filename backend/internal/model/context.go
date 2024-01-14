package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Context 请求上下文结构
type Context struct {
	Code    int            // 响应状态码
	Message string         // 消息内容
	Session *ghttp.Session // 当前Session管理对象
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// DefaultHandlerResponse 统一返回对象
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"msg" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// UserInfo 用户信息
type UserInfo struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
}
