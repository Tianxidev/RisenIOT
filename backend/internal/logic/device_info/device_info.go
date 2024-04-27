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
	"fmt"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

type sDeviceInfo struct {
}

func init() {
	service.RegisterDeviceInfo(New())
}

func New() *sDeviceInfo {
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

// Auth 设备授权
func (s *sDeviceInfo) Auth(ctx context.Context, sn, pwd string) (status bool, err error) {
	var deviceInfo *entity.SysDeviceInfo
	err = dao.SysDeviceInfo.Ctx(ctx).Where("sn=? and pwd=?", sn, pwd).Scan(&deviceInfo)
	liberr.ErrIsNil(ctx, err, "设备SN或密码错误")

	if deviceInfo == nil || deviceInfo.Id < 1 {
		return false, gerror.New("设备SN或密码错误")
	}

	return true, nil
}

// Info 获取设备信息
func (s *sDeviceInfo) Info(ctx context.Context, id int, sn string) (deviceInfo *model.DeviceAllInfo, err error) {
	deviceInfo = &model.DeviceAllInfo{}
	deviceInfoTable := dao.SysDeviceInfo.Table()
	deviceInfoColumns := dao.SysDeviceInfo.Columns()
	if id == 0 && len(sn) < 1 {
		err = gerror.New("参数错误")
		return
	}

	g.Log().Printf(ctx, "设备ID:%v, 设备SN:%v", id, sn)

	// 查询设备信息
	data, err := g.Model(deviceInfoTable).Where(deviceInfoColumns.Id, id).One()
	liberr.ErrIsNil(ctx, err, "查询设备信息失败")

	// 映射设备信息
	err = data.Struct(&deviceInfo.Info)
	liberr.ErrIsNil(ctx, err, "映射设备信息失败")

	// 查询设备产品信息
	deviceInfo.Kind, err = service.DeviceKind().Get(ctx, deviceInfo.Info.Kind)
	liberr.ErrIsNil(ctx, err, "查询设备产品信息失败, kindId:"+gconv.String(deviceInfo.Info.Kind))

	// 查询产品数据类型信息
	deviceInfo.CategoryList, err = service.DeviceCategory().KindGet(ctx, deviceInfo.Info.Kind)
	liberr.ErrIsNil(ctx, err, "查询产品数据类型信息失败, kindId:"+gconv.String(deviceInfo.Info.Kind))

	return
}

// Add 添加设备信息
func (s *sDeviceInfo) Add(ctx context.Context, req *device.InfoAddReq) (err error) {
	m := dao.SysDeviceInfo.Ctx(ctx)

	// 判断设备SN是否为空, 如果为空则随机生成一个SN
	if req.Sn == "" {
		req.Sn = guid.S([]byte(gconv.String(gtime.TimestampNano())))
	}

	// 判断设备密码是否为空, 如果为空则随机生成一个密码
	if req.Pwd == "" {
		req.Pwd = guid.S([]byte(gconv.String(gtime.TimestampNano())))
	}

	// 判断设备分组是否为空, 如果为空则默认分组
	if req.Group == 0 {
		req.Group = -1
	}

	// 根据当前用户ID和设备名称查询是否存在相同的设备名称
	count, err := m.Where("create_by=? and name=?", int(service.UserCtx().GetUserId(ctx)), req.Name).Count()
	liberr.ErrIsNil(ctx, err, "查询设备信息失败")
	liberr.ValueIsTrue(ctx, count > 0, "已存在同名设备, 请更换设备名称")

	deviceInfo := &entity.SysDeviceInfo{
		Name:     req.Name,
		Group:    req.Group,
		Sn:       req.Sn,
		Pwd:      req.Pwd,
		Kind:     req.Kind,
		Monitor:  req.Monitor,
		Status:   1,
		TimeOut:  30,
		CreateBy: int(service.UserCtx().GetUserId(ctx)),
	}

	// 写入数据库
	_, err = dao.SysDeviceInfo.Ctx(ctx).Data(deviceInfo).Insert()
	liberr.ErrIsNil(ctx, err, "添加设备信息失败")
	return
}

// Edit 编辑设备信息
func (s *sDeviceInfo) Edit(ctx context.Context, req *device.InfoEditReq) (err error) {
	m := dao.SysDeviceInfo.Ctx(ctx)

	// 根据设备ID和当前用户ID查询设备信息
	deviceInfo := &entity.SysDeviceInfo{}
	err = m.Where("id=? and create_by=?", req.Id, int(service.UserCtx().GetUserId(ctx))).Scan(&deviceInfo)
	liberr.ErrIsNil(ctx, err, "查询设备信息失败")
	liberr.ValueIsTrue(ctx, deviceInfo == nil, "设备信息不存在")

	// 判断设备分组是否为空, 如果为空则默认分组
	if req.Group == 0 {
		req.Group = -1
	}

	// 更新设备信息
	_, err = m.Data(g.Map{
		"name":     req.Name,
		"group":    req.Group,
		"kind":     req.Kind,
		"monitor":  req.Monitor,
		"location": req.Location,
	}).Where("id=? and create_by=?", req.Id, int(service.UserCtx().GetUserId(ctx))).Update()
	liberr.ErrIsNil(ctx, err, "编辑设备信息失败")
	return
}

// Delete 删除设备信息
func (s *sDeviceInfo) Delete(ctx context.Context, ids []int) (err error) {
	m := dao.SysDeviceInfo
	userId := service.UserCtx().GetUserId(ctx)

	// 根据设备ID和当前用户ID删除设备信息
	for _, id := range ids {
		query := fmt.Sprintf("%s=? and %s=?", m.Columns().Id, m.Columns().CreateBy)
		_, err = m.Ctx(ctx).Where(query, id, userId).Delete()
		liberr.ErrIsNil(ctx, err, "删除设备信息失败")
	}
	return
}

// UpdateDataLastTime 更新设备数据最后上报时间
func (s *sDeviceInfo) UpdateDataLastTime(ctx context.Context, id int) (err error) {
	_, err = dao.SysDeviceInfo.Ctx(ctx).Data(g.Map{
		"last_time": gtime.Now().String(),
	}).Where("id=?", id).Update()
	liberr.ErrIsNil(ctx, err, "更新设备数据最后上报时间失败")
	return
}

// Count 统计设备数量
func (s *sDeviceInfo) Count(ctx context.Context) (total int, err error) {

	// 统计我的设备数量
	total, err = dao.SysDeviceInfo.Ctx(ctx).Where("create_by=?", int(service.UserCtx().GetUserId(ctx))).Count()

	liberr.ErrIsNil(ctx, err, "统计设备数量失败")
	return
}

// OnlineCount 在线设备数量
func (s *sDeviceInfo) OnlineCount(ctx context.Context) (total int, err error) {
	// 统计我的在线设备数量
	userId := int(service.UserCtx().GetUserId(ctx))
	total, err = dao.SysDeviceInfo.Ctx(ctx).Where("create_by=? and status=1", userId).Count()
	liberr.ErrIsNil(ctx, err, "统计在线设备数量失败")
	return
}

// OfflineCount 离线设备数量
func (s *sDeviceInfo) OfflineCount(ctx context.Context) (total int, err error) {
	// 统计我的离线设备数量
	userId := int(service.UserCtx().GetUserId(ctx))
	total, err = g.Model(dao.SysDeviceInfo.Table()).Where("create_by=? and status = 0", userId).Count()
	liberr.ErrIsNil(ctx, err, "统计离线设备数量失败")
	return
}

// KeepAlive 设备保活状态检查
func (s *sDeviceInfo) KeepAlive(ctx context.Context, time *gtime.Time) {
	deviceList := make([]*entity.SysDeviceInfo, 0)

	// 查询所有设备信息
	err := g.Model(dao.SysDeviceInfo.Table()).Scan(&deviceList)
	liberr.ErrIsNil(ctx, err, "查询设备信息失败")

	// 遍历设备信息, 判断设备是否在线
	for _, deviceInfo := range deviceList {

		// deviceInfo.LastDataUpdateTime 转换为时间戳
		nowTime := time.Timestamp()
		lastTime := deviceInfo.LastDataUpdateTime.Timestamp()

		if deviceInfo.Status == 0 {
			continue
		}

		// 查询产品超时时间
		kindInfo, err := service.DeviceKind().Get(ctx, deviceInfo.Kind)
		liberr.ErrIsNil(ctx, err, "查询产品超时时间失败")

		// 判断设备最后上报时间是否超过设备心跳超时时间
		// 当前系统时间戳 - 设备最后上报时间戳 > 设备心跳超时时间
		if nowTime-lastTime > int64(kindInfo.TimeOut) {
			userInfo := new(entity.SysUser)

			// 设备离线
			g.Log().Info(ctx, "离线设备=>", deviceInfo.Name, "设备ID:", deviceInfo.Id)
			_, err = dao.SysDeviceInfo.Ctx(ctx).Where("id=?", deviceInfo.Id).Update(g.Map{
				"status": 0,
			})

			// 查询用户消息推送配置
			one, err := g.Model(dao.SysUser.Table()).Where("id=?", deviceInfo.CreateBy).One()
			err = one.Struct(userInfo)
			liberr.ErrIsNil(ctx, err, "查询用户消息推送配置失败")

			// 发送设备离线消息
			msg := fmt.Sprintf("[设备离线] 设备名称:%s, 设备ID:%d", deviceInfo.Name, deviceInfo.Id)

			// 消息url编码
			msg = gconv.String(gurl.Encode(msg))

			g.Log().Info(ctx, "发送离线消息给用户=>", userInfo.UserName, "-", userInfo.Pushdeer, "用户ID:", userInfo.Id, "推送设备=>", userInfo.Pushdeer, "消息内容:", msg)
			service.SysPush().TextMsgSendPushDeer(ctx, userInfo.Pushdeer, msg)

			liberr.ErrIsNil(ctx, err, "设备离线失败")
		} else {
			g.Log().Info(ctx, "在线设备=>", deviceInfo.Name, "设备ID:", deviceInfo.Id, "当前超时:", nowTime-lastTime, "设定超时:", kindInfo.TimeOut)
		}

	}

}
