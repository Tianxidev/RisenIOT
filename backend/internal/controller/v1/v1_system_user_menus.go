package v1

import (
	"backend/internal/model"
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserMenus(ctx context.Context, req *system.UserMenusReq) (res *system.UserMenusRes, err error) {
	var (
		permissions []string
		menuList    []*model.UserMenus
	)
	userId := service.UserCtx().GetUserId(ctx)
	menuList, permissions, err = service.SysUser().GetAdminRules(ctx, userId)
	res = &system.UserMenusRes{
		MenuList:    menuList,
		Permissions: permissions,
	}
	return
}
