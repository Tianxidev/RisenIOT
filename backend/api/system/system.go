// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"backend/api/system/v1"
)

type ISystemV1 interface {
	Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (res *v1.HeartbeatRes, err error)
	Pk(ctx context.Context, req *v1.PkReq) (res *v1.PkRes, err error)
	Version(ctx context.Context, req *v1.VersionReq) (res *v1.VersionRes, err error)
}
