package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DeptAdd(ctx context.Context, req *system.DeptAddReq) (res *system.DeptAddRes, err error) {
	err = service.SysDept().Add(ctx, req)
	return
}
