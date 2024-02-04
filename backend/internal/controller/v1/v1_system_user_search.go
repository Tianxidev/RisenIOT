package v1

import (
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserSearch(ctx context.Context, req *system.UserSearchReq) (res *system.UserSearchRes, err error) {
	var (
		total    interface{}
		userList []*entity.SysUser
	)
	res = new(system.UserSearchRes)
	total, userList, err = service.SysUser().List(ctx, req)
	if err != nil || total == 0 {
		return
	}
	res.Total = total
	res.UserList, err = service.SysUser().GetUsersRoleDept(ctx, userList)
	return
}
