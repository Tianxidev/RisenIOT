package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PostEdit(ctx context.Context, req *system.PostEditReq) (res *system.PostEditRes, err error) {
	err = service.SysPost().Edit(ctx, req)
	return
}
