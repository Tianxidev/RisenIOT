package deviceStrategy

import (
	"backend/api/v1/device"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
)

/// 设备策略

type sDeviceStrategy struct {
}

func New() *sDeviceStrategy {
	return &sDeviceStrategy{}
}

func init() {
	service.RegisterDeviceStrategy(New())
}

// Add 添加策略
func (s *sDeviceStrategy) Add(ctx context.Context, req *device.StrategyAddReq) (*device.StrategyAddRes, error) {
	var err error
	res := new(device.StrategyAddRes)

	// 获取用户ID
	userId := service.UserCtx().GetUserId(ctx)

	// 创建策略实体
	entity := entity.SysDeviceStrategy{
		Id:       0,
		Name:     req.Name,
		Type:     req.Type,
		Cron:     req.Cron,
		DeviceId: req.DeviceId,
		AuthorId: int(userId),
	}

	// 插入数据库
	_, err = dao.SysDeviceStrategy.Ctx(ctx).Insert(entity)
	liberr.ErrIsNil(ctx, err, "添加策略失败")

	return res, err
}

// List 查询策略列表
func (s *sDeviceStrategy) List(ctx context.Context, req *device.StrategySearchReq) (*device.StrategySearchRes, error) {
	res := new(device.StrategySearchRes)

	var list []*entity.SysDeviceStrategy

	// 根据上下文获取用户ID
	userId := service.UserCtx().GetUserId(ctx)

	// 查询策略列表
	m := dao.SysDeviceStrategy.Ctx(ctx)
	err := m.Where("author_id=?", int(userId)).Scan(&list)
	liberr.ErrIsNil(ctx, err, "查询策略列表失败")

	// 遍历策略列表返回res
	for _, v := range list {

		var strategyType string

		// 判断是定时策略还是数据策略
		if v.Type == consts.DeviceStrategyDataType {
			strategyType = "数据策略"
		} else {
			strategyType = "时间策略"
		}

		// 添加到响应列表
		res.List = append(res.List, &model.SysDeviceStrategy{
			Id:       v.Id,
			Name:     v.Name,
			Type:     strategyType,
			TypeId:   v.Type,
			Cron:     v.Cron,
			DeviceId: v.DeviceId,
		})
	}

	// 设置数量
	res.Total = len(list)

	return res, err
}
