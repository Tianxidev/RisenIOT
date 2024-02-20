package device_categoty

import (
	"backend/api/v1/device"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sDeviceCategory struct {
}

func init() {
	service.RegisterDeviceCategory(New())
}

func New() *sDeviceCategory {
	return &sDeviceCategory{}
}

// List 列表
func (s *sDeviceCategory) List(ctx context.Context, req *device.CategorySearchReq) (total, page int, list []*entity.SysDeviceCategory, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysDeviceCategory.Ctx(ctx)

	// 判断产品类型是否不为空
	liberr.ValueIsTrue(ctx, req.KindId == 0, "产品类型不能为空")

	// 根据产品类型获取产品数据类型
	m = m.Where(dao.SysDeviceCategory.Columns().KindId+" = ?", req.KindId)

	err = g.Try(ctx, func(ctx context.Context) {
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取设备数据列表失败")
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取总行数失败")
			return
		}
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		err = m.Page(page, req.PageSize).Order(order).Scan(&list)
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取数据失败")
		}
	})
	return
}

// KindGet 根据产品类型获取产品数据类型
func (s *sDeviceCategory) KindGet(ctx context.Context, kindId int) (list []*entity.SysDeviceCategory, err error) {
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
