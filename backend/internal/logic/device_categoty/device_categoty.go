package device_categoty

import (
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sDeviceCategory struct {
}

func init() {
	service.RegisterDeviceCategory(New())
}

func New() service.IDeviceCategory {
	return &sDeviceCategory{}
}

func (s *sDeviceCategory) KindGet(ctx context.Context, kindId int) (list []*entity.SysDeviceCategoty, err error) {
	m := dao.SysDeviceCategory.Ctx(ctx)
	if kindId != 0 {
		m = m.Where(dao.SysDeviceCategory.Columns().KindId+" = ?", kindId)
	}

	err = g.Try(ctx, func(ctx context.Context) {
		order := "id asc"
		err = m.Order(order).Scan(&list)
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取数据失败")
		}
	})
	return
}
