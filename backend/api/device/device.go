// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package device

import (
	"context"

	"backend/api/device/v1"
)

type IDeviceV1 interface {
	DeviceDataReceive(ctx context.Context, req *v1.DeviceDataReceiveReq) (res *v1.DeviceDataReceiveRes, err error)
	DeviceStatus(ctx context.Context, req *v1.DeviceStatusReq) (res *v1.DeviceStatusRes, err error)
}
