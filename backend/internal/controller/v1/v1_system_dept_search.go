package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DeptSearch(ctx context.Context, req *system.DeptSearchReq) (res *system.DeptSearchRes, err error) {
	res = new(system.DeptSearchRes)
	res.DeptList, err = service.SysDept().GetList(ctx, req)
	return
}
