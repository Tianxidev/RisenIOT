package deviceKind

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

type sDeviceKind struct {
}

func init() {
	service.RegisterDeviceKind(New())
}

func New() service.IDeviceKind {
	return &sDeviceKind{}
}

// Get 通过id获取
func (s *sDeviceKind) Get(ctx context.Context, id int) (info *entity.SysDeviceKind, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.SysDeviceKind.Ctx(ctx).Where(dao.SysDeviceKind.Columns().Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(ctx, "get device kind err", err, info)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// List 获取设备种类列表
func (s *sDeviceKind) List(ctx context.Context, req *device.KindSearchReq) (total, page int, list []*entity.SysDeviceKind, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysDeviceKind.Ctx(ctx)
	if req.Name != "" {
		m = m.Where(dao.SysDeviceKind.Columns().Name+" like ?", "%"+req.Name+"%")
	}
	if req.Mark != "" {
		m = m.Where(dao.SysDeviceKind.Columns().Mark+" = ?", req.Mark)
	}
	if req.TimeOut != "" {
		m = m.Where(dao.SysDeviceKind.Columns().TimeOut+" = ?", req.TimeOut)
	}
	if req.BeginTime != "" {
		m = m.Where(dao.SysDeviceKind.Columns().CreatedAt+" >= ", req.BeginTime)
	}
	if req.EndTime != "" {
		m = m.Where(dao.SysDeviceKind.Columns().CreatedAt+" <= ", req.EndTime)
	}
	if req.Id != 0 {
		m = m.Where(dao.SysDeviceKind.Columns().Id+" = ?", req.Id)
	}
	g.Log().Printf(ctx, "req param:%v \n", req)
	err = g.Try(ctx, func(ctx context.Context) {
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取DeviceKind列表失败")
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
