package deviceInfo

import (
	"backend/api/v1/device"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDeviceInfo struct {
}

func init() {
	service.RegisterDeviceInfo(New())
}

func New() service.IDeviceInfo {
	return &sDeviceInfo{}
}

// List 获取设备信息列表
func (s *sDeviceInfo) List(ctx context.Context, req *device.InfoSearchReq) (total, page int, list []*entity.SysDeviceInfo, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}

	m := dao.SysDeviceInfo.Ctx(ctx)
	if req.Name != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().Name+" like ?", "%"+req.Name+"%")
	}
	if req.Group != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().Group+" = ?", req.Group)
	}
	if req.Sn != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().Sn+" = ?", req.Sn)
	}
	if req.Pwd != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().Pwd+" = ?", req.Pwd)
	}
	if req.Kind != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().Kind+" = ?", req.Kind)
	}
	if req.Logo != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().Logo+" = ?", req.Logo)
	}
	if req.Monitor != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().Monitor+" = ?", req.Monitor)
	}
	if req.Location != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().Location+" = ?", req.Location)
	}
	if req.BeginTime != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().CreatedAt+" >=", req.BeginTime)
	}
	if req.EndTime != "" {
		m = m.Where(dao.SysDeviceInfo.Columns().CreatedAt+" <", req.EndTime)
	}

	at := dao.SysDeviceInfo.Table()
	ac := dao.SysDeviceInfo.Columns()
	bt := dao.SysDeviceStatus.Table()
	bc := dao.SysDeviceStatus.Columns()
	err = g.Try(ctx, func(ctx context.Context) {
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取设备信息表总行数失败")
		m = m.LeftJoin(bt, bt+"."+bc.DeviceId+"="+at+"."+ac.Id)
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		m = m.Fields(at + ".*")
		m = m.Page(page, req.PageSize)
		m = m.Order(order)
		err = m.Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取设备信息列表失败")
	})
	return
}

func (s *sDeviceInfo) Auth(ctx context.Context, sn, pwd string) (status bool, err error) {
	var deviceInfo *entity.SysDeviceInfo
	err = dao.SysDeviceInfo.Ctx(ctx).Where("sn=? and pwd=?", sn, pwd).Scan(&deviceInfo)
	liberr.ErrIsNil(ctx, err, "设备SN或密码错误")

	if deviceInfo == nil || deviceInfo.Id < 1 {
		return false, gerror.New("设备SN或密码错误")
	}

	return true, nil
}

func (s *sDeviceInfo) GetAllInfo(ctx context.Context, id int, sn string) (info *model.DeviceAllInfo, err error) {
	info = &model.DeviceAllInfo{}
	if id == 0 && len(sn) < 1 {
		err = gerror.New("参数错误")
		return
	}

	g.Log().Printf(ctx, "id:%v, sn:%v", id, sn)

	at := dao.SysDeviceInfo.Table()
	ac := dao.SysDeviceInfo.Columns()
	bt := dao.SysDeviceStatus.Table()
	bc := dao.SysDeviceStatus.Columns()

	if id != 0 {
		//err = dao.SysDeviceInfo.Ctx(ctx).LeftJoin(dao.SysDeviceStatus.Table(), dao.SysDeviceStatus.Table()+"."+dao.SysDeviceStatus.Columns().DeviceId+"="+dao.SysDeviceInfo.Table()+"."+dao.SysDeviceInfo.Columns().Id).Where(dao.SysDeviceInfo.Columns().Id, id).Scan(&info.Info)

		m := dao.SysDeviceInfo.Ctx(ctx).LeftJoin(bt, bt+"."+bc.DeviceId+"="+at+"."+ac.Id).Where(at+"."+ac.Id, id)
		err = m.Fields(bt + ".*").Scan(&info.Info)
		liberr.ErrIsNil(ctx, err, "不存在该设备信息")

	} else if len(sn) > 0 {
		//err = dao.SysDeviceInfo.Ctx(ctx).LeftJoin(dao.SysDeviceStatus.Table(), dao.SysDeviceStatus.Table()+"."+dao.SysDeviceStatus.Columns().DeviceId+"="+dao.SysDeviceInfo.Table()+"."+dao.SysDeviceInfo.Columns().Id).Where(dao.SysDeviceInfo.Columns().Sn, sn).Scan(&info.Info)
	} else {
		err = gerror.New("参数错误")
		return
	}

	if err != nil {
		liberr.ErrIsNil(ctx, err, "获取设备信息失败")
		return
	}

	if info == nil || err != nil {
		liberr.ErrIsNil(ctx, err, "获取设备信息失败")
		return
	}

	info.Kind, err = service.DeviceKind().Get(ctx, info.Info.Kind)
	if err != nil || info.Kind == nil {
		liberr.ErrIsNil(ctx, err, "获取设备类别失败, kindId:"+gconv.String(info.Info.Kind))
		return
	}
	info.CategoryList, err = service.DeviceCategory().KindGet(ctx, info.Info.Kind)
	if err != nil || info.CategoryList == nil {
		liberr.ErrIsNil(ctx, err, "获取设备类别列表失败, kindId:"+gconv.String(info.Info.Kind))
		return
	}
	return
}
