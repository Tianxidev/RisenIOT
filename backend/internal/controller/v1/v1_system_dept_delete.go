package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DeptDelete(ctx context.Context, req *system.DeptDeleteReq) (res *system.DeptDeleteRes, err error) {
	err = service.SysDept().Delete(ctx, req.Id)
	return
}
