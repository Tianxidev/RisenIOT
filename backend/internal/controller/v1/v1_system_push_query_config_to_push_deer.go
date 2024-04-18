package v1

import (
	"backend/internal/service"
	"context"

	"backend/api/v1/system"
)

func (c *ControllerSystem) PushQueryConfigToPushDeer(ctx context.Context, req *system.PushQueryConfigToPushDeerReq) (res *system.PushQueryConfigToPushDeerRes, err error) {
	res = new(system.PushQueryConfigToPushDeerRes)
	service.SysPush().QueryPushDeerConfig(ctx, res)
	return
}
