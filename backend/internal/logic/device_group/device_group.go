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

func New() *sDeviceGroup {
	return &sDeviceGroup{}
}

// List 查询设备分组列表
func (s *sDeviceGroup) List(ctx context.Context) (res *device.GroupListRes, err error) {
	t1 := dao.SysDeviceGroup
	m := t1.Ctx(ctx)
	var list []*entity.SysDeviceGroup
	res = &device.GroupListRes{}
	userId := int(service.UserCtx().GetUserId(ctx))

	// 查询设备分组列表
	err = m.Scan(&list)
	liberr.ErrIsNil(ctx, err, "查询设备分组列表失败")

	// 遍历设备分组列表
	for _, v := range list {

		// 跳过不是当前用户创建的设备分组
		if v.CreateBy != userId && v.Id != -1 {
			continue
		}

		// 添加设备分组到返回列表
		res.List = append(res.List, g.Map{
			"id":      v.Id,
			"name":    v.Name,
			"remarks": v.Remarks,
		})
	}

	res.Total = len(res.List)

	return
}

// Add 添加设备分组
func (s *sDeviceGroup) Add(ctx context.Context, req *device.GroupAddReq) (err error) {
	t1 := dao.SysDeviceGroup
	m := t1.Ctx(ctx)

	// 添加设备分组
	_, err = m.Data(g.Map{
		t1.Columns().Name:     req.Name,
		t1.Columns().Remarks:  req.Remarks,
		t1.Columns().CreateBy: int(service.UserCtx().GetUserId(ctx)),
	}).Insert()
	liberr.ErrIsNil(ctx, err, "添加设备分组失败")

	return

}

// Edit 编辑
func (s *sDeviceGroup) Edit(ctx context.Context, req *device.GroupEditReq) (err error) {
	t1 := dao.SysDeviceGroup
	userID := int(service.UserCtx().GetUserId(ctx))
	m := t1.Ctx(ctx)

	// 编辑设备分组
	m = m.Data(g.Map{
		t1.Columns().Name:    req.Name,
		t1.Columns().Remarks: req.Remarks,
	})
	m = m.Where(t1.Columns().Id, req.Id)

	// 判断是否是超级管理员
	if !service.SysUser().IsSuperAdmin(ctx, userID) {
		m = m.Where(t1.Columns().CreateBy, userID)
	}

	_, err = m.Update()
	liberr.ErrIsNil(ctx, err, "编辑设备分组失败")

	return
}

// Del 删除设备分组
func (s *sDeviceGroup) Del(ctx context.Context, ids []int) (err error) {
	var count int

	t1 := dao.SysDeviceGroup
	t2 := dao.SysDeviceInfo
	m := t1.Ctx(ctx)

	// 遍历设备分组列表
	for _, v := range ids {

		// 跳过默认分组
		if v == -1 {
			continue
		}

		// 查询分组下是否有设备
		count, err = t2.Ctx(ctx).Where(t2.Columns().Group, v).Count()
		liberr.ErrIsNil(ctx, err, "查询设备分组失败")
		liberr.ValueIsTrue(ctx, count > 0, "设备分组下存在设备, 请先删除设备")

		// 删除设备分组
		m.Where(t1.Columns().CreateBy, int(service.UserCtx().GetUserId(ctx)))
		m = m.Where(t1.Columns().Id, v)
		_, err = m.Delete()
		liberr.ErrIsNil(ctx, err, "删除设备分组失败")

	}

	return
}
