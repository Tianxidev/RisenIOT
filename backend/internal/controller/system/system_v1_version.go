package system

import (
	"context"

	"backend/api/system/v1"
)

// Version 获取系统版本
func (c *ControllerV1) Version(ctx context.Context, req *v1.VersionReq) (res *v1.VersionRes, err error) {
	res = &v1.VersionRes{}
	res.Version = "0.0.1"
	return res, nil
}
