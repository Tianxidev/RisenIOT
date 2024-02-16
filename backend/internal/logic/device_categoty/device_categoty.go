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
	if req.KindId != 0 {
		m = m.Where(dao.SysDeviceCategory.Columns().KindId+" = ?", req.KindId)
	}
	if req.Name != "" {
		m = m.Where(dao.SysDeviceCategory.Columns().Name+" like ?", "%"+req.Name+"%")
	}
	if req.Mark != "" {
		m = m.Where(dao.SysDeviceCategory.Columns().Mark+" = ?", req.Mark)
	}
	if req.DataType != "" {
		m = m.Where(dao.SysDeviceCategory.Columns().DataType+" = ?", req.DataType)
	}
	if req.Unit != "" {
		m = m.Where(dao.SysDeviceCategory.Columns().Unit+" = ?", req.Unit)
	}
	if req.Ratio != "" {
		m = m.Where(dao.SysDeviceCategory.Columns().Ratio+" = ?", req.Ratio)
	}
	if req.Format != "" {
		m = m.Where(dao.SysDeviceCategory.Columns().Format+" = ?", req.Format)
	}
	if req.HomeShow != "" {
		m = m.Where(dao.SysDeviceCategory.Columns().HomeShow+" = ?", req.HomeShow)
	}
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
