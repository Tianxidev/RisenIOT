package v1

import (
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PersonalInfo(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error) {
	res = new(system.PersonalInfoRes)
	userId := service.UserCtx().GetUserId(ctx)
	res.User, err = service.SysUser().GetUserInfoById(ctx, userId)
	var dept *entity.SysDept
	dept, err = service.SysDept().GetByDeptId(ctx, res.User.DeptId)
	res.DeptName = dept.DeptName
	allRoles, err := service.SysRole().GetRoleList(ctx)
	roles, err := service.SysUser().GetAdminRole(ctx, userId, allRoles)
	name := make([]string, len(roles))
	roleIds := make([]uint, len(roles))
	for k, v := range roles {
		name[k] = v.Name
		roleIds[k] = v.Id
	}
	res.Roles = name
	if err != nil {
		return
	}
	return
}
