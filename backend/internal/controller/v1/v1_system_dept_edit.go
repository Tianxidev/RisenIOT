package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DeptEdit(ctx context.Context, req *system.DeptEditReq) (res *system.DeptEditRes, err error) {
	err = service.SysDept().Edit(ctx, req)
	return
}
