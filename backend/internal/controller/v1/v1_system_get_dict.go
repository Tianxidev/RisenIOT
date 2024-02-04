package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) GetDict(ctx context.Context, req *system.GetDictReq) (res *system.GetDictRes, err error) {
	res, err = service.SysDictData().GetDictWithDataByType(ctx, req)
	return
}
