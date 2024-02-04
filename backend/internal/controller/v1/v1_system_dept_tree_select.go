package v1

import (
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) DeptTreeSelect(ctx context.Context, req *system.DeptTreeSelectReq) (res *system.DeptTreeSelectRes, err error) {
	var deptList []*entity.SysDept
	deptList, err = service.SysDept().GetList(ctx, &system.DeptSearchReq{
		Status: "1", //正常状态数据
	})
	if err != nil {
		return
	}
	res = new(system.DeptTreeSelectRes)
	res.Deps = service.SysDept().GetListTree(0, deptList)
	return
}
