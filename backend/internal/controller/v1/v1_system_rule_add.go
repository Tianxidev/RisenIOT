package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) RuleAdd(ctx context.Context, req *system.RuleAddReq) (res *system.RuleAddRes, err error) {
	err = service.SysAuthRule().Add(ctx, req)
	return
}
