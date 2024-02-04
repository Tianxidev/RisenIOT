package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) SysUserOnlineSearch(ctx context.Context, req *system.SysUserOnlineSearchReq) (res *system.SysUserOnlineSearchRes, err error) {
	res, err = service.SysUserOnline().GetOnlineListPage(ctx, req)
	return
}
