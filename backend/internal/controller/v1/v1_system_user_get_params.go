package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) UserGetParams(ctx context.Context, req *system.UserGetParamsReq) (res *system.UserGetParamsRes, err error) {
	res = new(system.UserGetParamsRes)
	res.RoleList, err = service.SysRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	res.Posts, err = service.SysPost().GetUsedPost(ctx)
	return
}
