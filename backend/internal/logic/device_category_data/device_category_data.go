package deviceCategoryData

import (
	"backend/api/v1/device"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
)

type sDeviceCategoryData struct {
}

func init() {
	service.RegisterDeviceCategoryData(New())
}
func New() service.IDeviceCategoryData {
	return &sDeviceCategoryData{}
}

func (s *sDeviceCategoryData) Add(ctx context.Context, req *device.CategoryDataAddReq) (err error) {
	_, err = dao.SysDeviceCategoryData.Ctx(ctx).Insert(entity.SysDeviceCategoryData{
		CategoryId: req.CategoryId,
		DeviceId:   req.DeviceId,
		DataInt:    int(req.DataInt),
		DataStr:    req.DataStr,
		DataDouble: req.DataDouble,
		CreatedAt:  req.CreatedAt,
	})
	liberr.ErrIsNil(ctx, err, "保存设备数据失败")
	return
}
