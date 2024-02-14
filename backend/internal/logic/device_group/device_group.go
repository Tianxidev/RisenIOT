package deviceGroup

import (
	"backend/api/v1/device"
	"backend/internal/dao"
	"backend/internal/service"
	"context"
)

type sDeviceGroup struct {
}

func init() {
	service.RegisterDeviceGroup(New())
}

func New() service.IDeviceGroup {
	return &sDeviceGroup{}
}

// GroupList 查询设备分组列表
func (s *sDeviceGroup) GroupList(ctx context.Context) (res *device.GroupListRes, err error) {
	m := dao.SysDeviceGroup.Ctx(ctx)
	err = m.Scan(res.List)
	return
}
