package deviceStatus

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
	"github.com/gogf/gf/v2/os/gtime"
)

type sDeviceStatus struct {
}

func init() {
	service.RegisterDeviceStatus(New())
}

func New() *sDeviceStatus {
	return &sDeviceStatus{}
}

// List 获取状态列表
func (s *sDeviceStatus) List(ctx context.Context, req *device.StatusSearchReq) (total, page int, list []*entity.SysDeviceStatus, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysDeviceStatus.Ctx(ctx)
	if req.DeviceId != "" {
		m = m.Where(dao.SysDeviceStatus.Columns().DeviceId+" = ?", req.DeviceId)
	}
	if req.Status != "" {
		m = m.Where(dao.SysDeviceStatus.Columns().Status+" = ?", req.Status)
	}
	if req.TimeOut != "" {
		m = m.Where(dao.SysDeviceStatus.Columns().TimeOut+" = ?", req.TimeOut)
	}
	if req.UpTime != "" {
		m = m.Where(dao.SysDeviceStatus.Columns().UpTime+" = ?", req.UpTime)
	}
	if req.DownTime != "" {
		m = m.Where(dao.SysDeviceStatus.Columns().DownTime+" = ?", req.DownTime)
	}
	if req.LastDataUpdateTime != "" {
		m = m.Where(dao.SysDeviceStatus.Columns().LastDataUpdateTime+" = ?", req.LastDataUpdateTime)
	}
	err = g.Try(ctx, func(ctx context.Context) {
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取设备状态列表失败")
		order := "status_id asc"
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

// ChangeStatus 修改状态
func (s *sDeviceStatus) ChangeStatus(ctx context.Context, deviceId int, status int) error {

	time := gtime.Now().Timestamp()

	// 时间戳转换为时间字符串
	timeStr := gtime.NewFromTimeStamp(time).String()

	// 更新设备信息表
	_, _ = g.Model(dao.SysDeviceInfo.Table()).Where(dao.SysDeviceInfo.Columns().Id, deviceId).Update(g.Map{
		dao.SysDeviceInfo.Columns().Status:             status,
		dao.SysDeviceInfo.Columns().LastDataUpdateTime: timeStr,
		dao.SysDeviceInfo.Columns().UpTime:             timeStr,
		dao.SysDeviceInfo.Columns().DownTime:           timeStr,
	})

	//// 离线
	//if status == consts.DeviceStatusOffLine {
	//	_, err := dao.SysDeviceStatus.Ctx(ctx).Where(dao.SysDeviceStatus.Columns().DeviceId+" = ?", deviceId).Save(g.Map{
	//		dao.SysDeviceStatus.Columns().DeviceId:           deviceId,
	//		dao.SysDeviceStatus.Columns().Status:             status,
	//		dao.SysDeviceStatus.Columns().DownTime:           timeStr,
	//		dao.SysDeviceStatus.Columns().LastDataUpdateTime: timeStr,
	//	})
	//	return err
	//}
	//
	//// 在线
	//if status == consts.DeviceStatusOnLine {
	//	_, err := dao.SysDeviceStatus.Ctx(ctx).Where(dao.SysDeviceStatus.Columns().DeviceId+" = ?", deviceId).Save(g.Map{
	//		dao.SysDeviceStatus.Columns().DeviceId:           deviceId,
	//		dao.SysDeviceStatus.Columns().Status:             status,
	//		dao.SysDeviceStatus.Columns().UpTime:             timeStr,
	//		dao.SysDeviceStatus.Columns().LastDataUpdateTime: timeStr,
	//	})
	//	return err
	//}
	//
	//// 上报数据
	//if status == consts.DeviceStatusDataUp {
	//	_, err := dao.SysDeviceStatus.Ctx(ctx).Where(dao.SysDeviceStatus.Columns().DeviceId+" = ?", deviceId).Save(g.Map{
	//		dao.SysDeviceStatus.Columns().DeviceId:           deviceId,
	//		dao.SysDeviceStatus.Columns().Status:             status,
	//		dao.SysDeviceStatus.Columns().LastDataUpdateTime: time,
	//	})
	//
	//	return err
	//}

	return nil
}
