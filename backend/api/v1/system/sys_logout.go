package system

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta   `path:"/system/login" tags:"登录" method:"post" summary:"用户登录"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type UserLoginRes struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	MenuList    []*model.UserMenus  `json:"menuList"`
	Permissions []string            `json:"permissions"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/system/logout" tags:"登录" method:"get" summary:"退出登录"`
	commonApi.Author
}

type UserLoginOutRes struct {
}
