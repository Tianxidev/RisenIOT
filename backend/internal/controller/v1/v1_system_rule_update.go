package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RuleUpdate(ctx context.Context, req *system.RuleUpdateReq) (res *system.RuleUpdateRes, err error) {
	err = service.SysAuthRule().Update(ctx, req)
	return
}
