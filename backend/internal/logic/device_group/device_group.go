package deviceGroup

import (
	"backend/api/v1/device"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
	var list []*entity.SysDeviceGroup
	res = &device.GroupListRes{}

	// 查询设备分组列表
	err = m.Scan(&list)
	liberr.ErrIsNil(ctx, err, "查询设备分组列表失败")

	// 遍历设备分组列表
	for _, v := range list {
		res.List = append(res.List, g.Map{
			"id":      v.Id,
			"name":    v.Name,
			"remarks": v.Remarks,
		})
	}

	return
}
