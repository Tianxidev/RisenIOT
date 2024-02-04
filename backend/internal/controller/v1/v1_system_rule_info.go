package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RuleInfo(ctx context.Context, req *system.RuleInfoReq) (res *system.RuleInfoRes, err error) {
	res = new(system.RuleInfoRes)
	res.Rule, err = service.SysAuthRule().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.RoleIds, err = service.SysAuthRule().GetMenuRoles(ctx, req.Id)
	return
}
