package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RuleGetParams(ctx context.Context, req *system.RuleGetParamsReq) (res *system.RuleGetParamsRes, err error) {
	res = new(system.RuleGetParamsRes)
	res.Roles, err = service.SysRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	res.Menus, err = service.SysAuthRule().GetIsMenuList(ctx)
	return
}
