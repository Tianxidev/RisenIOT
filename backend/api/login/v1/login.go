package v1

import (
	"backend/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLoginReq 用户登录请求
type UserLoginReq struct {
	g.Meta   `path:"/login" tags:"登录" method:"post" summary:"用户登录"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

// UserLoginRes 用户登录响应
type UserLoginRes struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	Permissions []string            `json:"permissions"`
}
