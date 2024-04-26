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

func New() *sDeviceKind {
	return &sDeviceKind{}
}

// Get 查询产品详情
func (s *sDeviceKind) Get(ctx context.Context, id int) (info *entity.SysDeviceKind, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	one, err := g.Model(dao.SysDeviceKind.Table()).Where("id = ?", id).One()
	liberr.ErrIsNil(ctx, err, "查询产品详情失败")
	err = one.Struct(&info)
	liberr.ErrIsNil(ctx, err, "产品详情映射失败")
	return
}

// List 查询产品类型列表
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

	err = g.Try(ctx, func(ctx context.Context) {
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取产品类型列表失败")
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取总行数失败")
			return
		}
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}

		// 根据用户id查询创建产品类型以及公共产品类型
		m = m.Where("create_by = ? or public = 1", int(service.UserCtx().GetUserId(ctx)))
		err = m.Page(page, req.PageSize).Order(order).Scan(&list)
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取产品类型数据失败")
		}
	})
	return
}

// Add 添加产品类型
func (s *sDeviceKind) Add(ctx context.Context, req *device.KindAddReq) (err error) {
	liberr.ValueIsTrue(ctx, req.TimeOut <= 0, "超时时间不能小于0")
	t1 := dao.SysDeviceKind

	// 判断产品类型是否存在
	m := dao.SysDeviceKind.Ctx(ctx)
	m = m.Where(t1.Columns().Name, req.Name)
	m = m.Where(t1.Columns().CreateBy, int(service.UserCtx().GetUserId(ctx)))
	kind, err := m.One()
	liberr.ErrIsNil(ctx, err, "查询产品类型失败")
	liberr.ValueIsTrue(ctx, kind != nil, "产品类型已存在")

	// 添加产品类型
	_, err = dao.SysDeviceKind.Ctx(ctx).Data(&entity.SysDeviceKind{
		Name:     req.Name,
		Mark:     req.Mark,
		TimeOut:  req.TimeOut,
		CreateBy: int(service.UserCtx().GetUserId(ctx)),
	}).Insert()
	liberr.ErrIsNil(ctx, err, "添加产品类型失败")
	return
}

// Edit 编辑产品类型
func (s *sDeviceKind) Edit(ctx context.Context, req *device.KindEditReq) (err error) {
	kind := new(entity.SysDeviceKind)
	liberr.ValueIsTrue(ctx, req.TimeOut <= 0, "超时时间不能小于0")
	t1 := dao.SysDeviceKind
	m := dao.SysDeviceKind.Ctx(ctx)
	m = m.Where(t1.Columns().Id, req.Id)
	m = m.Where(t1.Columns().CreateBy, int(service.UserCtx().GetUserId(ctx)))
	one, err := m.One()
	liberr.ErrIsNil(ctx, err, "查询产品类型失败")
	liberr.ValueIsTrue(ctx, one == nil, "产品类型不存在")

	// 映射产品类型
	err = one.Struct(kind)
	liberr.ErrIsNil(ctx, err, "产品类型映射失败")

	// 编辑产品类型
	_, err = dao.SysDeviceKind.Ctx(ctx).Data(&entity.SysDeviceKind{
		Id:       kind.Id,
		Name:     req.Name,
		Mark:     req.Mark,
		TimeOut:  req.TimeOut,
		CreateBy: int(service.UserCtx().GetUserId(ctx)),
	}).Where(t1.Columns().Id, req.Id).Update()
	liberr.ErrIsNil(ctx, err, "编辑产品类型失败")
	return
}

// Del 删除产品类型
func (s *sDeviceKind) Del(ctx context.Context, id []int) (err error) {
	t1 := dao.SysDeviceKind
	t2 := dao.SysDeviceInfo

	// 遍历删除产品类型
	for _, v := range id {

		// 判断产品类型下是否有设备
		count, err := t2.Ctx(ctx).Where(t2.Columns().Kind, v).Count()
		liberr.ErrIsNil(ctx, err, "查询设备信息失败")
		liberr.ValueIsTrue(ctx, count > 0, "产品类型下存在设备, 请先删除设备")

		// 删除产品类型
		m1 := t1.Ctx(ctx)
		m1 = m1.Where(t1.Columns().Id, v)
		m1 = m1.Where(t1.Columns().CreateBy, int(service.UserCtx().GetUserId(ctx)))
		_, err = m1.Delete()
		liberr.ErrIsNil(ctx, err, "删除产品类型失败")

	}

	return
}
