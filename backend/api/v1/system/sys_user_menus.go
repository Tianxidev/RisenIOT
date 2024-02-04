package system

import (
	commonApi "backend/api/v1/common"
	"backend/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type UserMenusReq struct {
	g.Meta `path:"/system/user/getUserMenus" tags:"用户管理" method:"get" summary:"获取用户菜单"`
	commonApi.Author
}

type UserMenusRes struct {
	g.Meta      `mime:"application/json"`
	MenuList    []*model.UserMenus `json:"menuList"`
	Permissions []string           `json:"permissions"`
}
