package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RuleDelete(ctx context.Context, req *system.RuleDeleteReq) (res *system.RuleDeleteRes, err error) {
	err = service.SysAuthRule().DeleteMenuByIds(ctx, req.Ids)
	return
}
